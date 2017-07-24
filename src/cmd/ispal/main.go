package main

import (
	"fmt"
	"os"
	"strings"
)

var delimiter = " "

func main() {
	inputStrings := os.Args[1:]

	// report if no arguments
	if len(inputStrings) == 0 {
		fmt.Println("no arguments recieved")
		return
	}

	// initalise results array and populate
	results := make([]string, len(inputStrings))
	for i, testString := range inputStrings {
		if err := isBasicLatin(testString); err != nil {
			results[i] = err.Error()
		} else {
			results[i] = fmt.Sprintf("%t", isCaseSensitivePalindrome(testString))
		}
	}
	fmt.Println(strings.Join(results, delimiter))
}

// Checks for the upper and lowercase latin alphabet
func isBasicLatin(testStr string) error {
	for _, charachter := range testStr {
		if charachter < 'A' || charachter > 'z' ||
			(charachter < 'a' && charachter > 'Z') {
			return fmt.Errorf("invaldChar:'%s'", string(charachter))
		}
	}
	return nil
}

func isCaseSensitivePalindrome(testStr string) bool {
	for i, j := 0, len(testStr)-1; i < j; {
		if testStr[i] != testStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
