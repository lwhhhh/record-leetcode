package algorithm

func removeElements(head *ListNode, val int) *ListNode {
	temp := &ListNode{0, nil}
	temp.Next = head
	pointer := temp

	for pointer.Next != nil {
		if pointer.Next.Val == val {
			pointer.Next = pointer.Next.Next
		} else {
			pointer = pointer.Next
		}
	}
	return temp.Next
}
