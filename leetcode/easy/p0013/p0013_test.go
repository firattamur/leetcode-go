package leetcode

import (
	"testing"
)

func TestP0013Case1(t *testing.T) {
	s := "III"
	expected := 3

	result := RomanToInt(s)
	if result != expected {
		t.Errorf("RomanToInt(%s) = %d, expected %d", s, result, expected)
	}
}

func TestP0013Case2(t *testing.T) {
	s := "IV"
	expected := 4

	result := RomanToInt(s)
	if result != expected {
		t.Errorf("RomanToInt(%s) = %d, expected %d", s, result, expected)
	}
}

func TestP0013Case3(t *testing.T) {
	s := "IX"
	expected := 9

	result := RomanToInt(s)
	if result != expected {
		t.Errorf("RomanToInt(%s) = %d, expected %d", s, result, expected)
	}
}

func TestP0013Case4(t *testing.T) {
	s := "LVIII"
	expected := 58

	result := RomanToInt(s)
	if result != expected {
		t.Errorf("RomanToInt(%s) = %d, expected %d", s, result, expected)
	}
}
