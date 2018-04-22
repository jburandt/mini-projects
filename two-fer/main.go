package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("What is your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("One for you, one for me.")
	} else {
		fmt.Printf("One for %s, one for me.\n", text)
	}
}
