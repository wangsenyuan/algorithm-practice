package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)

	res := solve(n, k, []int{a, b}, []int{c, d})

	if res == nil {
		fmt.Println(-1)
		return
	}

	for _, row := range res {
		s := fmt.Sprintf("%v", row)
		fmt.Println(s[1 : len(s)-1])
	}

}

func solve(n int, k int, first []int, second []int) [][]int {
	if n == 4 || k <= n {
		return nil
	}
	a, b := first[0], first[1]
	c, d := second[0], second[1]

	special := make([]bool, n+1)
	special[a] = true
	special[b] = true
	special[c] = true
	special[d] = true

	id := make([]int, n)

	for i, x := 2, 1; i < n-2; i++ {
		for special[x] {
			x++
		}
		id[i] = x
		x++
	}

	id[0] = a
	id[1] = c
	id[n-2] = b
	id[n-1] = d

	res := make([][]int, 2)
	res[0] = append(res[0], id[0], id[1])
	for i := 2; i < n-2; i++ {
		res[0] = append(res[0], id[i])
	}
	res[0] = append(res[0], id[n-1], id[n-2])

	res[1] = append(res[1], id[1], id[0])

	for i := 2; i < n; i++ {
		res[1] = append(res[1], id[i])
	}
	return res
}
