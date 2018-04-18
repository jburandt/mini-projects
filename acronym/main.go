package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Give me an acronym to generate")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// read input
	text := scanner.Text()
	// split strings by space " "
	words := strings.Split(text, " ")

	res := ""

	for _, word := range words {
		// add first character of each string to slice
		res = res + string([]rune(word)[0])
	}

	fmt.Println(res)
}
