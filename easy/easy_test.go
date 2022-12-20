package easy

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome(12321) {
		t.Fatal("test failed!")
	}
	t.Log("test success!")
}

func TestLongestCommonPrefix(t *testing.T) {
	print(LongestCommonPrefix([]string{"ab", "a"}))
}
