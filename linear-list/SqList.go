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

func newInstance() *SqList {
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

/**
是否已满
*/
func (list *SqList) isFull() bool {
	return list.length == MAXSIZE
}

/**
添加数据
*/
func (list *SqList) append(e ElemType) bool {
	if list.isFull() {
		fmt.Println("list is full")
		return false
	}
	list.data[list.length] = e
	list.length++
	return true
}
func (list *SqList) delete(i int, e *ElemType) string {
	var k int
	if list.length == 0 {
		return ERROR
	}

	if i < 1 || i > list.length {
		return ERROR
	}
	*e = list.data[i-1]
	if i < list.length {
		for k = i; k < list.length; k++ {
			list.data[k-1] = list.data[k]
		}
	}
	list.length--
	return OK

}

/**打印*/
func (list *SqList) printf() {
	for i := 0; i < len(list.data); i++ {
		fmt.Println(i, list.data[i])
	}
}

func main() {
	list := newInstance()
	list.printf()
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	list.append(5)
	fmt.Println("insert:", list.insert(1, 7))
	list.printf()
	var del ElemType
	fmt.Println("delete:", list.delete(2, &del))
	fmt.Println("delete e =", del)
}
