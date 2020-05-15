package main

import (
	"fmt"
	"os"
)

type StaticLinkList struct {
	data   int
	cursor int //游标，存放后续数据下标，为0是无指向
}

//初始化数组状态
func initList(size int) []StaticLinkList {

	list := make([]StaticLinkList, size)
	for i := 0; i < size-1; i++ {
		list[i].cursor = i + 1
	}
	list[size-1].cursor = 0
	return list
}

/*若备用空间链表非空，则返回分配的结点下标,否则返回0*/
func mallocSLL(list []StaticLinkList) int {
	i := list[0].cursor
	if i == 0 {
		os.Exit(0)
	}
	list[0].cursor = list[i].cursor
	return i
}
func insert(list []StaticLinkList, i int, data int) string {
	var j, k, l int
	k = len(list) - 1
	if i < 1 || i > len(list)+1 {
		return "error"
	}
	j = mallocSLL(list)
	if j > 0 {
		list[j].data = data
		for l = 1; l <= i-1; l++ {
			k = list[k].cursor
		}
		list[j].cursor = list[k].cursor
		list[k].cursor = j
		return "ok"
	}
	return "error"
}
func main() {
	list := initList(5)
	insert(list, 1, 11)
	insert(list, 1, 22)
	insert(list, 1, 33)
	insert(list, 1, 44)

	for i, arr := range list {
		fmt.Println("i=", i, "cursor=", arr.cursor, "data=", arr.data)
	}

}
