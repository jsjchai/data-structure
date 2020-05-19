package main

import "fmt"

const MAXSIZE = 10

type SqQueue struct {
	data  []int
	front int //头指针
	rear  int //尾指针，若队列不空，指向队列尾元素的下一个位置
}

func newInstance() *SqQueue {
	return &SqQueue{data: make([]int, MAXSIZE), front: 0, rear: 0}
}

func (q *SqQueue) queueLength() int {
	return (q.rear - q.front + MAXSIZE) % MAXSIZE
}

/**入队列**/
func (q *SqQueue) enQueue(e int) string {

	//队列已满
	if ((q.rear + 1) % MAXSIZE) == q.front {
		return "ERROR"
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % MAXSIZE
	return "OK"
}

/**出队列**/
func (q *SqQueue) deQueue() *int {

	//空队列
	if q.front == q.rear {
		return nil
	}
	e := q.data[q.front]
	q.front = (q.front + 1) % MAXSIZE
	return &e
}

func main() {
	q := newInstance()
	q.enQueue(1)
	q.enQueue(2)
	q.enQueue(3)
	q.enQueue(4)
	q.enQueue(5)

	for q.front != q.rear {
		e := *q.deQueue()
		fmt.Println(e)
	}

}
