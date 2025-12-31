package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readNNums(reader *bufio.Reader, n int) []int {
	arr := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &arr[i])
	}
	return arr
}

func drive(reader *bufio.Reader) string {
	var n, m, l int
	fmt.Fscan(reader, &n, &m, &l)
	A := readNNums(reader, n)
	B := readNNums(reader, m)
	C := readNNums(reader, l)
	return solve(A, B, C)
}

type Card struct {
	val  int
	hand int
}

func solve(A []int, B []int, C []int) string {
	// len(A) + len(B) + len(C) <= 12
	var cards []Card
	for _, v := range A {
		cards = append(cards, Card{val: v, hand: 1})
	}
	for _, v := range B {
		cards = append(cards, Card{val: v, hand: 2})
	}

	for _, v := range C {
		cards = append(cards, Card{val: v, hand: 0})
	}

	slices.SortFunc(cards, func(x, y Card) int {
		return cmp.Or(x.val-y.val, x.hand-y.hand)
	})

	n := len(cards)

	bases := make([]int, n+1)
	bases[0] = 1
	for i := range n {
		bases[i+1] = bases[i] * 3
	}

	// n <= 12
	find := func(mask int, w int) []int {
		var res []int
		for i := 0; mask > 0; i++ {
			v := mask % 3
			if v == w {
				res = append(res, i)
			}
			mask /= 3
		}
		return res
	}

	get := func(mask int, i int) int {
		return mask % bases[i+1] / bases[i]
	}

	play := func(mask int, i int, j int) int {
		v := get(mask, i)
		mask -= v * bases[i]
		if j >= 0 {
			mask += v * bases[j]
		}
		return mask
	}

	dp := make([][]int, bases[n])
	for i := range dp {
		dp[i] = make([]int, 2)
		for j := range 2 {
			dp[i][j] = -1
		}
	}

	var f func(mask int, player int) int
	f = func(mask int, player int) (res int) {
		arr := find(mask, player)
		if len(arr) == 0 {
			// bad luck
			return 0
		}

		if dp[mask][player-1] != -1 {
			return dp[mask][player-1]
		}

		defer func() {
			dp[mask][player-1] = res
		}()

		table := find(mask, 0)
		// len(table) > 0 always true

		for _, v := range arr {
			// 可以直接放进去
			newMask := play(mask, v, -1)
			if f(newMask, 3-player) == 0 {
				res = 1
				return
			}
			for _, w := range table {
				if cards[w].val >= cards[v].val {
					break
				}
				// can switch
				newMask := play(mask, v, w)
				if f(newMask, 3-player) == 0 {
					res = 1
					return
				}
			}
		}

		return
	}

	var flag int
	for i := range n {
		flag += bases[i] * cards[i].hand
	}

	ans := f(flag, 1)

	if ans == 1 {
		return "Takahashi"
	}

	return "Aoki"
}
