// calc-cli is a scientific calculator for the command line.
//
// It supports basic arithmetic, scientific functions, unit conversions,
// and interactive mode with history.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		runInteractive()
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "eval", "-e":
		if err := runEval(args); err != nil {
			fmt.Fprintf(os.Stderr, "calc-cli: %v\n", err)
			os.Exit(1)
		}
	case "convert":
		if err := runConvert(args); err != nil {
			fmt.Fprintf(os.Stderr, "calc-cli: %v\n", err)
			os.Exit(1)
		}
	case "help", "--help", "-h":
		printUsage()
	case "version", "--version", "-v":
		fmt.Println("calc-cli v0.1.0")
	default:
		if err := runEval(append([]string{command}, args...)); err != nil {
			fmt.Fprintf(os.Stderr, "calc-cli: %v\n", err)
			os.Exit(1)
		}
	}
}

func runEval(args []string) error {
	expr := strings.Join(args, " ")
	if expr == "" {
		return fmt.Errorf("expression is required")
	}

	result, err := evaluate(expr)
	if err != nil {
		return err
	}

	fmt.Printf("%.10g\n", result)
	return nil
}

func runConvert(args []string) error {
	if len(args) < 3 {
		fmt.Println("Usage: calc-cli convert <value> <from> <to>")
		fmt.Println("\nAvailable units:")
		printUnitCategories()
		return nil
	}

	value, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return fmt.Errorf("invalid value: %w", err)
	}

	from := strings.ToLower(args[1])
	to := strings.ToLower(args[2])

	result, err := convertUnit(value, from, to)
	if err != nil {
		return err
	}

	fmt.Printf("%.10g %s = %.10g %s\n", value, from, result, to)
	return nil
}

func runInteractive() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("calc-cli - Scientific Calculator")
	fmt.Println("Type 'help' for commands, 'quit' to exit.")
	fmt.Println()

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		switch strings.ToLower(line) {
		case "quit", "exit", "q":
			return
		case "help", "h":
			printUsage()
			continue
		}

		result, err := evaluate(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			continue
		}

		fmt.Printf("= %.10g\n", result)
	}
}

func evaluate(expr string) (float64, error) {
	expr = strings.TrimSpace(expr)
	expr = strings.ToLower(expr)

	// Handle special functions
	if strings.HasPrefix(expr, "sqrt(") || strings.HasPrefix(expr, "sqrt ") {
		return handleUnaryFunc(expr, math.Sqrt, "sqrt")
	}
	if strings.HasPrefix(expr, "sin(") || strings.HasPrefix(expr, "sin ") {
		return handleUnaryFunc(expr, math.Sin, "sin")
	}
	if strings.HasPrefix(expr, "cos(") || strings.HasPrefix(expr, "cos ") {
		return handleUnaryFunc(expr, math.Cos, "cos")
	}
	if strings.HasPrefix(expr, "tan(") || strings.HasPrefix(expr, "tan ") {
		return handleUnaryFunc(expr, math.Tan, "tan")
	}
	if strings.HasPrefix(expr, "asin(") || strings.HasPrefix(expr, "asin ") {
		return handleUnaryFunc(expr, math.Asin, "asin")
	}
	if strings.HasPrefix(expr, "acos(") || strings.HasPrefix(expr, "acos ") {
		return handleUnaryFunc(expr, math.Acos, "acos")
	}
	if strings.HasPrefix(expr, "atan(") || strings.HasPrefix(expr, "atan ") {
		return handleUnaryFunc(expr, math.Atan, "atan")
	}
	if strings.HasPrefix(expr, "log(") || strings.HasPrefix(expr, "log ") {
		return handleUnaryFunc(expr, math.Log, "log")
	}
	if strings.HasPrefix(expr, "log2(") || strings.HasPrefix(expr, "log2 ") {
		return handleUnaryFunc(expr, math.Log2, "log2")
	}
	if strings.HasPrefix(expr, "log10(") || strings.HasPrefix(expr, "log10 ") {
		return handleUnaryFunc(expr, math.Log10, "log10")
	}
	if strings.HasPrefix(expr, "exp(") || strings.HasPrefix(expr, "exp ") {
		return handleUnaryFunc(expr, math.Exp, "exp")
	}
	if strings.HasPrefix(expr, "abs(") || strings.HasPrefix(expr, "abs ") {
		return handleUnaryFunc(expr, math.Abs, "abs")
	}
	if strings.HasPrefix(expr, "ceil(") || strings.HasPrefix(expr, "ceil ") {
		return handleUnaryFunc(expr, math.Ceil, "ceil")
	}
	if strings.HasPrefix(expr, "floor(") || strings.HasPrefix(expr, "floor ") {
		return handleUnaryFunc(expr, math.Floor, "floor")
	}

	// Handle constants
	expr = replaceConstants(expr)

	// Simple expression parser
	return parseExpression(expr)
}

func handleUnaryFunc(expr string, fn func(float64) float64, fnStr string) (float64, error) {
	// Strip function name and optional space/paren
	if strings.HasPrefix(expr, fnStr+"(") {
		expr = expr[len(fnStr)+1:]
		// Strip matching closing paren
		if len(expr) > 0 && expr[len(expr)-1] == ')' {
			expr = expr[:len(expr)-1]
		}
	} else if strings.HasPrefix(expr, fnStr+" ") {
		expr = expr[len(fnStr)+1:]
	}
	expr = strings.TrimSpace(expr)

	// Replace constants (pi, e) before parsing
	expr = replaceConstants(expr)

	arg, err := parseExpression(expr)
	if err != nil {
		return 0, fmt.Errorf("invalid argument for %s: %w", fnStr, err)
	}
	return fn(arg), nil
}

func replaceConstants(expr string) string {
	expr = strings.ReplaceAll(expr, "pi", fmt.Sprintf("(%.20f)", math.Pi))
	expr = strings.ReplaceAll(expr, "e", fmt.Sprintf("(%.20f)", math.E))
	return expr
}

func parseExpression(expr string) (float64, error) {
	expr = strings.ReplaceAll(expr, " ", "")
	expr = strings.ReplaceAll(expr, "^", "**")

	// Simple recursive descent parser
	p := &parser{input: expr, pos: 0}
	result, err := p.parseExpr()
	if err != nil {
		return 0, err
	}

	if p.pos < len(p.input) {
		return 0, fmt.Errorf("unexpected character at position %d: %c", p.pos, p.input[p.pos])
	}

	return result, nil
}

type parser struct {
	input string
	pos   int
}

func (p *parser) parseExpr() (float64, error) {
	left, err := p.parseTerm()
	if err != nil {
		return 0, err
	}

	for p.pos < len(p.input) && (p.input[p.pos] == '+' || p.input[p.pos] == '-') {
		op := p.input[p.pos]
		p.pos++
		right, err := p.parseTerm()
		if err != nil {
			return 0, err
		}
		if op == '+' {
			left += right
		} else {
			left -= right
		}
	}

	return left, nil
}

func (p *parser) parseTerm() (float64, error) {
	left, err := p.parsePower()
	if err != nil {
		return 0, err
	}

	for p.pos < len(p.input) && (p.input[p.pos] == '*' || p.input[p.pos] == '/') {
		op := p.input[p.pos]
		p.pos++
		right, err := p.parsePower()
		if err != nil {
			return 0, err
		}
		if op == '*' {
			left *= right
		} else {
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			left /= right
		}
	}

	return left, nil
}

func (p *parser) parsePower() (float64, error) {
	base, err := p.parseUnary()
	if err != nil {
		return 0, err
	}

	for p.pos+1 < len(p.input) && p.input[p.pos] == '*' && p.input[p.pos+1] == '*' {
		p.pos += 2
		exp, err := p.parsePower()
		if err != nil {
			return 0, err
		}
		base = math.Pow(base, exp)
	}

	return base, nil
}

func (p *parser) parseUnary() (float64, error) {
	if p.pos < len(p.input) && (p.input[p.pos] == '+' || p.input[p.pos] == '-') {
		op := p.input[p.pos]
		p.pos++
		val, err := p.parseUnary()
		if err != nil {
			return 0, err
		}
		if op == '-' {
			return -val, nil
		}
		return val, nil
	}
	return p.parseNumber()
}

func (p *parser) parseNumber() (float64, error) {
	if p.pos >= len(p.input) {
		return 0, fmt.Errorf("unexpected end of expression")
	}

	if p.input[p.pos] == '(' {
		p.pos++
		val, err := p.parseExpr()
		if err != nil {
			return 0, err
		}
		if p.pos >= len(p.input) || p.input[p.pos] != ')' {
			return 0, fmt.Errorf("missing closing parenthesis")
		}
		p.pos++
		return val, nil
	}

	start := p.pos
	for p.pos < len(p.input) && (isDigit(p.input[p.pos]) || p.input[p.pos] == '.') {
		p.pos++
	}

	if start == p.pos {
		return 0, fmt.Errorf("expected number at position %d", p.pos)
	}

	return strconv.ParseFloat(p.input[start:p.pos], 64)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// Unit conversion system
func convertUnit(value float64, from, to string) (float64, error) {
	if from == to {
		return value, nil
	}

	// Length
	lengthUnits := map[string]float64{
		"mm": 0.001, "cm": 0.01, "m": 1, "km": 1000,
		"in": 0.0254, "ft": 0.3048, "yd": 0.9144, "mi": 1609.344,
	}
	if f, ok := lengthUnits[from]; ok {
		if t, ok := lengthUnits[to]; ok {
			return value * f / t, nil
		}
	}

	// Weight
	weightUnits := map[string]float64{
		"mg": 0.000001, "g": 0.001, "kg": 1, "t": 1000,
		"oz": 0.0283495, "lb": 0.453592,
	}
	if f, ok := weightUnits[from]; ok {
		if t, ok := weightUnits[to]; ok {
			return value * f / t, nil
		}
	}

	// Temperature (special case)
	tempUnits := map[string]bool{"c": true, "f": true, "k": true}
	if tempUnits[from] && tempUnits[to] {
		return convertTemp(value, from, to), nil
	}

	// Data
	dataUnits := map[string]float64{
		"b": 1, "kb": 1024, "mb": 1048576, "gb": 1073741824, "tb": 1099511627776,
	}
	if f, ok := dataUnits[from]; ok {
		if t, ok := dataUnits[to]; ok {
			return value * f / t, nil
		}
	}

	return 0, fmt.Errorf("unknown units: %s, %s", from, to)
}

func convertTemp(value float64, from, to string) float64 {
	celsius := value
	switch from {
	case "f":
		celsius = (value - 32) * 5 / 9
	case "k":
		celsius = value - 273.15
	}

	switch to {
	case "f":
		return celsius*9/5 + 32
	case "k":
		return celsius + 273.15
	default:
		return celsius
	}
}

func printUnitCategories() {
	fmt.Println("  Length: mm, cm, m, km, in, ft, yd, mi")
	fmt.Println("  Weight: mg, g, kg, t, oz, lb")
	fmt.Println("  Temperature: c, f, k")
	fmt.Println("  Data: b, kb, mb, gb, tb")
}

func printUsage() {
	usage := `
calc-cli - Scientific Calculator for the Command Line

USAGE:
    calc-cli [command] [arguments]

COMMANDS:
    calc-cli                    Start interactive mode
    calc-cli eval "2 + 3 * 4"   Evaluate expression
    calc-cli "2 + 3 * 4"        Same as eval (shortcut)
    calc-cli convert 100 c f    Convert units

SUPPORTED OPERATIONS:
    + - * /                     Basic arithmetic
    ** or ^                     Exponentiation
    ( )                         Grouping

FUNCTIONS:
    sqrt(x)   Square root
    sin(x)    Sine (radians)
    cos(x)    Cosine (radians)
    tan(x)    Tangent (radians)
    asin(x)   Arc sine
    acos(x)   Arc cosine
    atan(x)   Arc tangent
    log(x)    Natural logarithm
    log2(x)   Base-2 logarithm
    log10(x)  Base-10 logarithm
    exp(x)    Exponential (e^x)
    abs(x)    Absolute value
    ceil(x)   Ceiling
    floor(x)  Floor

CONSTANTS:
    pi        3.14159...
    e         2.71828...

CONVERSIONS:
    Length:        mm, cm, m, km, in, ft, yd, mi
    Weight:        mg, g, kg, t, oz, lb
    Temperature:   c, f, k
    Data:          b, kb, mb, gb, tb

EXAMPLES:
    calc-cli "2 + 3"                    # 5
    calc-cli "sqrt(16)"                 # 4
    calc-cli "sin(pi/2)"                # 1
    calc-cli "2 ** 10"                  # 1024
    calc-cli convert 100 c f            # 212 f
    calc-cli convert 5 km mi            # 3.10686 mi
    calc-cli convert 1 gb mb            # 1024 mb
`
	fmt.Println(strings.TrimSpace(usage))
}
