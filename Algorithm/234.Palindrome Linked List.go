package main

import (
	"fmt"
	//"math/rand"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversed(head *ListNode) *ListNode {
	var current *ListNode = head
	var temp *ListNode = nil
	var last *ListNode = nil
	for current != nil {
		if current == head {
			temp = current.Next
			current.Next = nil
			last = head
			current = temp
		} else {
			temp = current.Next
			current.Next = last
			last = current
			current = temp
		}
	}
	//var current1 *ListNode = last
	//for current1 != nil {
	//	//fmt.Printf("%d ", current1.Val)
	//	current1 = current1.Next
	//}
	////fmt.Println()
	return last
}
func isPalindrome(head *ListNode) bool {
	var count int
	current := head
	for i := 0; current != nil; i++ {
		count += 1
		current = current.Next
	}
	if count == 0 {
		return true
	}
	//mt.Printf("count=%d\n", count)
	// if isParlinerome the length must be odd

	var front_half_head *ListNode
	front_half_head = nil
	tempNode := head
	current1 := head
	for i := 0; i < count/2; i++ {
		if front_half_head == nil {
			front_half_head = head
			current1 = current1.Next
			tempNode = current1
			front_half_head.Next = tempNode
		} else {
			current1 = current1.Next
			tempNode.Next = current1
			tempNode = current1
		}
	}
	if count%2 != 0 {
		current1 = current1.Next
	}
	var latter_half_head *ListNode
	latter_half_head = nil
	tempNode = current1
	for current1 != nil {
		if latter_half_head == nil {
			latter_half_head = tempNode
			tempNode = current1.Next
			current1 = current1.Next
			latter_half_head.Next = tempNode
		} else {
			current1 = current1.Next
			tempNode.Next = current1
			tempNode = current1
		}
	}

	//fmt.Println("\n\nfront_half:")
	//tempNode1 := front_half_head
	//for i := 0; i < count/2; i++ {
	//	fmt.Printf("%d ", tempNode1.Val)
	//	tempNode1 = tempNode1.Next
	//}
	//fmt.Println()
	//fmt.Println("\nlatter_half")
	//tempNode1 = latter_half_head
	//for tempNode1 != nil {
	//	fmt.Printf("%d ", tempNode1.Val)
	//	tempNode1 = tempNode1.Next
	//}
	//fmt.Println()
	latter_half_head = reversed(latter_half_head)
	node1 := front_half_head
	node2 := latter_half_head
	//fmt.Printf("node2=%v", node2)
	for i := 0; i < count/2; i++ {
		if node1.Val != node2.Val {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true

}

func main() {
	head := &ListNode{1, &ListNode{1, nil}}
	isPalindrome(head)
}
