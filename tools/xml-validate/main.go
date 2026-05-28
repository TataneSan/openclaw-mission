package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"os"
)

func run() error {
	input := flag.String("i", "", "chemin du fichier XML (vide pour stdin)")
	pretty := flag.Bool("p", false, "affiche le XML formate si valide")
	flag.Parse()

	var reader io.Reader
	if *input != "" {
		f, err := os.Open(*input)
		if err != nil {
			return fmt.Errorf("erreur ouverture fichier: %w", err)
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("erreur lecture: %w", err)
	}

	// Validation XML via Unmarshal
	var raw interface{}
	if err := xml.Unmarshal(data, &raw); err != nil {
		fmt.Fprintf(os.Stderr, "ERREUR: %v\n", err)
		fmt.Fprintf(os.Stderr, "Validation echouee\n")
		os.Exit(1)
		return nil
	}

	fmt.Fprintln(os.Stderr, "Validation reussie: XML valide")

	if *pretty {
		pretty, _ := xml.MarshalIndent(raw, "", "  ")
		fmt.Println(string(pretty))
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "xml-validate: %v\n", err)
		os.Exit(1)
	}
}
