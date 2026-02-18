package main

import "linkedlist/listnode"

func main() {
	head := listnode.InsertAtBack(nil, 1)
	head.Next = listnode.NewListNode(2, nil)
	head.Next.Next = listnode.NewListNode(3, nil)
	head = listnode.InsertAtBack(head, 4)
	head = listnode.InsertAtBack(head, 5)

	listnode.TraverseList(head)

	head = listnode.InsertAtFront(head, 6)
	listnode.TraverseList(head)

	head = listnode.DeleteNode(head, 3)
	listnode.TraverseList(head)

	head = listnode.ReverseList(head)
	listnode.TraverseList(head)

	head = listnode.DeleteNode(head, 6)
	listnode.TraverseList(head)
}
