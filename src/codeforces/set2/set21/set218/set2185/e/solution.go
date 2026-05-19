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
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Fprintln(writer, s[1:len(s)-1])
	}
}

func drive(reader *bufio.Reader) []int {
	var n, m, k int
	fmt.Fscan(reader, &n, &m, &k)
	robots := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &robots[i])
	}
	spikes := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &spikes[i])
	}
	var commands string
	fmt.Fscan(reader, &commands)
	return solve(robots, spikes, commands)
}

const inf = 1 << 60

type data struct {
	min_l int
	max_r int
}

func solve(robots []int, spikes []int, commands string) []int {
	slices.Sort(spikes)

	n := len(commands)
	dp := make([]data, n+1)

	var move int
	for i := range n {
		if commands[i] == 'L' {
			move--
		} else {
			move++
		}
		dp[i+1].min_l = min(dp[i].min_l, move)
		dp[i+1].max_r = max(dp[i].max_r, move)
	}

	check := func(pos int, k int) bool {
		d := sort.SearchInts(spikes, pos)
		// spikes[r] >= pos
		l, r := -inf, inf
		if d < len(spikes) {
			r = spikes[d]
		}
		if d > 0 {
			l = spikes[d-1]
		}
		dl := dp[k].min_l
		dr := dp[k].max_r
		return pos+dl <= l || pos+dr >= r
	}

	diff := make([]int, len(commands)+1)
	diff[0] = len(robots)

	for _, v := range robots {
		if check(v, len(commands)) {
			j := sort.Search(len(commands), func(j int) bool {
				return check(v, j)
			})
			diff[j-1]--
		}

	}
	for i := 1; i < len(commands); i++ {
		diff[i] += diff[i-1]
	}
	return diff[:len(commands)]
}
