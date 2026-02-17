package listnode

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

func NewListNode(data int, next *ListNode) *ListNode {
	return &ListNode{
		Data: data,
		Next: next,
	}
}

func InsertAtBack(head *ListNode, data int) *ListNode {
	newNode := NewListNode(data, nil)

	if head == nil {
		return newNode
	}
	current := head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return head
}

func InsertAtFront(head *ListNode, data int) *ListNode {
	newHead := NewListNode(data, head)
	return newHead
}

func DeleteNode(head *ListNode, data int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Data == data {
		return head.Next
	}
	current := head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return head
		}
		current = current.Next
	}
	return head
}

func TraverseList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Print(current.Data, " ")
		current = current.Next
	}
	fmt.Println()
}
