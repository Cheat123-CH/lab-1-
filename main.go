package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"yourmodugo mod initle/utils/crack"
)

func usageAndExit() {
	fmt.Println("Usage: go run main.go <wordlist-file> <target-sha1>")
	fmt.Println("Example: go run main.go wordlists/nord_vpn.txt aa1c7d931cf140bb35a5a16adeb83a551649c3b9")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		usageAndExit()
	}

	wordlistPath := os.Args[1]
	target := strings.ToLower(strings.TrimSpace(os.Args[2]))

	// create/open verbose log file in project root: verbose.txt
	verboseFilePath := "verbose.txt"
	vf, err := os.OpenFile(verboseFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open verbose file: %v\n", err)
		os.Exit(2)
	}
	defer vf.Close()

	// Setup logging to both stdout and file
	mw := ioMultiWriter(os.Stdout, vf)
	logger := log.New(mw, "", log.LstdFlags)

	// record start
	start := time.Now()
	logger.Printf("Starting SHA1 cracking. wordlist=%s target=%s\n", filepath.Base(wordlistPath), target)

	f, err := os.Open(wordlistPath)
	if err != nil {
		logger.Fatalf("Failed to open wordlist: %v\n", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNo := 0
	var foundWord string
	for scanner.Scan() {
		lineNo++
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}

		hash := crack.HashSHA1(word)
		// verbose log each attempt
		logger.Printf("Line %d: trying word='%s' sha1=%s\n", lineNo, word, hash)

		if hash == target {
			foundWord = word
			logger.Printf("== MATCH FOUND == line=%d word='%s' sha1=%s\n", lineNo, word, hash)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		logger.Printf("Error reading wordlist: %v\n", err)
	}

	elapsed := time.Since(start)
	if foundWord != "" {
		logger.Printf("Completed in %s — password: '%s'\n", elapsed.String(), foundWord)
		fmt.Printf("\nFOUND: %s\n", foundWord)
	} else {
		logger.Printf("Completed in %s — password not found in the provided wordlist.\n", elapsed.String())
		fmt.Println("\nNOT FOUND in wordlist.")
	}
}

// ioMultiWriter avoids importing io.MultiWriter directly to show the implementation
// (but you can replace this with io.MultiWriter(os.Stdout, vf) if you prefer).
func ioMultiWriter(w1, w2 *os.File) *multiWriter {
	return &multiWriter{w1: w1, w2: w2}
}

type multiWriter struct {
	w1 *os.File
	w2 *os.File
}

func (m *multiWriter) Write(p []byte) (n int, err error) {
	n, err = m.w1.Write(p)
	if err != nil {
		return
	}
	_, err = m.w2.Write(p)
	return
}
