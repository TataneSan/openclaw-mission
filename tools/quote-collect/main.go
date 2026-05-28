package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Quote struct {
	ID        int
	Text      string
	Author    string
	Source    string
	Category  string
	Favorite  int
	CreatedAt string
}

func openDB() (*sql.DB, error) {
	dataDir := os.Getenv("HOME") + "/.quote-collect"
	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, fmt.Errorf("cannot create data directory: %w", err)
	}
	db, err := sql.Open("sqlite3", dataDir+"/quotes.db?_journal_mode=WAL")
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %w", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS quotes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			author TEXT DEFAULT 'Unknown',
			source TEXT DEFAULT '',
			category TEXT DEFAULT 'general',
			favorite INTEGER DEFAULT 0,
			created_at TEXT NOT NULL
		);
	`)
	return db, err
}

func add(db *sql.DB, text, author, source, category string) error {
	_, err := db.Exec(
		"INSERT INTO quotes (text, author, source, category, created_at) VALUES (?, ?, ?, ?, ?)",
		text, author, source, category, time.Now().UTC().Format(time.RFC3339),
	)
	if err != nil {
		return err
	}
	fmt.Printf("Added quote #%d.\n", 0)
	return nil
}

func listQuotes(db *sql.DB, category, sortBy string, favOnly bool) error {
	var rows *sql.Rows
	var err error

	base := "SELECT id, text, author, source, category, favorite, created_at FROM quotes WHERE 1=1"
	args := []interface{}{}

	if favOnly {
		base += " AND favorite = 1"
	}
	if category != "" && category != "all" {
		base += " AND category = ?"
		args = append(args, category)
	}

	switch sortBy {
	case "author":
		base += " ORDER BY author ASC"
	case "date", "newest":
		base += " ORDER BY created_at DESC"
	case "random":
		base += " ORDER BY RANDOM()"
	default:
		base += " ORDER BY id ASC"
	}

	rows, err = db.Query(base, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.Author, &q.Source, &q.Category, &q.Favorite, &q.CreatedAt); err != nil {
			return err
		}
		quotes = append(quotes, q)
	}

	if len(quotes) == 0 {
		fmt.Println("No quotes found.")
		return nil
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, q := range quotes {
		fav := ""
		if q.Favorite == 1 {
			fav = "*"
		}
		trunc := q.Text
		if len(trunc) > 60 {
			trunc = trunc[:57] + "..."
		}
		fmt.Fprintf(w, "[%d] %s \"%s\" — %s (%s%s)\n", q.ID, fav, trunc, q.Author, q.Category, q.Source)
	}
	w.Flush()
	fmt.Printf("\nTotal: %d quote(s)\n", len(quotes))
	return nil
}

func showQuote(db *sql.DB, id int) error {
	var q Quote
	err := db.QueryRow(
		"SELECT id, text, author, source, category, favorite, created_at FROM quotes WHERE id = ?",
		id,
	).Scan(&q.ID, &q.Text, &q.Author, &q.Source, &q.Category, &q.Favorite, &q.CreatedAt)
	if err == sql.ErrNoRows {
		return fmt.Errorf("quote %d not found", id)
	}
	if err != nil {
		return err
	}

	fav := ""
	if q.Favorite == 1 {
		fav = " [favorite]"
	}
	fmt.Printf("\n  \"%s\"\n", q.Text)
	fmt.Printf("  — %s\n", q.Author)
	if q.Source != "" {
		fmt.Printf("  Source: %s\n", q.Source)
	}
	fmt.Printf("  Category: %s%s\n", q.Category, fav)
	fmt.Printf("  Added: %s\n", q.CreatedAt[:10])
	fmt.Println()
	return nil
}

func randomQuote(db *sql.DB) error {
	var q Quote
	err := db.QueryRow(
		"SELECT id, text, author, source, category, favorite, created_at FROM quotes ORDER BY RANDOM() LIMIT 1",
	).Scan(&q.ID, &q.Text, &q.Author, &q.Source, &q.Category, &q.Favorite, &q.CreatedAt)
	if err == sql.ErrNoRows {
		fmt.Println("No quotes in collection. Add some first!")
		return nil
	}
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("  \"%s\"\n", q.Text)
	fmt.Printf("  — %s\n", q.Author)
	if q.Source != "" {
		fmt.Printf("  Source: %s\n", q.Source)
	}
	fmt.Printf("  Category: %s  ID: %d\n", q.Category, q.ID)
	fmt.Println()
	return nil
}

func toggleFav(db *sql.DB, id int) error {
	var current int
	err := db.QueryRow("SELECT favorite FROM quotes WHERE id = ?", id).Scan(&current)
	if err == sql.ErrNoRows {
		return fmt.Errorf("quote %d not found", id)
	}
	if err != nil {
		return err
	}
	newVal := 1
	if current == 1 {
		newVal = 0
	}
	_, err = db.Exec("UPDATE quotes SET favorite = ? WHERE id = ?", newVal, id)
	if err != nil {
		return err
	}
	if newVal == 1 {
		fmt.Printf("Quote %d marked as favorite.\n", id)
	} else {
		fmt.Printf("Quote %d removed from favorites.\n", id)
	}
	return nil
}

func remove(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM quotes WHERE id = ?", id)
	if err != nil {
		return err
	}
	n, _ := result.RowsAffected()
	if n == 0 {
		return fmt.Errorf("quote %d not found", id)
	}
	fmt.Printf("Removed quote %d.\n", id)
	return nil
}

func search(db *sql.DB, query string) error {
	rows, err := db.Query(
		"SELECT id, text, author, source, category, favorite, created_at FROM quotes WHERE text LIKE ? OR author LIKE ? OR source LIKE ? ORDER BY id",
		"%"+query+"%", "%"+query+"%", "%"+query+"%",
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.Author, &q.Source, &q.Category, &q.Favorite, &q.CreatedAt); err != nil {
			return err
		}
		quotes = append(quotes, q)
	}

	if len(quotes) == 0 {
		fmt.Printf("No quotes matching %q\n", query)
		return nil
	}

	for _, q := range quotes {
		fav := ""
		if q.Favorite == 1 {
			fav = " *"
		}
		trunc := q.Text
		if len(trunc) > 60 {
			trunc = trunc[:57] + "..."
		}
		fmt.Printf("[%d]%s \"%s\" — %s\n", q.ID, fav, trunc, q.Author)
	}
	fmt.Printf("\nFound: %d quote(s)\n", len(quotes))
	return nil
}

func showStats(db *sql.DB) error {
	var total int
	db.QueryRow("SELECT COUNT(*) FROM quotes").Scan(&total)
	var favCount int
	db.QueryRow("SELECT COUNT(*) FROM quotes WHERE favorite = 1").Scan(&favCount)

	rows, _ := db.Query("SELECT category, COUNT(*) as cnt FROM quotes GROUP BY category ORDER BY cnt DESC")
	var cats []struct {
		Name string
		Cnt  int
	}
	for rows.Next() {
		var c struct {
			Name string
			Cnt  int
		}
		rows.Scan(&c.Name, &c.Cnt)
		cats = append(cats, c)
	}
	rows.Close()

	fmt.Println("Quote Collection Statistics")
	fmt.Println("---------------------------")
	fmt.Printf("Total quotes: %d\n", total)
	fmt.Printf("Favorites: %d\n", favCount)
	if len(cats) > 0 {
		fmt.Println("\nBy category:")
		for _, c := range cats {
			fmt.Printf("  %-20s %d\n", c.Name+":", c.Cnt)
		}
	}
	return nil
}

func exportJSON(db *sql.DB) error {
	rows, err := db.Query("SELECT id, text, author, source, category, favorite, created_at FROM quotes ORDER BY id")
	if err != nil {
		return err
	}
	defer rows.Close()

	var quotes []Quote
	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.Author, &q.Source, &q.Category, &q.Favorite, &q.CreatedAt); err != nil {
			return err
		}
		quotes = append(quotes, q)
	}

	data, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func printUsage() {
	fmt.Println(`quote-collect — Collect and organize quotes

Usage:
  quote-collect add <text> [--author NAME] [--source SOURCE] [--category CAT]
  quote-collect list [--category CAT] [--sort id|author|date|random] [--fav]
  quote-collect show <id>
  quote-collect random
  quote-collect fav <id>           Toggle favorite
  quote-collect remove <id>
  quote-collect search <query>
  quote-collect stats
  quote-collect export [--json]

Examples:
  quote-collect add "The only way to do great work is to love what you do" --author "Steve Jobs" --category inspiration
  quote-collect add "Talk is cheap. Show me the code." --author "Linus Torvalds" --category programming
  quote-collect list
  quote-collect list --category inspiration
  quote-collect list --sort random
  quote-collect list --fav
  quote-collect random
  quote-collect fav 3
  quote-collect search code
  quote-collect stats
  quote-collect export --json

Data stored in ~/.quote-collect/quotes.db`)
}

func main() {
	_ = rand.Seed // silence unused warning if any

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	db, err := openDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: quote-collect add <text> [--author NAME] [--source SOURCE] [--category CAT]")
			os.Exit(1)
		}
		text := os.Args[2]
		author := "Unknown"
		source := ""
		category := "general"
		i := 3
		for i < len(os.Args) {
			switch os.Args[i] {
			case "--author":
				if i+1 < len(os.Args) {
					author = os.Args[i+1]
					i += 2
				}
			case "--source":
				if i+1 < len(os.Args) {
					source = os.Args[i+1]
					i += 2
				}
			case "--category":
				if i+1 < len(os.Args) {
					category = os.Args[i+1]
					i += 2
				}
			default:
				i++
			}
		}
		if err := add(db, text, author, source, category); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "list":
		category := ""
		sortBy := "id"
		favOnly := false
		for i := 2; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "--category":
				if i+1 < len(os.Args) {
					category = os.Args[i+1]
					i++
				}
			case "--sort":
				if i+1 < len(os.Args) {
					sortBy = os.Args[i+1]
					i++
				}
			case "--fav":
				favOnly = true
			}
		}
		if err := listQuotes(db, category, sortBy, favOnly); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "show":
		if len(os.Args) < 3 {
			fmt.Println("Usage: quote-collect show <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %s\n", os.Args[2])
			os.Exit(1)
		}
		if err := showQuote(db, id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "random":
		if err := randomQuote(db); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "fav":
		if len(os.Args) < 3 {
			fmt.Println("Usage: quote-collect fav <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %s\n", os.Args[2])
			os.Exit(1)
		}
		if err := toggleFav(db, id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: quote-collect remove <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid ID: %s\n", os.Args[2])
			os.Exit(1)
		}
		if err := remove(db, id); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: quote-collect search <query>")
			os.Exit(1)
		}
		if err := search(db, strings.Join(os.Args[2:], " ")); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "stats":
		if err := showStats(db); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "export":
		if err := exportJSON(db); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

	case "help", "--help", "-h":
		printUsage()

	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}
