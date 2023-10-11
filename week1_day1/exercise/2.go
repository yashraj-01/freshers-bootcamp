package main

import (
	"fmt"
	"unicode"
)

type Stack []interface{}

func (s *Stack) push(x interface{}) {
	*s = append(*s, x)
}

func (s *Stack) pop() interface{} {
	l := len(*s)
	if l == 0 {
		return nil
	}
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

func (s *Stack) top() interface{} {
	l := len(*s)
	return (*s)[l-1]
}

func (s *Stack) isEmpty() bool {
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
	var stNodes Stack
	var stOps Stack

	for _, ch := range s {
		if unicode.IsDigit(ch) {
			stNodes.push(&Node{ch, nil, nil})
		} else if ch == '(' {
			stOps.push(ch)
		} else if ch == ')' {
			for stOps.top() != '(' {
				right := stNodes.pop().(*Node)
				left := stNodes.pop().(*Node)
				op := stOps.pop().(int32)
				stNodes.push(&Node{op, left, right})
			}
			stOps.pop()
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			for !stOps.isEmpty() && compare(stOps.top().(int32), ch) {
				right := stNodes.pop().(*Node)
				left := stNodes.pop().(*Node)
				op := stOps.pop().(int32)
				stNodes.push(&Node{op, left, right})
			}
			stOps.push(ch)
		}
	}

	for !stOps.isEmpty() {
		right := stNodes.pop().(*Node)
		left := stNodes.pop().(*Node)
		op := stOps.pop().(int32)
		stNodes.push(&Node{op, left, right})
	}

	return stNodes.top().(*Node)
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
