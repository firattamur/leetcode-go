package leetcode

import (
	"testing"
)

func TestP0009Case1(t *testing.T) {
	x := 121
	expected := true

	result := IsPalindrome(x)
	if result != expected {
		t.Errorf("IsPalindrome(%d) = %t, expected %t", x, result, expected)
	}
}

func TestP0009Case2(t *testing.T) {
	x := -121
	expected := false

	result := IsPalindrome(x)
	if result != expected {
		t.Errorf("IsPalindrome(%d) = %t, expected %t", x, result, expected)
	}
}

func TestP0009Case3(t *testing.T) {
	x := 10
	expected := false

	result := IsPalindrome(x)
	if result != expected {
		t.Errorf("IsPalindrome(%d) = %t, expected %t", x, result, expected)
	}
}

