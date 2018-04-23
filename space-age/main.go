package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// input must be in seconds old
	fmt.Println("Give me your age in seconds:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// read input
	text := scanner.Text()
	x, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return
	}
	calc(x)

}

func calc(input float64) {
	// earth days are 365.25
	earthDays := 365.25
	// convert seconds to days (60 * 60 * 24) as float
	sectoDays := 86400.0

	conversion1 := input / sectoDays
	conversion2 := conversion1 / earthDays

	// other planets
	mercury := conversion2 / 0.2408467
	venus := conversion2 / 0.61519726
	mars := conversion2 / 1.8808158
	jupiter := conversion2 / 11.862615
	saturn := conversion2 / 29.447498
	uranus := conversion2 / 84.016846
	neptune := conversion2 / 164.79132

	fmt.Printf("You are %.2f earth years old.\n", conversion2)
	fmt.Printf("You are %.2f mercury years old.\n", mercury)
	fmt.Printf("You are %.2f venus years old.\n", venus)
	fmt.Printf("You are %.2f mars years old.\n", mars)
	fmt.Printf("You are %.2f jupiter years old.\n", jupiter)
	fmt.Printf("You are %.2f saturn years old.\n", saturn)
	fmt.Printf("You are %.2f uranus years old.\n", uranus)
	fmt.Printf("You are %.2f neptune years old.\n", neptune)
	return
}
