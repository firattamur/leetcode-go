package p0014

import "testing"

func TestP0014Case1(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	expected := "fl"

	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("LongestCommonPrefix(%v) = %s, expected %s", strs, result, expected)
	}
}

func TestP0014Case2(t *testing.T) {
	strs := []string{"dog", "racecar", "car"}
	expected := ""

	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("LongestCommonPrefix(%v) = %s, expected %s", strs, result, expected)
	}
}

func TestP0014Case3(t *testing.T) {
	strs := []string{"", "racecar", "car"}
	expected := ""

	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("LongestCommonPrefix(%v) = %s, expected %s", strs, result, expected)
	}
}

func TestP0014Case4(t *testing.T) {
	strs := []string{}
	expected := ""

	result := LongestCommonPrefix(strs)
	if result != expected {
		t.Errorf("LongestCommonPrefix(%v) = %s, expected %s", strs, result, expected)
	}
}
