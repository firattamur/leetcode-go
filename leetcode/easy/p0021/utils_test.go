package p0021

import (
	"reflect"
	"testing"
)

func TestArrayToListNodeWithNotEmptyValues(t *testing.T) {
	values := []int{1, 2, 3}
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	result := ArrayToListNode(values)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ArrayToListNode(%v) = %v, expected %v", values, result, expected)
	}
}

func TestArrayToListNodeWithEmptyValues(t *testing.T) {
	values := []int{}
	expected := (*ListNode)(nil)

	result := ArrayToListNode(values)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ArrayToListNode(%v) = %v, expected %v", values, result, expected)
	}
}
