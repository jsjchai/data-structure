package main

import "fmt"

type StackNode struct {
	data int
	next *StackNode
}

/**链栈*/
type LinkStack struct {
	top   *StackNode
	count int
}

func (stack *LinkStack) push(e int) {
	s := StackNode{data: e, next: stack.top}
	stack.top = &s
	stack.count++
}

func (stack *LinkStack) isEmpty() bool {
	return stack.count == 0
}

func (stack *LinkStack) pop() *int {
	if stack.isEmpty() {
		return nil
	}
	e := stack.top.data
	stack.top = stack.top.next
	stack.count--
	return &e
}

func main() {
	stack := LinkStack{top: nil, count: 0}
	stack.push(11)
	stack.push(21)
	e := *stack.pop()

	fmt.Println(e)

}
