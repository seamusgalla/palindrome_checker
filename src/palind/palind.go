package palind

import (
	"fmt"
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

	// Convert to runes i.e. unicode code points. Allows extened charachter sets
	// and averts the danger of more than one byte representing a charachter.
	testRunes := []rune(testStr)

	if err := isLatin(testStr); err != nil {
		return false, err
	}

	// decide on comapare function
	var compare compFunc = caseInsensitiveCompare
	if caseSen {
		compare = caseSensitiveCompare
	}

	for i, j := 0, len(testRunes)-1; i < j; {
		if !compare(testRunes[i], testRunes[j]) {
			return false, nil
		}
		i++
		j--
	}
	return true, nil
}

type compFunc func(rune, rune) bool

func caseSensitiveCompare(char1, char2 rune) bool {
	return char1 == char2
}

func caseInsensitiveCompare(char1, char2 rune) bool {
	return ensureCaps(char1) == ensureCaps(char2)
}

func ensureCaps(char rune) rune {
	if char >= 'a' && char <= 'z' || (char >= 'à' && char <= 'þ') {
		return char - capsByteDiff
	}
	return char
}

// Checks for the upper and lowercase latin alphabet
func isLatin(testStr string) error {
	for _, char := range testStr {
		if !(isBasicLatinChar(char) || isExtendedLatinChar(char)) {
			return fmt.Errorf("invalidChar:'%s'", string(char))
		}
	}
	return nil
}

func isBasicLatinChar(char rune) bool {
	return ((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z'))
}

func isExtendedLatinChar(char rune) bool {
	return char >= 'À' && char <= 'ÿ' && char != '×' && char != '÷'
}

// removes any occurances of the array of chars from the string str
func removeChars(str string, chars []string) string {
	for _, char := range chars {
		str = replace(str, char, "")
	}
	return str
}

// custom replace written for demo purposes. Should use strings.Replace
func replace(str, find, replS string) string {
	for i := 0; i < len(str); i++ {
		if str[i] == find[0] {
			// only check if equal if there are enough remaining bytes to compare to.
			if len(str[i:]) >= len(find) && str[i:i+len(find)] == find {
				str = str[:i] + replS + str[i+len(find):]
				i += len(replS) - 1
			}
		}
	}
	return str
}

// No longer used after benchmarks showed it slower than current implementation
// Kept around for benchmarks.
func caseInsensitiveCompareOld(char1, char2 rune) bool {
	if char1 >= 'A' && char1 <= 'Z' {
		return char1 == char2 || char1 == (char2-capsByteDiff)
	}
	return char1 == char2 || char1 == (char2+capsByteDiff)
}
