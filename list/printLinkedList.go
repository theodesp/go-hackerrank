package list

import "fmt"

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode {
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

/*
If you're new to linked lists, this is a great exercise for learning about them.
Given a pointer to the head node of a linked list, print its elements in order,
one element per line. If the head pointer is null (indicating the list is empty),
don’t print anything.

Input Format

The first line of input contains , the number of elements in the linked list.
The next  lines contain one element each, which are the elements of the linked list.

Note: Do not read any input from stdin/console. Complete the printLinkedList function
in the editor below.
 */

func PrintLinkedList(head *SinglyLinkedListNode) {
	if head.next == nil {
		fmt.Println(head.data)
		return
	}
	for head != nil {
		fmt.Println(head.data)
		head = head.next
	}
}
