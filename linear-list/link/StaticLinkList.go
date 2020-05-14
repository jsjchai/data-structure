package main

const MAX_SIZE = 10

type StaticLinkList struct {
	data   int
	cursor int //游标，存放后续数据下标，为0是无指向
}

//初始化数组状态
func initList(size int) []StaticLinkList {

	list := make([]StaticLinkList, size)
	for i := 0; i < MAX_SIZE-1; i++ {
		list[i].cursor = i + 1
	}
	list[MAX_SIZE-1].cursor = 0
	return list
}

/*若备用空间链表非空，则返回分配的结点下标,否则返回0*/
func mallocSLL(list []StaticLinkList) int {
	i := list[0].cursor
	if list[0].cursor >= 0 {
		list[0].cursor = list[i].cursor
	}
	return i
}
