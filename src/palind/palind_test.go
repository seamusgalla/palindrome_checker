package palind

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	res, err := IsPalindrome("Level", false, false, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_Fail(t *testing.T) {
	res, err := IsPalindrome("palindrome", false, false, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != false {
		t.Error("Test should return 'false'")
	}
}

func TestIsPalindrome_InvalidInput(t *testing.T) {
	_, err := IsPalindrome("gr8", false, false, false)
	if err == nil {
		t.Error("Test should reject the '8' charachter")
	}
}

func TestIsPalindrome_IgnoreWhiteSpace(t *testing.T) {
	res, err := IsPalindrome("\tWas it a car or a cat I saw\n", true, false, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_IgnorePunctuation(t *testing.T) {
	res, err := IsPalindrome("race,car!", false, true, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_IgnoreWhiteSpaceAndPunctuation(t *testing.T) {
	res, err := IsPalindrome("A man, a plan, a canal, Panama!", true, true, false)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != true {
		t.Error("Test should return 'true'")
	}
}

func TestIsPalindrome_CaseSensitive(t *testing.T) {
	res, err := IsPalindrome("Kayak", true, true, true)
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	if res != false {
		t.Error("Test should return 'false'")
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

func TestIsLatin(t *testing.T) {
	var err error
	var want string
	err = isLatin("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		t.Errorf("Unexpected Error: %s", err.Error())
	}
	err = isLatin("thenumber6")
	want = "invalidChar:'6'"
	if err.Error() != want {
		t.Errorf("Wrong Error: %s\nWant: %s", err.Error(), want)
	}
	err = isLatin("somePunc#")
	want = "invalidChar:'#'"
	if err.Error() != want {
		t.Errorf("Wrong Error: %s\nWant: %s", err.Error(), want)
	}
}

func TestReplace(t *testing.T) {
	var want = "My dog"
	var res = replace("My cat", "cat", "dog")
	if res != want {
		t.Errorf("Wrong Result: %s\nWant: %s", res, want)
	}
	want = "My tall cat"
	res = replace("My tiny cat", "tiny", "tall")
	if res != want {
		t.Errorf("Wrong Result: %s\nWant: %s", res, want)
	}
	want = "My very very tiny cat"
	res = replace("My really really tiny cat", "really", "very")
	if res != want {
		t.Errorf("Wrong Result: %s\nWant: %s", res, want)
	}
}

func BenchmarkCaseInsensitiveCompareOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		caseInsensitiveCompareOld('b', 'B')
		caseInsensitiveCompareOld('P', 'a')
		caseInsensitiveCompareOld('x', 'x')
	}
}

func BenchmarkCaseInsensitiveCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		caseInsensitiveCompare('b', 'B')
		caseInsensitiveCompare('P', 'a')
		caseInsensitiveCompare('x', 'x')
	}
}

func BenchmarkCustomReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		replace("My really really tall cat", "really", "very")
	}
}

func BenchmarkStandardLibraryReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Replace("My really really tall cat", "really", "very", -1)
	}
}
