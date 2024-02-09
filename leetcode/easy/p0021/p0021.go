package p0021

/*

Leetcode Problem 21: Merge Two Sorted Lists

https://leetcode.com/problems/merge-two-sorted-lists/

*/

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	currentHead := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			currentHead.Next = l1
			l1 = l1.Next
		} else {
			currentHead.Next = l2
			l2 = l2.Next
		}
		currentHead = currentHead.Next
	}

	if l1 != nil {
		currentHead.Next = l1
	} else {
		currentHead.Next = l2
	}

	return head.Next
}
