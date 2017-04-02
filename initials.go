package main

import (
	"fmt"
	"strings"
	"unicode"
	"os"
	"bufio"
)

func getName() string {
	// Prompt and read user's input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	return strings.ToUpper(name)
}

func showInitials(name string) {
	for i := 0; i < len(name); i++ {
		// If current character is letter
		if unicode.IsLetter(rune(name[i])) {
			// Show it
			fmt.Printf("%c", name[i])
			// And keep iterate until the end of the word
			for unicode.IsLetter(rune(name[i + 1])) {
				i++;
			}
		}
	}
}

func main() {
	showInitials(getName())
}
