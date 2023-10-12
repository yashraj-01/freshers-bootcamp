package main

import (
	"fmt"
	"unicode"
)

type Stack[T any] []T

func (s *Stack[T]) push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) pop() T {
	l := len(*s)
	if l == 0 {
		var res T
		return res
	}
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

func (s *Stack[T]) top() T {
	l := len(*s)
	return (*s)[l-1]
}

func (s *Stack[T]) isEmpty() bool {
	l := len(*s)
	return l == 0
}

type Node struct {
	val   int32
	left  *Node
	right *Node
}

// Returns true if op1 is a operator and priority(op1) >= priority(op2)
func compare(op1, op2 int32) bool {
	if op1 == '(' || op1 == ')' {
		return false
	}
	return op1 == '*' || op1 == '/' || op2 == '+' || op2 == '-'
}

func genExpTree(s string) *Node {
	var stNodes Stack[*Node]
	var stOps Stack[int32]

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			stNodes.push(&Node{ch, nil, nil})
		} else if ch == '(' {
			stOps.push(ch)
		} else if ch == ')' {
			for stOps.top() != '(' {
				right := stNodes.pop()
				left := stNodes.pop()
				op := stOps.pop()
				stNodes.push(&Node{op, left, right})
			}
			stOps.pop()
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			for !stOps.isEmpty() && compare(stOps.top(), ch) {
				right := stNodes.pop()
				left := stNodes.pop()
				op := stOps.pop()
				stNodes.push(&Node{op, left, right})
			}
			stOps.push(ch)
		}
	}

	for !stOps.isEmpty() {
		right := stNodes.pop()
		left := stNodes.pop()
		op := stOps.pop()
		stNodes.push(&Node{op, left, right})
	}

	return stNodes.top()
}

func main() {
	s := "1+2+3+4+5"
	expTree := genExpTree(s)
	fmt.Println("In-order traversal:")
	inOrder(expTree)
	fmt.Println("\nPost-order traversal:")
	postOrder(expTree)
	fmt.Println()
}

func postOrder(root *Node) {
	if root == nil {
		return
	}
	postOrder(root.left)
	postOrder(root.right)
	fmt.Printf("%c ", root.val)
}

func inOrder(root *Node) {
	if root == nil {
		return
	}
	inOrder(root.left)
	fmt.Printf("%c ", root.val)
	inOrder(root.right)
}
