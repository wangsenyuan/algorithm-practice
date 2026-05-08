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
		fmt.Fprintln(writer, res)
	}
}

func drive(reader *bufio.Reader) int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	a := make([][]int, k)
	for i := range k {
		a[i] = make([]int, 3)
		fmt.Fscan(reader, &a[i][0], &a[i][1], &a[i][2])
	}
	return solve(n, m, a)
}

const mod = 1_000_000_007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func mul(a, b int) int {
	return a * b % mod
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res = mul(res, a)
		}
		a = mul(a, a)
		b >>= 1
	}
	return res
}

func solve(n int, m int, a [][]int) int {
	if n*m == 1 {
		if len(a) == 0 {
			return 2
		}
		return 1
	}
	var inner int
	var corner int
	var other int

	if n >= 2 && m >= 2 {
		inner = (n - 2) * (m - 2)
		corner = 4
		other = 2*(n-2) + 2*(m-2)
	} else if n == 1 {
		inner = 0
		corner = 2
		other = m - 2
	} else {
		// m == 1
		inner = 0
		corner = 2
		other = n - 2
	}

	for _, cur := range a {
		r, c := cur[0]-1, cur[1]-1
		if r > 0 && r < n-1 && c > 0 && c < m-1 {
			inner--
		} else if r == 0 && c == 0 || r == 0 && c == m-1 || r == n-1 && c == 0 || r == n-1 && c == m-1 {
			corner--
		} else {
			other--
		}
	}

	res := pow(2, inner+corner)
	if other > 0 {
		return mul(res, pow(2, other-1))
	}
	type pair struct {
		first  int
		second int
	}
	marked := make(map[pair]int)

	for _, cur := range a {
		marked[pair{cur[0], cur[1]}] = cur[2]
	}

	// 绿色的都当作黑色处理
	checkEmpty := func(r int, c int, color int) bool {
		if color == 0 {
			return false
		}
		// 这个格子是空的
		if _, ok := marked[pair{r, c}]; !ok {
			return true
		}
		return false
	}

	checkDifferent := func(r int, c int, color int) bool {
		if c2, ok := marked[pair{r, c}]; ok && c2 != color {
			return true
		}
		return false
	}

	var cnt int
	var cnt2 int
	for _, cur := range a {
		r, c, color := cur[0], cur[1], cur[2]
		if r > 1 && checkEmpty(r-1, c, color) {
			cnt++
		}
		if r < n && checkEmpty(r+1, c, color) {
			cnt++
		}

		if c > 1 && checkEmpty(r, c-1, color) {
			cnt++
		}
		if c < m && checkEmpty(r, c+1, color) {
			cnt++
		}

		if r > 1 && checkDifferent(r-1, c, color) {
			cnt2++
		}

		if r < n && checkDifferent(r+1, c, color) {
			cnt2++
		}
		if c > 1 && checkDifferent(r, c-1, color) {
			cnt2++
		}
		if c < m && checkDifferent(r, c+1, color) {
			cnt2++
		}
	}
	cnt2 /= 2

	cnt += cnt2

	if cnt%2 == 1 {
		return 0
	}

	return res
}
