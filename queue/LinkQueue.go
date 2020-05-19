package main

import "fmt"

type QNode struct {
	data int
	next *QNode
}

type LinkQueue struct {
	front *QNode
	rear  *QNode
}

func (q *LinkQueue) enQueue(e int) {
	s := &QNode{data: e, next: nil}

	if q.rear == nil {
		q.rear = s
		q.front.next = s
	} else {
		//把拥有元素e新结点s赋值给原队尾结点的后继
		q.rear.next = s
		// 把当前的s设置为队尾结点
		q.rear = s
	}

}

func (q *LinkQueue) deQueue() *int {
	if q.front == nil {
		return nil
	}
	p := q.front.next
	e := p.data
	q.front.next = p.next
	if q.rear == p {
		q.rear = q.front
	}
	return &e
}

func main() {
	q := new(LinkQueue)
	q.front = new(QNode)
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	for q.front != q.rear {
		fmt.Println(*q.deQueue())
	}

}
