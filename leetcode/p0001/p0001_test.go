package leetcode

import (
	"reflect"
	"testing"
)

func TestP0001Case1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	result := TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TwoSum(%v, %d) = %v, expected %v", nums, target, result, expected)
	}
}

func TestP0001Case2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}

	result := TwoSum(nums, target)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TwoSum(%v, %d) = %v, expected %v", nums, target, result, expected)
	}
}

func TestP0001Case3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}

	result := TwoSum(nums, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TwoSum(%v, %d) = %v, expected %v", nums, target, result, expected)
	}
}

func TestP0001Case4(t *testing.T) {
	nums := []int{2, 3, 1}
	target := 6
	expected := []int(nil)

	result := TwoSum(nums, target)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("TwoSum(%v, %d) = %v, expected %v", nums, target, result, expected)
	}
}
