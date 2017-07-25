package palind

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	res, err := IsPalindrome("Level", false, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_Fail(t *testing.T) {
	res, err := IsPalindrome("palindrome", false, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != false {
		t.Error("Test should return 'false'")
	}
}

func TestIsPalindrome_InvalidInput(t *testing.T) {
	_, err := IsPalindrome("gr8", false, false)
	if err == nil {
		t.Error("Test should reject the '8' charachter")
	}
}

func TestIsPalindrome_IgnoreWhiteSpace(t *testing.T) {
	res, err := IsPalindrome("\tWas it a car or a cat I saw\n", true, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_IgnorePunctuation(t *testing.T) {
	res, err := IsPalindrome("race,car!", false, true)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_IgnoreWhiteSpaceAndPunctuation(t *testing.T) {
	res, err := IsPalindrome("A man, a plan, a canal, Panama!", true, true)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestRemoveChars(t *testing.T) {
	testStr := "Remove all m's and spaces"
	want := "Reoveall'sandspaces"
	res := removeChars(testStr, []string{"m", " "})
	if strings.Compare(want, res) != 0 {
		t.Errorf("Output not as expected.\nWanted:%s\nGot:%s", want, res)
	}
}

func TestIsBasicLatin(t *testing.T) {
	var err error
	err = isBasicLatin("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}

	if err = isBasicLatin("thenumber6"); err.Error() != "invalidChar:'6'" {
		t.Errorf("Wrong Error: %s", err.Error())
	}
	if err = isBasicLatin("somePunc#"); err.Error() != "invalidChar:'#'" {
		t.Errorf("Wrong Error: %s", err.Error())
	}
}
