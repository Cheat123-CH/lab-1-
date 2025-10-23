package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"golang.org/x/crypto/sha3"
	"encoding/hex"
)

// Function to compute hash and return hexadecimal string
func computeHash(algorithm string, input string) string {
	var hash []byte

	switch algorithm {
	case "MD5":
		sum := md5.Sum([]byte(input))
		hash = sum[:]
	case "SHA1":
		sum := sha1.Sum([]byte(input))
		hash = sum[:]
	case "SHA256":
		sum := sha256.Sum256([]byte(input))
		hash = sum[:]
	case "SHA512":
		sum := sha512.Sum512([]byte(input))
		hash = sum[:]
	case "SHA3":
		sum := sha3.Sum512([]byte(input))
		hash = sum[:]
	default:
		return ""
	}

	return hex.EncodeToString(hash)
}

func main() {
	var input1, input2 string

	fmt.Println("===== Hash Proof Program =====")
	fmt.Print("Enter first input: ")
	fmt.Scanln(&input1)
	fmt.Print("Enter second input: ")
	fmt.Scanln(&input2)

	algorithms := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3"}

	fmt.Println("\n--- Hash Comparison Results ---")
	for _, algo := range algorithms {
		hash1 := computeHash(algo, input1)
		hash2 := computeHash(algo, input2)

		fmt.Printf("\nAlgorithm: %s\n", algo)
		fmt.Printf("Input1 Hash: %s\n", hash1)
		fmt.Printf("Input2 Hash: %s\n", hash2)

		if hash1 == hash2 {
			fmt.Printf("Result: Match\n")
		} else {
			fmt.Printf("Result:Match!\n")
		}
	}
}
