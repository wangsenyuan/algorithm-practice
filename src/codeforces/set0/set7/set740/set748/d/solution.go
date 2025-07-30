package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ans := process(reader)
	fmt.Println(ans)
}

func process(reader *bufio.Reader) int {
	var m, n int
	fmt.Fscan(reader, &m, &n)
	words := make([]string, m)
	a := make([]int, m)
	for i := range m {
		fmt.Fscan(reader, &words[i], &a[i])
	}
	return solve(words, a)
}

func solve(words []string, a []int) int {
	id := make(map[string]int)

	var pair []int

	var scores [][]int
	var flag []int

	add := func(s string, score int) {
		if v, ok := id[s]; !ok {
			id[s] = len(scores)
			scores = append(scores, []int{score})
			flag = append(flag, 0)
			pair = append(pair, -1)
		} else {
			scores[v] = append(scores[v], score)
		}
	}

	for i := range words {
		s := words[i]
		add(s, a[i])

		r := reverse(s)

		if s != r {
			if v, ok := id[r]; ok && pair[id[s]] == -1 {
				pair[id[s]] = v
				pair[v] = id[s]
			}
		} else {
			flag[id[s]] = 1
		}
	}

	for i := range scores {
		slices.Sort(scores[i])
		slices.Reverse(scores[i])
	}

	var ans int
	marked := make([]bool, len(scores))

	var arr []int

	for i := range len(scores) {
		if marked[i] {
			continue
		}
		if flag[i] == 1 {
			// 自己和自己可以组成回文的
			for len(scores[i]) >= 2 && scores[i][0] >= 0 && scores[i][1] >= 0 {
				ans += scores[i][0] + scores[i][1]
				scores[i] = scores[i][2:]
			}
			// 剩下那个要么放在最中间，要么放在两边，但是有可能使用一个负数
			if len(scores[i]) > 0 && scores[i][0] > 0 {
				arr = append(arr, i)
			}
		} else if pair[i] != -1 {
			j := pair[i]
			marked[j] = true
			for u := 0; u < len(scores[i]) && u < len(scores[j]) && scores[i][u]+scores[j][u] > 0; u++ {
				ans += scores[i][u] + scores[j][u]
			}
		}
	}

	if len(arr) == 1 {
		ans += scores[arr[0]][0]
	} else if len(arr) > 1 {
		m := len(arr)
		pref := make([]int, m+1)
		for i := range m {
			pref[i+1] = pref[i]
			id := arr[i]
			if len(scores[id]) > 1 {
				pref[i+1] += max(0, scores[id][0]+scores[id][1])
			}
			// else 它只能放一个，所以没法放在两边
		}
		var best int
		var suf int
		for i := m - 1; i >= 0; i-- {
			id := arr[i]
			best = max(best, pref[i]+suf+scores[id][0])
			// 把id放在中间
			if len(scores[id]) > 1 {
				suf += max(0, scores[id][0]+scores[id][1])
			}
		}
		ans += best
	}

	return ans
}

func reverse(s string) string {
	buf := []byte(s)
	slices.Reverse(buf)
	return string(buf)
}
