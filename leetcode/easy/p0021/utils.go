package p0021

type ListNode struct {
	Val  int
	Next *ListNode
}

func ArrayToListNode(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  values[0],
		Next: nil,
	}

	current := head
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{
			Val:  values[i],
			Next: nil,
		}
		current = current.Next
	}

	return head
}
