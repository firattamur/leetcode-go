package p0021

import (
	"reflect"
	"testing"
)

func TestP0021Case1(t *testing.T) {
	l1 := ArrayToListNode([]int{1, 2, 4})
	l2 := ArrayToListNode([]int{1, 3, 4})
	expected := ArrayToListNode([]int{1, 1, 2, 3, 4, 4})

	result := MergeTwoLists(l1, l2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeTwoLists(%v, %v) = %v, expected %v", l1, l2, result, expected)
	}
}

func TestP0021Case2(t *testing.T) {
	l1 := ArrayToListNode([]int{})
	l2 := ArrayToListNode([]int{})
	expected := ArrayToListNode([]int{})

	result := MergeTwoLists(l1, l2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeTwoLists(%v, %v) = %v, expected %v", l1, l2, result, expected)
	}
}

func TestP0021Case3(t *testing.T) {
	l1 := ArrayToListNode([]int{})
	l2 := ArrayToListNode([]int{0})
	expected := ArrayToListNode([]int{0})

	result := MergeTwoLists(l1, l2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeTwoLists(%v, %v) = %v, expected %v", l1, l2, result, expected)
	}
}
