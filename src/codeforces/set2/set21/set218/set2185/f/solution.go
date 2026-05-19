package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		for _, ans := range res {
			fmt.Fprintln(writer, ans)
		}
	}
}

func drive(reader *bufio.Reader) []int {
	var n, q int
	fmt.Fscan(reader, &n, &q)
	a := make([]int, 1<<n)
	for i := range 1 << n {
		fmt.Fscan(reader, &a[i])
	}
	queries := make([][]int, q)
	for i := range q {
		queries[i] = make([]int, 2)
		fmt.Fscan(reader, &queries[i][0], &queries[i][1])
	}
	return solve(a, queries)
}

func solve(a []int, queries [][]int) []int {
	// n = 1 << h
	n := len(a)

	tr := make([]int, 4*n)

	var f func(i int, l int, r int)

	f = func(i int, l int, r int) {
		if l == r {
			tr[i] = a[l]
		} else {
			mid := (l + r) >> 1
			f(i*2+1, l, mid)
			f(i*2+2, mid+1, r)
			tr[i] = tr[i*2+1] ^ tr[i*2+2]
		}
	}

	f(0, 0, n-1)

	var play func(i int, l int, r int, pos int, val int, cnt int) int

	play = func(i int, l int, r int, pos int, val int, cnt int) int {
		if l == r {
			return cnt
		}
		mid := (l + r) >> 1
		if pos <= mid {
			if tr[i*2+1]^val < tr[i*2+2] {
				cnt += r - mid
			}
			return play(i*2+1, l, mid, pos, val, cnt)
		}
		if tr[i*2+1] >= tr[i*2+2]^val {
			cnt += mid - l + 1
		}
		return play(i*2+2, mid+1, r, pos, val, cnt)
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		j, v := cur[0]-1, cur[1]
		v ^= a[j]
		ans[i] = play(0, 0, n-1, j, v, 0)
	}

	return ans
}
