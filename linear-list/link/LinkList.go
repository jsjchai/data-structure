package main

import (
	"fmt"
	"math/rand"
)

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
	if p == nil || j > i {
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

func (list *LinkList) delete(i int, e *ElemType) string {
	var j int
	var p, q *Node
	p = list.head

	for j = 1; p.next != nil && j < i; j++ {
		p = p.next
	}
	if p.next == nil || j > i {
		return ERROR
	}

	q = p.next
	p.next = q.next
	*e = q.data
	list.length--
	return OK

}

func (list *LinkList) createListHead(n int) {
	var p *Node
	for i := 0; i < n; i++ {
		var num = ElemType(rand.Intn(100) + 1)
		fmt.Println(num)
		p = &Node{data: num, next: list.head.next}
		list.head.next = p
		list.length++
	}
}

func (list *LinkList) createListTail(n int) {
	var p, r *Node
	r = list.head
	for i := 0; i < n; i++ {
		var num = ElemType(rand.Intn(100) + 1)
		fmt.Println(num)
		p = &Node{data: num, next: nil}
		r.next = p
		r = p
		list.length++
	}
	r.next = nil
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

	list := LinkList{head: &Node{data: -1, next: nil}, length: 1}
	//list.createListHead(10)
	list.createListTail(10)
	fmt.Println("------------")
	list.println()

	list.insert(1, 6)
	list.insert(2, 7)
	list.insert(3, 8)
	list.insert(2, 9)
	list.println()
	var e ElemType
	fmt.Println("delete", list.delete(1, &e))
	fmt.Println("e = ", e)
	list.println()
	fmt.Println("get", list.get(2, &e))
	fmt.Println("e = ", e)

}
