package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const capsByteDiff = 'a' - 'A'

var delimiter string
var allMustBeValid bool
var caseSensitive bool

func init() {
	flag.StringVar(&delimiter, "d", " ", "define the delimiter between results")
	flag.BoolVar(&allMustBeValid, "allValid", false,
		"program will fail if any argument is invalid")
	flag.BoolVar(&caseSensitive, "caseSen", false, "make the check case sensitive")
}

func main() {
	flag.Parse()
	inputStrings := flag.Args()

	// report if no arguments
	if len(inputStrings) == 0 {
		fmt.Println("no arguments recieved")
		return
	}

	// decide on comparisionFunction
	var compFunc func(byte, byte) bool
	if caseSensitive {
		compFunc = caseSensitiveCompare
	} else {
		compFunc = caseInsensitiveCompare
	}

	// initalise results array and populate
	results := make([]string, len(inputStrings))
	for i, testString := range inputStrings {
		if err := isBasicLatin(testString); err != nil {
			if allMustBeValid {
				fmt.Fprintf(os.Stderr, "Argument %d invalid=>%s\n", i, err.Error())
				return
			}
			results[i] = err.Error()
		} else {
			results[i] = fmt.Sprintf("%t", testPalindrome(testString,
				compFunc))
		}
	}
	fmt.Println(strings.Join(results, delimiter))
	return
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

func testPalindrome(testStr string, compareFunc func(byte, byte) bool) bool {
	for i, j := 0, len(testStr)-1; i < j; {
		if !compareFunc(testStr[i], testStr[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func caseSensitiveCompare(char1, char2 byte) bool {
	return char1 == char2
}

func caseInsensitiveCompare(char1, char2 byte) bool {
	if char1 >= 'A' && char1 <= 'Z' {
		return char1 == char2 || char1 == (char2-capsByteDiff)
	}
	return char1 == char2 || char1 == (char2+capsByteDiff)
}
