package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Hey, I'm Bob. What do you want?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	if len(text) == 0 {
		fmt.Println("Fine. Be that way!")
	}else {
		conversation(text)
	}
}

func conversation(input string) {
	inputParse := input[len(input)-1:]
	inputParseTwo := input[len(input)-2:]

	if inputParseTwo == "?!" || inputParseTwo == "!?" {
		fmt.Println("Calm down, I know what I'm doing!")
		return
	}

	switch inputParse {
	case "?":
		fmt.Println("Sure.")
	case "!":
		fmt.Println("Whoa, chill out!")
	default:
		fmt.Println("Whatever")
	}

	return
}