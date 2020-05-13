package main

import (
	"fmt"
)

const (
	MAXSIZE int    = 10 //最大存储容量
	OK      string = "ok"
	ERROR   string = "error"
)

//元素类型
type ElemType int

type SqList struct {
	data   []ElemType //数组存储数据元素，最大值MAXSIZE
	length int        //当前长度
}

func new() *SqList {
	return &SqList{data: make([]ElemType, MAXSIZE)}
}

/**
获取元素
*/
func (list *SqList) getElm(i int, e ElemType) string {
	if i < 0 || i > list.length {
		return ERROR
	}
	e = list.data[i]
	return OK
}

/**
在某个位置插入元素
*/
func (list *SqList) insert(i int, e ElemType) string {
	var k int
	if MAXSIZE == list.length {
		return ERROR
	}

	if i < 1 || i > list.length+1 {
		return ERROR
	}

	if i <= list.length {

		//将要插入位置后数据元素向后移动一位
		for k = list.length - 1; k >= i-1; k-- {
			list.data[k+1] = list.data[k]
		}
	}
	list.data[i-1] = e
	list.length++
	return OK
}

func (list *SqList) isFull() bool {
	return list.length == MAXSIZE
}

func (list *SqList) append(e ElemType) bool {
	if list.isFull() {
		fmt.Println("list is full")
		return false
	}
	list.data[list.length] = e
	list.length++
	return true
}

/**打印*/
func (list *SqList) prinf() {
	for i := 0; i < len(list.data); i++ {
		fmt.Println(i, list.data[i])
	}
}

func main() {
	list := new()
	list.prinf()
	list.append(2)
	fmt.Println("insert:", list.insert(1, 7))
	list.prinf()
}
