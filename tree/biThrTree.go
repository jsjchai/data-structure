package main

import "fmt"

const (
	LinK   int = 0 //表示指向左右孩子指针
	Thread int = 1 //表示指向前驱或后继的线索
)

type BiThrTree struct {
	data   int
	lchild *BiThrTree
	rchild *BiThrTree
	lTag   int //左标志
	rTag   int //右标志
}

//全局变量，始终指向刚刚访问过的结点
var pre *BiThrTree

/* 中序遍历进行中序线索化 */
func inThreading(p *BiThrTree) {
	if p == nil {
		return
	}
	inThreading(p.lchild)
	if p.lchild == nil {
		p.lTag = Thread
		p.lchild = pre
	}
	if pre.rchild == nil {
		pre.rTag = Thread
		pre.rchild = p
	}
	pre = p
	inThreading(p.rchild)
}

/* t指向头结点，头结点左链lchild指向根结点，头结点右链rchild指向中序遍历的最后一个结点，中序遍历二叉线索链表t*/
func inOrderTraverseThr(t *BiThrTree) {
	p := t.lchild
	for p != t {
		for p.lTag == LinK {
			p = p.lchild
		}
		fmt.Print(p.data)
		for p.rTag == Thread && p.rchild != t {
			p = p.rchild
			fmt.Print(p.data)
		}
		p = p.rchild

	}
}

func main() {

}
