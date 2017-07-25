package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"palind"
)

var delimiter string
var allMustBeValid, ignoreWhite, ignorePunc bool

func init() {
	flag.StringVar(&delimiter, "d", " ", "define the delimiter between results")
	flag.BoolVar(&allMustBeValid, "allValid", false,
		"program will fail if any argument is invalid")
	flag.BoolVar(&ignorePunc, "ignorePunc", false,
		"the program will ignore punctuation marks i.e. '!' or ':'")
	flag.BoolVar(&ignoreWhite, "ignoreWhite", false,
		"the program will ignore white space including tabs and returns")
}

func main() {
	flag.Parse()
	args := flag.Args()

	// report usage if no arguments
	if len(args) == 0 {
		fmt.Println("ispal takes one or more strings as arguments and checks if they are palindromes")
		return
	}

	// initalise variables used in the following loop
	results := make([]string, len(args))
	var result bool
	var err error

	for i, testStr := range args {
		result, err = palind.IsPalindrome(testStr, ignoreWhite, ignorePunc)
		if err != nil {
			if allMustBeValid {
				fmt.Fprintf(os.Stderr, "Argument %d => %s\n", i, err.Error())
				return
			}
			results[i] = err.Error()
		} else {
			results[i] = fmt.Sprintf("%t", result)
		}
	}

	// return results to STDOUT
	fmt.Println(strings.Join(results, delimiter))
	return
}
