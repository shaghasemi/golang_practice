package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("\n")
	scanner := bufio.NewScanner(os.Stdin)

	var codes map[string]string

	scanner.Scan()
	numberOfCountries, _ := strconv.Atoi(scanner.Text())

	codes = make(map[string]string)

	for i := 0; i < numberOfCountries; i++ {
		scanner.Scan()
		content := scanner.Text()
		parts := strings.Fields(content)
		codes[parts[1]] = parts[0]
	}

	var callCountries []string

	scanner.Scan()
	numberOfCalls, _ := strconv.Atoi(scanner.Text())

	callCountries = make([]string, numberOfCalls)

	for j := range numberOfCalls {
		scanner.Scan()
		callNumber := (scanner.Text())[0:3]
		a, ok := codes[callNumber]
		if ok {
			callCountries[j] = a
		} else {
			callCountries[j] = "Invalid Number"
		}
	}
}
