package main

import "fmt"

func getNext(t string, next []int) {
	i := 0
	j := 0
	next[0] = -1
	rs := []rune(t)
	for i < len(t) {
		if j == 0 || rs[i] == rs[j] {
			i++
			j++
			next[i] = j
		} else {
			j = next[j]
		}
	}
}

/** 返回子串t在主串s中第pos个字符之后的位置 若不存在，则函数返回值为-1*/
/* T非空,1<=pos<=len(s)*/
func indexKMP(s string, t string, pos int) int {
	i := pos
	j := 0
	next := make([]int, 20)
	getNext(t, next)
	fmt.Println(next)
	for i < len(s) && j < len(t) {
		if j == -1 || s[i] == t[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j >= len(t) {
		return i - j
	}

	return -1
}

func main() {
	fmt.Println(indexKMP("abcdefgab", "bcdef", 0))
	fmt.Println(indexKMP("aaaabcde", "aaaaax", 0))
}
