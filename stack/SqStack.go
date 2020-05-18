package main

import "fmt"

const MAXSIZE int = 5

type SqStack struct {
	data []int
	top  int //用于栈顶指针
}

func (stack *SqStack) push(e int) string {
	if stack.top == MAXSIZE-1 {
		return "ERROR"
	}
	stack.top++
	stack.data[stack.top] = e
	return "OK"
}

func (stack *SqStack) pop(e *int) string {
	if stack.top == -1 {
		return "ERROR"
	}
	*e = stack.data[stack.top]
	stack.top--
	return "OK"
}

func main() {

	stack := SqStack{data: make([]int, MAXSIZE), top: -1}
	stack.push(5)
	stack.push(4)
	stack.push(3)
	stack.push(2)
	stack.push(1)

	fmt.Println("出栈：")

	var e int
	for stack.top >= 0 {
		stack.pop(&e)
		fmt.Print(e, " ")
	}

}
