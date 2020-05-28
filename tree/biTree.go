package main

import (
	"fmt"
)

type BiTree struct {
	data   int
	lchild *BiTree
	rchild *BiTree
}

/*前序遍历*/
func preOrderTraverse(t *BiTree) {
	if t == nil {
		return
	}
	fmt.Print(t.data)
	preOrderTraverse(t.lchild)
	preOrderTraverse(t.rchild)
}

/*中序遍历*/
func inOrderTraverse(t *BiTree) {
	if t == nil {
		return
	}
	inOrderTraverse(t.lchild)
	fmt.Print(t.data)
	inOrderTraverse(t.rchild)
}

/*后序遍历*/
func postOrderTraverse(t *BiTree) {
	if t == nil {
		return
	}
	preOrderTraverse(t.lchild)
	preOrderTraverse(t.rchild)
	fmt.Print(t.data)
}

func main() {

	tree := new(BiTree)
	tree.data = 1
	tree.rchild = &BiTree{data: 3, lchild: nil, rchild: nil}
	tree.lchild = &BiTree{data: 2, lchild: nil, rchild: nil}
	fmt.Println("前序遍历：")
	preOrderTraverse(tree)
	fmt.Println()
	fmt.Println("中序遍历：")
	inOrderTraverse(tree)
	fmt.Println()
	fmt.Println("后序遍历：")
	postOrderTraverse(tree)

}
