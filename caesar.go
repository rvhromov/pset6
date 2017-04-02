package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

const AlphabetLength = 26

func getKey() int {
	key := 0
	fmt.Print("Enter the key: ")
	fmt.Scanf("%d", &key)

	return key
}

func checkKey(key int) bool {
	if key < 0 {
		fmt.Println("Key must be positive integer.")
		return false
	}

	return true
}

func getPlaintext() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter plaintext: ")
	plaintext, _ := reader.ReadString('\n')

	return plaintext
}

func caeserCiphering(plaintext string, key int) {
	ciphertext := []byte(plaintext)

	for i := 0; i < len(plaintext); i++ {
		if unicode.IsLetter(rune(plaintext[i])) {
			if unicode.IsUpper(rune(plaintext[i])) {
				ciphertext[i] = ((plaintext[i] + byte(key) - 'A') % AlphabetLength) + 'A'
			} else {
				ciphertext[i] = ((plaintext[i] + byte(key) - 'a') % AlphabetLength) + 'a'
			}
		}
	}

	fmt.Println("Ciphertext: ", string(ciphertext));
}

func main() {
	// Gets key from user
	key := getKey()
	// Checks if this key is positive, if not - stops executing
	if !checkKey(key) {
		return
	}
	// Ciphers the plaintext
	caeserCiphering(getPlaintext(), key)
}
