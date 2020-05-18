package main

import "fmt"

const MAX_SIZE = 10

/**两栈共享空间*/
type SqDoubleStack struct {
	data []int
	top1 int //栈1栈顶指针
	top2 int //栈2栈顶指针
}

/**
  e 栈顶元素
  stackNumber 1号栈还是2号栈，1,2
*/
func (stack *SqDoubleStack) push(e int, stackNumber int) string {

	//栈已满,不能push新元素
	if stack.top1+1 == stack.top2 {
		return "ERROR"
	}
	if stackNumber == 1 {
		stack.top1++
		stack.data[stack.top1] = e
	} else if stackNumber == 2 {
		stack.top2--
		stack.data[stack.top2] = e
	} else {
		return "ERROR"
	}

	return "OK"
}

func (stack *SqDoubleStack) pop(e *int, stackNumber int) string {
	if stackNumber == 1 {
		if stack.top1 == -1 {
			return "ERROR"
		}
		*e = stack.data[stack.top1]
		stack.top1--
	} else if stackNumber == 2 {
		if stack.top2 == MAX_SIZE {
			return "ERROR"
		}
		*e = stack.data[stack.top2]
		stack.top2++
	}
	return "OK"
}

func main() {
	stack := SqDoubleStack{data: make([]int, MAX_SIZE), top1: -1, top2: MAX_SIZE}
	stack.push(10, 1)
	stack.push(20, 2)

	var e int
	stack.pop(&e, 1)
	fmt.Println(e)
	stack.pop(&e, 2)
	fmt.Println(e)

}
