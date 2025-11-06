package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, _, _, d, res := drive(reader)
	fmt.Fprintln(writer, d)
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1])
	}
}

func drive(reader *bufio.Reader) (s int, a []int, b []int, gadgets [][]int, d int, res [][]int) {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k, &s)

	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	b = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &b[i])
	}
	gadgets = make([][]int, m)
	for i := range m {
		gadgets[i] = make([]int, 2)
		fmt.Fscan(reader, &gadgets[i][0], &gadgets[i][1])
	}
	d, res = solve(k, s, a, b, gadgets)
	return
}

type Gadget struct {
	id    int
	price int
}

func solve(k int, s int, a []int, b []int, gadgets [][]int) (int, [][]int) {
	n := len(a)
	// m := len(gadgets)
	var dollars []Gadget
	var pounds []Gadget

	for i, cur := range gadgets {
		if cur[0] == 1 {
			dollars = append(dollars, Gadget{id: i, price: cur[1]})
		} else {
			pounds = append(pounds, Gadget{id: i, price: cur[1]})
		}
	}

	slices.SortFunc(dollars, func(a, b Gadget) int {
		return a.price - b.price
	})
	slices.SortFunc(pounds, func(a, b Gadget) int {
		return a.price - b.price
	})

	dp := make([][2]int, n)

	for i := 0; i < n; i++ {
		dp[i][0] = i
		dp[i][1] = i
		if i > 0 && a[i] >= a[dp[i-1][0]] {
			dp[i][0] = dp[i-1][0]
		}
		if i > 0 && b[i] >= b[dp[i-1][1]] {
			dp[i][1] = dp[i-1][1]
		}
	}

	var res [][]int

	check := func(d int) bool {
		if d == 0 {
			return false
		}
		d1 := dp[d-1][0]
		d2 := dp[d-1][1]
		p1 := a[d1]
		p2 := b[d2]

		res = res[:0]

		for x, i, j := s, 0, 0; len(res) < k && x > 0 && (i < len(dollars) || j < len(pounds)); {
			if j == len(pounds) || i < len(dollars) && dollars[i].price*p1 <= pounds[j].price*p2 {
				x -= dollars[i].price * p1
				if x < 0 {
					break
				}
				res = append(res, []int{dollars[i].id + 1, d1 + 1})
				i++
			} else {
				x -= pounds[j].price * p2
				if x < 0 {
					break
				}
				res = append(res, []int{pounds[j].id + 1, d2 + 1})
				j++
			}
			if x < 0 {
				break
			}
		}

		return len(res) == k
	}

	if !check(n) {
		return -1, nil
	}
	d := sort.Search(n, check)
	check(d)
	return d, res
}
