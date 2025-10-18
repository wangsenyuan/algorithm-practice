package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	res := drive(reader)
	for _, v := range res {
		fmt.Fprintln(writer, v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m int
	fmt.Fscan(reader, &n, &m)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	q := make([][]int, m)
	for i := 0; i < m; i++ {
		q[i] = make([]int, 2)
		fmt.Fscan(reader, &q[i][0], &q[i][1])
	}
	return solve(a, q)
}

const H = 7

func solve(a []int, q [][]int) []int {
	mx := slices.Max(a)
	var primes []int
	lpf := make([]int, mx+1)
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			lpf[i] = i
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > mx {
				break
			}
			lpf[p*i] = p
			if i%p == 0 {
				break
			}
		}
	}

	type query struct {
		id int
		l  int
	}

	n := len(a)

	arr := make([][]query, n+1)
	for i, cur := range q {
		l, r := cur[0], cur[1]
		arr[r] = append(arr[r], query{id: i, l: l})
	}

	var right_most [2*H + 1]int
	best := make([][H + 1]int, mx+1)

	ans := make([]int, len(q))

	for i, v := range a {
		var w int

		mul := []int{1}

		for v > 1 {
			var e int
			p := lpf[v]
			for v%p == 0 {
				v /= p
				e ^= 1
			}
			if e == 1 {
				w++
				// 这里貌似不会去迭代新加入的元素
				for _, m := range mul {
					mul = append(mul, m*p)
				}
			}
		}
		for mask, m := range mul {
			common := bits.OnesCount8(uint8(mask))
			for w2 := common; w2 <= H; w2++ {
				op := w + w2 - 2*common
				// best[m][w]表示将某个a[?]变成m的最大的下标
				right_most[op] = max(right_most[op], best[m][w2])
			}
			best[m][w] = i + 1
		}

		for _, qr := range arr[i+1] {
			for j := 0; j <= 2*H; j++ {
				if right_most[j] >= qr.l {
					ans[qr.id] = j
					break
				}
			}
		}
	}
	return ans
}
