package palind

import (
	"fmt"
	"strings"
)

const (
	capsByteDiff = 'a' - 'A'
)

var (
	whiteChars = []string{" ", "\t", "\n", "\v", "\f", "\r"}
	puncChars  = []string{",", ".", "!", "?", ":", ";", "'", "\""}
)

// IsPalindrome takes a word to test for being a palindrome and a second bool
// argument stating wether or not it should be case sensitive.
func IsPalindrome(testStr string, ignoreWhite, ignorePunc, caseSen bool) (bool,
	error) {
	// determine and remove ignore set
	var ignoreSet []string
	if ignoreWhite {
		ignoreSet = append(ignoreSet, whiteChars...)
	}
	if ignorePunc {
		ignoreSet = append(ignoreSet, puncChars...)
	}
	testStr = removeChars(testStr, ignoreSet)

	if err := isBasicLatin(testStr); err != nil {
		return false, err
	}

	// decide on comapare function
	var compare compFunc = caseInsensitiveCompare
	if caseSen {
		compare = caseSensitiveCompare
	}

	for i, j := 0, len(testStr)-1; i < j; {
		if !compare(testStr[i], testStr[j]) {
			return false, nil
		}
		i++
		j--
	}
	return true, nil
}

type compFunc func(byte, byte) bool

func caseInsensitiveCompare(char1, char2 byte) bool {
	if char1 >= 'A' && char1 <= 'Z' {
		return char1 == char2 || char1 == (char2-capsByteDiff)
	}
	return char1 == char2 || char1 == (char2+capsByteDiff)
}

func caseSensitiveCompare(char1, char2 byte) bool {
	return char1 == char2
}

// Checks for the upper and lowercase latin alphabet
func isBasicLatin(testStr string) error {
	for _, charachter := range testStr {
		if charachter < 'A' || charachter > 'z' ||
			(charachter < 'a' && charachter > 'Z') {
			return fmt.Errorf("invalidChar:'%s'", string(charachter))
		}
	}
	return nil
}

// removes any occurances of the array of chars from the string str
func removeChars(str string, chars []string) string {
	for _, char := range chars {
		str = strings.Replace(str, char, "", -1)
	}
	return str
}
