package main

import "fmt"

const (
	OK    string = "ok"
	ERROR string = "error"
)

//元素类型
type ElemType int

type Node struct {
	data ElemType
	next *Node
}

type LinkList struct {
	head   *Node
	length int
}

func (list *LinkList) get(i int, e *ElemType) string {
	var j int
	p := list.head.next
	for j = 1; j < i; j++ {
		p = p.next
	}
	if p == nil || j > 1 {
		return ERROR
	}
	*e = p.data
	return OK
}

func (list *LinkList) insert(i int, e ElemType) string {

	p := list.head
	var j int
	for j = 1; j < i; j++ {
		p = p.next
	}
	if j > i {
		return ERROR
	}
	s := Node{data: e, next: p.next}
	p.next = &s
	list.length++
	return OK
}

func (list *LinkList) println() {
	if list.head == nil || list.length == 0 {
		return
	}

	var temp = list.head
	for i := 0; i < list.length; i++ {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func main() {

	list := LinkList{head: &Node{data: 4, next: nil}, length: 1}
	list.insert(1, 6)
	list.println()
}
