// file-signature verifies file integrity using cryptographic signatures.
//
// Supports Ed25519 and RSA keys for signing and verification.
// Usage: file-signature <command> [options]
//
// Commands:
//   generate  Generate a new key pair (Ed25519 or RSA)
//   sign      Sign a file with a private key
//   verify    Verify a file signature against a public key
//   list      List available commands and options
package main

import (
	"crypto"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "generate":
		cmdGenerate(os.Args[2:])
	case "sign":
		cmdSign(os.Args[2:])
	case "verify":
		cmdVerify(os.Args[2:])
	case "list", "-h", "--help", "help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("file-signature - File signature verification tool")
	fmt.Println()
	fmt.Println("Usage: file-signature <command> [options]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  generate <key-type> <output-dir>  Generate a new key pair")
	fmt.Println("                                    key-type: ed25519 (default) or rsa")
	fmt.Println("  sign <file> <private-key> [output] Sign a file with a private key")
	fmt.Println("                                      output: signature file (default: file.sig)")
	fmt.Println("  verify <file> <signature> <public-key>  Verify a file signature")
	fmt.Println("  list                          List available commands")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  file-signature generate ed25519 ./keys")
	fmt.Println("  file-signature generate rsa ./keys")
	fmt.Println("  file-signature sign myfile.tar keys/private.pem")
	fmt.Println("  file-signature verify myfile.tar myfile.tar.sig keys/public.pem")
}

func cmdGenerate(args []string) {
	keyType := "ed25519"
	outputDir := "."

	if len(args) > 0 {
		keyType = args[0]
	}
	if len(args) > 1 {
		outputDir = args[1]
	}

	keyType = strings.ToLower(keyType)
	if keyType != "ed25519" && keyType != "rsa" {
		fmt.Fprintf(os.Stderr, "Error: key type must be 'ed25519' or 'rsa', got '%s'\n", keyType)
		os.Exit(1)
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot create directory %s: %v\n", outputDir, err)
		os.Exit(1)
	}

	privPath := filepath.Join(outputDir, "private.pem")
	pubPath := filepath.Join(outputDir, "public.pem")

	var privPEM, pubPEM []byte
	var bitSize int

	if keyType == "ed25519" {
		pub, priv, err := ed25519.GenerateKey(rand.Reader)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to generate Ed25519 key: %v\n", err)
			os.Exit(1)
		}

		privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to marshal private key: %v\n", err)
			os.Exit(1)
		}
		privPEM = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})

		pubBytes, err := x509.MarshalPKIXPublicKey(pub)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to marshal public key: %v\n", err)
			os.Exit(1)
		}
		pubPEM = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
	} else {
		bitSize = 2048
		if len(args) > 2 {
			fmt.Fscanf(strings.NewReader(args[2]), "%d", &bitSize)
		}
		if bitSize < 1024 || bitSize > 4096 {
			fmt.Fprintf(os.Stderr, "Error: RSA key size must be between 1024 and 4096\n")
			os.Exit(1)
		}

		priv, err := rsa.GenerateKey(rand.Reader, bitSize)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to generate RSA key: %v\n", err)
			os.Exit(1)
		}

		privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to marshal private key: %v\n", err)
			os.Exit(1)
		}
		privPEM = pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})

		pubBytes, err := x509.MarshalPKIXPublicKey(&priv.PublicKey)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to marshal public key: %v\n", err)
			os.Exit(1)
		}
		pubPEM = pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes})
	}

	if err := os.WriteFile(privPath, privPEM, 0600); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot write private key: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(pubPath, pubPEM, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot write public key: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %s key pair\n", strings.ToUpper(keyType))
	fmt.Printf("  Private key: %s\n", privPath)
	fmt.Printf("  Public key:  %s\n", pubPath)
}

func cmdSign(args []string) {
	if len(args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: sign requires <file> <private-key> [output]\n")
		os.Exit(1)
	}

	filePath := args[0]
	keyPath := args[1]
	sigPath := filePath + ".sig"
	if len(args) > 2 {
		sigPath = args[2]
	}

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot open file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot read file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	hashed := hasher.Sum(nil)

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot read key %s: %v\n", keyPath, err)
		os.Exit(1)
	}

	privKey, err := parsePrivateKey(keyData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid private key: %v\n", err)
		os.Exit(1)
	}

	sig, err := signData(privKey, hashed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: signing failed: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(sigPath, sig, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot write signature: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Signed: %s\n", filePath)
	fmt.Printf("Signature: %s\n", sigPath)
	fmt.Printf("Algorithm: SHA256 with %s\n", keyTypeName(privKey))
}

func cmdVerify(args []string) {
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Error: verify requires <file> <signature> <public-key>\n")
		os.Exit(1)
	}

	filePath := args[0]
	sigPath := args[1]
	keyPath := args[2]

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot open file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot read file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	hashed := hasher.Sum(nil)

	sigData, err := os.ReadFile(sigPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot read signature %s: %v\n", sigPath, err)
		os.Exit(1)
	}

	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: cannot read key %s: %v\n", keyPath, err)
		os.Exit(1)
	}

	pubKey, err := parsePublicKey(keyData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid public key: %v\n", err)
		os.Exit(1)
	}

	valid := verifyData(pubKey, hashed, sigData)
	if valid {
		fmt.Println("OK: signature is valid")
		fmt.Printf("  File: %s\n", filePath)
		fmt.Printf("  Signature: %s\n", sigPath)
		fmt.Printf("  Algorithm: SHA256 with %s\n", keyTypeName(pubKey))
	} else {
		fmt.Println("FAIL: signature verification failed")
		fmt.Printf("  File: %s\n", filePath)
		fmt.Printf("  Signature: %s\n", sigPath)
		os.Exit(2)
	}
}

func parsePrivateKey(data []byte) (crypto.Signer, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM block found")
	}

	if block.Type == "PRIVATE KEY" {
		key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		switch k := key.(type) {
		case ed25519.PrivateKey:
			return k, nil
		case *rsa.PrivateKey:
			return k, nil
		default:
			return nil, fmt.Errorf("unsupported private key type")
		}
	}

	return nil, fmt.Errorf("unsupported PEM type: %s", block.Type)
}

func parsePublicKey(data []byte) (crypto.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("no PEM block found")
	}

	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("unsupported PEM type: %s", block.Type)
	}

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func signData(privKey crypto.Signer, hashed []byte) ([]byte, error) {
	switch k := privKey.(type) {
	case ed25519.PrivateKey:
		return ed25519.Sign(k, hashed), nil
	case *rsa.PrivateKey:
		return rsa.SignPSS(rand.Reader, k, crypto.SHA256, hashed, &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
			Hash:       crypto.SHA256,
		})
	default:
		return nil, fmt.Errorf("unsupported key type for signing")
	}
}

func verifyData(pubKey crypto.PublicKey, hashed []byte, sig []byte) bool {
	switch k := pubKey.(type) {
	case ed25519.PublicKey:
		return ed25519.Verify(k, hashed, sig)
	case *rsa.PublicKey:
		err := rsa.VerifyPSS(k, crypto.SHA256, hashed, sig, &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
			Hash:       crypto.SHA256,
		})
		return err == nil
	default:
		return false
	}
}

func keyTypeName(key any) string {
	switch key.(type) {
	case ed25519.PublicKey, ed25519.PrivateKey:
		return "Ed25519"
	case *rsa.PublicKey, *rsa.PrivateKey:
		return "RSA"
	default:
		return "unknown"
	}
}
