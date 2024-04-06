package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	cur, prev := head, head
	for cur != nil {
		// if the head is the value to be removed, move head to the next
		if cur == head && val == cur.Val {
			head = head.Next
			cur = head
			prev = head
			// if current val is the value to be removed, remove it
		} else if cur.Val == val {
			prev.Next = cur.Next
			cur = prev.Next
			// nothing to remove, move to the next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return head
}
