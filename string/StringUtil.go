package main

import "fmt"

/*T为非空串 若主串s中第pos个字符之后存在与T相等的子串，则返回第一个这样的子串在s中的位置，否则返回0*/
func index(s string, t string, pos int) int {
	if pos <= 0 {
		return 0
	}

	n := len(s)
	m := len(t)
	i := pos
	rs := []rune(s)
	for i <= n-m+1 {
		sub := string(rs[i-1 : m+i-1])
		fmt.Println(sub)
		if sub != t {
			i++
		} else {
			return i
		}
	}
	return 0
}

/*返回子串T在主串s中第pos个字符之后的位置 若不存在，则函数返回值为0*/
/*T非空，1<=pos<=len(s)*/
func index1(s string, t string, pos int) int {
	i := pos
	j := 1
	rs := []rune(s)
	rt := []rune(t)
	for i <= len(s) && j <= len(t) {
		fmt.Println(i, " ", rs[i-1], " ", j, " ", rt[j-1])
		if rs[i-1] == rt[j-1] {
			i++
			j++
		} else {
			i = i - j + 2 /*i退回到上次匹配首位的下一位*/
			j = 1
		}
	}
	if j > len(t) {
		return i - len(t)
	}

	return 0
}

func main() {
	//fmt.Println(index("abcadacd", "cd", 1))
	//fmt.Println(index1("goodgoogle", "google", 1))
	fmt.Println(index1("00000000000000000000000000000000000001", "00000001", 1))
}
