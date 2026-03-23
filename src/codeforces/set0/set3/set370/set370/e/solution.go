package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	best, res := solve(a)
	if best < 0 {
		fmt.Println(-1)
		return
	}

	fmt.Println(best)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i, x := range res {
		if i > 0 {
			fmt.Fprint(writer, " ")
		}
		fmt.Fprint(writer, x)
	}
	fmt.Fprintln(writer)
}

type pair struct {
	first  int
	second int
}

func solve(a []int) (best int, res []int) {
	n := len(a)
	if a[0] > 1 {
		return -1, nil
	}
	// Any valid diary must start with book 1.
	a[0] = 1
	limit := n/2 + 1
	L := make([]int, limit+1)
	R := make([]int, limit+1)
	for i := range L {
		L[i] = n + 1
		R[i] = -1
	}
	for i, v := range a {
		if v > n/2 {
			return -1, nil
		}
		if v == 0 {
			continue
		}
		L[v] = min(L[v], i)
		R[v] = max(R[v], i)
	}
	var obs []int
	for v := 1; v <= limit; v++ {
		if L[v] > R[v] {
			continue
		}
		if R[v]-L[v]+1 > 5 {
			return -1, nil
		}
		obs = append(obs, v)
	}

	for i := 1; i < len(obs); i++ {
		if R[obs[i-1]] >= L[obs[i]] {
			return -1, nil
		}
	}

	m := len(obs)
	dp := make([][5]bool, m)
	pre := make([][5]pair, m)
	start := make([][5]int, m)
	for i := range m {
		for j := range 5 {
			pre[i][j] = pair{-1, -1}
			start[i][j] = -1
		}
	}

	for i, v := range obs {
		baseLen := R[v] - L[v] + 1
		for right := 0; right < 5 && R[v]+right < n && baseLen+right <= 5; right++ {
			if right > 0 && a[R[v]+right] != 0 {
				break
			}
			end := R[v] + right
			for left := L[v]; left >= max(0, end-4); left-- {
				if left < L[v] && a[left] != 0 {
					break
				}
				if end-left+1 < 2 {
					continue
				}
				if i == 0 {
					if feasible(left, v-1) {
						dp[i][right] = true
						start[i][right] = left
						break
					}
					continue
				}
				gapBooks := v - obs[i-1] - 1
				for prevRight := 0; prevRight < 5; prevRight++ {
					if !dp[i-1][prevRight] {
						continue
					}
					prevEnd := R[obs[i-1]] + prevRight
					if feasible(left-prevEnd-1, gapBooks) {
						dp[i][right] = true
						pre[i][right] = pair{i - 1, prevRight}
						start[i][right] = left
						break
					}
				}
				if dp[i][right] {
					break
				}
			}
		}
	}

	lastRight := -1
	suffixBooks := -1
	lastVal := obs[m-1]
	for right := 0; right < 5; right++ {
		if !dp[m-1][right] {
			continue
		}
		end := R[lastVal] + right
		books := maxBooks(n - end - 1)
		if books < 0 {
			continue
		}
		if lastVal+books > best {
			best = lastVal + books
			lastRight = right
			suffixBooks = books
		}
	}
	if lastRight < 0 {
		return -1, nil
	}

	segL := make([]int, m)
	segR := make([]int, m)
	for i, right := m-1, lastRight; i >= 0; i-- {
		segL[i] = start[i][right]
		segR[i] = R[obs[i]] + right
		right = pre[i][right].second
	}

	res = make([]int, n)
	fillBooks(res, 0, 1, buildLens(segL[0], obs[0]-1))
	for i, v := range obs {
		for j := segL[i]; j <= segR[i]; j++ {
			res[j] = v
		}
		if i+1 < m {
			fillBooks(res, segR[i]+1, v+1, buildLens(segL[i+1]-segR[i]-1, obs[i+1]-v-1))
		}
	}
	fillBooks(res, segR[m-1]+1, obs[m-1]+1, buildLens(n-segR[m-1]-1, suffixBooks))
	return
}

func feasible(days int, books int) bool {
	if books == 0 {
		return days == 0
	}
	return 2*books <= days && days <= 5*books
}

func maxBooks(days int) int {
	switch {
	case days == 0:
		return 0
	case days == 1:
		return -1
	default:
		return days / 2
	}
}

func buildLens(days int, books int) []int {
	if books == 0 {
		return nil
	}
	lens := make([]int, books)
	for i := range lens {
		lens[i] = 2
	}
	extra := days - 2*books
	for i := 0; i < books && extra > 0; i++ {
		add := min(3, extra)
		lens[i] += add
		extra -= add
	}
	return lens
}

func fillBooks(res []int, pos int, book int, lens []int) {
	for _, ln := range lens {
		for range ln {
			res[pos] = book
			pos++
		}
		book++
	}
}
