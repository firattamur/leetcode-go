package p0020

import "testing"

func TestP0020Case1(t *testing.T) {
	s := "()"
	expected := true

	result := IsValid(s)
	if result != expected {
		t.Errorf("IsValid(%s) = %t, expected %t", s, result, expected)
	}
}

func TestP0020Case2(t *testing.T) {
	s := "()[]{}"
	expected := true

	result := IsValid(s)
	if result != expected {
		t.Errorf("IsValid(%s) = %t, expected %t", s, result, expected)
	}
}

func TestP0020Case3(t *testing.T) {
	s := "(]"
	expected := false

	result := IsValid(s)
	if result != expected {
		t.Errorf("IsValid(%s) = %t, expected %t", s, result, expected)
	}
}
