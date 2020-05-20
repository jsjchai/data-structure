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

func main() {
	fmt.Println(index("abcadacd", "cd", 1))
}
