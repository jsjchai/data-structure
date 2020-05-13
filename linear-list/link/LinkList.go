package link

import "fmt"

const (
	MAXSIZE int    = 10 //最大存储容量
	OK      string = "ok"
	ERROR   string = "error"
)

//元素类型
type ElemType int

type Node struct {
	data ElemType
	next *Node
}

type LinkList struct {
	node *Node
}

func (list *LinkList) get(i int, e *ElemType) string {
	var j int
	p := list.node.next
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

	p := list.node
	var j int
	for j = 1; j < i; j++ {
		p = p.next
	}
	if j > i {
		return ERROR
	}
	var s = &Node{data: e, next: p.next}
	p.next = s
	return OK
}

func (list *LinkList) println() {
	fmt.Println(list.node.data)
}

func main() {
	var list LinkList
	list.insert(1, 6)
	list.println()
}
