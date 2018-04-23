package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// string comparisons must be equal length
// compare at the index location of the string

func main() {
	fmt.Println("Enter your first string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("Enter your second string: ")
	scanner.Scan()
	text2 := scanner.Text()

	// verify the lengths match
	if len(text) != len(text2) {
		fmt.Println("Error: Input lengths do not match")
		return
	}

	// convert strings to upper case to make sure its a valid comparison
	text = strings.ToUpper(text)
	text2 = strings.ToUpper(text2)

	// iterate over both strings and write a counter for every mismatch
	x := 0
	for i := range text {
		if text[i] != text2[i] {
			x = x + 1
		}
	}
	fmt.Printf("The different between the two strands is %d\n", x)
}
