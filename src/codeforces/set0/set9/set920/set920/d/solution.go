package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	_, _, _, ok, res := drive(reader)
	if !ok {
		fmt.Fprintln(writer, "NO")
		return
	}
	fmt.Fprintln(writer, "YES")
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1], cur[2])
	}
}

func drive(reader *bufio.Reader) (a []int, K int, V int, ok bool, res [][]int) {
	var n int
	fmt.Fscan(reader, &n, &K, &V)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ok, res = solve(a, K, V)
	return
}

func solve(a []int, K int, V int) (bool, [][]int) {
	if K == 0 {
		if slices.Contains(a, V) {
			return true, nil
		}
		return false, nil
	}
	if V == 0 {
		// 肯定可以
		if a[0] == 0 {
			return true, nil
		}
		c := (a[0] + K - 1) / K
		return true, [][]int{{c, 1, 2}}
	}
	n := len(a)
	// V > 0
	dp := make([]*big.Int, n)
	var sum int
	for i, v := range a {
		sum += v
		dp[i] = big.NewInt(0)
		dp[i].SetBit(dp[i], 0, 1) // empty set has sum 0
		if i > 0 {
			cloneBits(dp[i], dp[i-1])
			for x := range K {
				if dp[i-1].Bit(x) == 1 {
					dp[i].SetBit(dp[i], (x+v)%K, 1)
				}
			}
		}
		dp[i].SetBit(dp[i], v%K, 1)
	}

	if sum < V || dp[n-1].Bit(V%K) == 0 {
		return false, nil
	}
	var pos int
	for pos < n && dp[pos].Bit(V%K) == 0 {
		pos++
	}
	var res [][]int
	var current int
	marked := make([]bool, n)

	add := func(i int) {
		marked[i] = true
		current += a[i]
		if i != pos {
			cnt := (a[i] + K - 1) / K
			res = append(res, []int{cnt, i + 1, pos + 1})
		}
	}

	want := V % K
	for i := pos; i >= 0; i-- {
		if a[i]%K == want {
			add(i)
			break
		}
		if i > 0 {
			v := a[i]
			for x := range K {
				if dp[i-1].Bit(x) == 1 && (x+v)%K == want {
					add(i)
					want = x
					break
				}
			}
		}
	}

	// pick extra: first unmarked non-pos tank
	extra := -1
	for i := 0; i < n; i++ {
		if i != pos && !marked[i] {
			extra = i
			break
		}
	}
	if extra == -1 {
		// all non-pos tanks are in S (marked/emptied into pos), so current >= V
		if current > V {
			dst := 0
			if pos == 0 {
				dst = 1
			}
			cnt := (current - V) / K
			res = append(res, []int{cnt, pos + 1, dst + 1})
		}
		return true, res
	}
	// consolidate all non-S tanks (except extra) into extra
	for i := 0; i < n; i++ {
		if !marked[i] && i != extra && a[i] > 0 {
			cnt := (a[i] + K - 1) / K
			res = append(res, []int{cnt, i + 1, extra + 1})
		}
	}
	// one bulk adjustment between pos and extra
	if current < V {
		cnt := (V - current) / K
		res = append(res, []int{cnt, extra + 1, pos + 1})
	} else if current > V {
		cnt := (current - V) / K
		res = append(res, []int{cnt, pos + 1, extra + 1})
	}

	return true, res
}

func cloneBits(dst *big.Int, src *big.Int) {
	dst.Set(src)
}
