package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
)

const AlphabetLength = 26

func getKey() string {
	key := ""
	fmt.Print("Enter the key: ")
	fmt.Scanf("%s", &key)

	return key
}

func getPlaintext() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter plaintext: ")
	plaintext, _ := reader.ReadString('\n')

	return plaintext
}

func isKeyContainsOnlyLetters(key string) bool {
	for i := 0; i < len(key); i++ {
		if !unicode.IsLetter(rune(key[i])) {
			fmt.Println("Keyword must only contains letters")
			return false
		}
	}

	return true
}

func vigenereCiphering(plaintext string, key string) {
	j := 0
	keyLength := len(key)
	ciphertext := []byte(plaintext)

	for i := 0; i < len(plaintext); i++ {
		// If the current character is letter - cipher this letter
		if unicode.IsLetter(rune(plaintext[i])) {
			ciphertext[i] = letterCiphering(plaintext[i], key[j % keyLength])
			j++
		}
	}

	fmt.Println("Ciphertext: ", string(ciphertext))
}

func letterCiphering(letter byte, key byte) byte {
	cipherletter := letter

	if unicode.IsUpper(rune(letter)) {
		// If letter and key are in upper case, else letter in upper but key in lower
		if unicode.IsUpper(rune(key)) {
			cipherletter = ((letter - 'A' + (key - 'A')) % AlphabetLength) + 'A'
		} else {
			cipherletter = ((letter - 'A' + (key - 'a')) % AlphabetLength) + 'A'
		}
	} else {
		// If letter in lower but key in upper case, else letter and key are in lower
		if unicode.IsUpper(rune(key)) {
			cipherletter = ((letter - 'a' + (key - 'A')) % AlphabetLength) + 'a'
		} else {
			cipherletter = ((letter - 'a' + (key - 'a')) % AlphabetLength) + 'a'
		}
	}

	return cipherletter
}

func main() {
	// Gets key from user
	key := getKey()
	// Checks if this key contains letters, if not - stops executing
	if !isKeyContainsOnlyLetters(key) {
		return
	}
	// Ciphers the plaintext
	vigenereCiphering(getPlaintext(), key)
}
