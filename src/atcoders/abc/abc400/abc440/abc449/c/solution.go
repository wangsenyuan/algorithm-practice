package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n, L, R int
	fmt.Fscan(reader, &n, &L, &R)
	var s string
	fmt.Fscan(reader, &s)
	return solve(L, R, s)
}

func solve(L int, R int, s string) int {
	pos := make([][]int, 26)
	var res int
	for i := range s {
		x := int(s[i] - 'a')
		// L <= i - j <= R
		//  i - R <= j <= i - L
		if i-L >= 0 {
			p1 := sort.SearchInts(pos[x], i-L+1)
			p2 := sort.SearchInts(pos[x], i-R)
			res += p1 - p2
		}
		pos[x] = append(pos[x], i)
	}

	return res
}
