package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	// input must be in seconds old
	fmt.Println("Give me your age in seconds <TEST> use 866000000")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// read input
	text := scanner.Text()
	// split strings by space " "
	x, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return
	}
	earth(x)
	
}


func earth(input float64) {
	// earth days are 365.25
	earthDays := 365.25
	// convert seconds to days (60 * 60 * 24) as float
	sectoDays := 86400.0

	conversion1 := input / sectoDays
	conversion2 := conversion1 / earthDays
	fmt.Printf("You are about %.2f earth years old.\n", conversion2)
	return
}