package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) Prepend(n *Node) *Node {
	n.Next = l.Head
	l.Head = n
	l.Length++

	return l.Head
}

func (l *LinkedList) Append(n *Node) *Node {
	curr := l.Head
	if curr == nil {
		return l.Prepend(n)
	}
	// traverse to the end
	for curr != nil && curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = n
	l.Length++

	return l.Head
}

func (l *LinkedList) InsertAtKthPos(n *Node, k int) *Node {
	if l.Head == nil {
		return l.Append(n)
	}

	if k == 1 {
		return l.Prepend(n)
	}

	curr, count := l.Head, 1
	for curr != nil && curr.Next != nil {
		curr = curr.Next
		count++
	}

	// if k > list Length
	if curr == nil || curr.Next == nil {
		return l.Append(n)
	}

	// you are in (k-1)th pos
	temp := curr.Next
	curr.Next = n
	n.Next = temp
	l.Length++

	return l.Head
}

func (l *LinkedList) InsertBeforeX(n, x int) *Node {
	nodeToBeAdded := &Node{}
	nodeToBeAdded.Val = x

	if l.Head == nil {
		return l.Prepend(nodeToBeAdded)
	}

	if l.Head.Val == n {
		return l.Prepend(nodeToBeAdded)
	}

	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == n {
			temp := curr.Next
			curr.Next = nodeToBeAdded
			nodeToBeAdded.Next = temp
			l.Length++

			return l.Head
		}
		curr = curr.Next
	}

	return nil
}

func (l *LinkedList) DeleteHead() *Node {
	if l.Head == nil {
		return nil
	}
	curr := l.Head
	l.Head = curr.Next
	curr.Next = nil
	l.Length--

	return l.Head
}

func (l *LinkedList) DeleteTail() *Node {
	if l.Head == nil {
		return nil
	}

	curr := l.Head
	// traverse to the second last Node
	for curr != nil && curr.Next != nil && curr.Next.Next != nil {
		curr = curr.Next
	}

	curr.Next = nil
	l.Length--

	return l.Head
}

func (l *LinkedList) DeleteKthNode(k int) *Node {
	if l.Head == nil {
		return nil
	}
	if k < 1 {
		return l.Head
	}
	curr, count := l.Head, 1
	for curr != nil && count < k-1 {
		curr = curr.Next
		count++
	}

	// if k > list Length
	if curr == nil || curr.Next == nil {
		return l.Head
	}

	// you are (k-1)th Node
	Next := curr.Next.Next
	curr.Next.Next = nil
	curr.Next = Next
	l.Length--

	return l.Head
}

func (l *LinkedList) DeleteNodeWithValueX(x int) *Node {
	if l.Head == nil {
		return nil
	}
	if l.Head.Val == x {
		return l.DeleteHead()
	}

	curr := l.Head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == x {
			temp := curr.Next.Next
			curr.Next.Next = nil
			curr.Next = temp
			l.Length--

			return l.Head
		}
		curr = curr.Next
	}

	return nil
}

func (l *LinkedList) print() {
	curr := l.Head

	for curr != nil {
		fmt.Printf("%v -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("NULL")
}

func main() {
	list := LinkedList{}
	Node1 := &Node{Val: 10}
	Node2 := &Node{Val: 20}
	Node3 := &Node{Val: 30}
	Node4 := &Node{Val: 40}

	list.Prepend(Node1)
	list.Prepend(Node2)
	list.Prepend(Node3)
	list.Append(Node4)

	list.print()

	list.InsertBeforeX(30, 15)
	list.print()
}
