package main

/**
孩子表示法
*/
type CTNode struct {
	child int
	next  *CTNode
}

/*
表头结构
*/
type CTBox struct {
	data       int
	firstchild *CTNode
}

type CTree struct {
	nodes []CTBox
	r     int
	n     int
}

func main() {

}
