package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func drive(reader *bufio.Reader) []int {
	var x, y string
	fmt.Fscan(reader, &x, &y)
	var n int
	fmt.Fscan(reader, &n)
	L := make([]int, n)
	R := make([]int, n)
	C := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &L[i], &R[i], &C[i])
	}

	return solve(x, y, L, R, C)
}

func solve(x string, y string, L []int, R []int, C []string) []int {
	var freq [][26]int

	var f0 [26]int

	var w []int

	pref0 := make([][]int, len(x)+1)
	pref0[0] = make([]int, 26)
	for i := range len(x) {
		pref0[i+1] = make([]int, 26)
		f0[int(x[i]-'a')]++
		copy(pref0[i+1], pref0[i])
		pref0[i+1][int(x[i]-'a')]++
	}

	freq = append(freq, f0)
	w = append(w, len(x))

	pref1 := make([][]int, len(y)+1)
	pref1[0] = make([]int, 26)

	var f1 [26]int
	for i := range len(y) {
		f1[int(y[i]-'a')]++
		pref1[i+1] = make([]int, 26)
		copy(pref1[i+1], pref1[i])
		pref1[i+1][int(y[i]-'a')]++
	}

	freq = append(freq, f1)
	w = append(w, len(y))

	inf := slices.Max(R) + 2

	for {
		var f [26]int
		f0 = freq[len(freq)-2]
		f1 = freq[len(freq)-1]
		ok := false
		for i := range 26 {
			f[i] = min(inf, f0[i]+f1[i])
			if f[i] < inf && f[i] > 0 {
				// f[i] = 0, not valid
				ok = true
			}
		}
		freq = append(freq, f)
		w = append(w, min(inf, w[len(w)-2])+w[len(w)-1])

		if !ok {
			break
		}
	}

	// 在s[i][:r]中，找到c的数量
	var play func(i int, r int, c int) int
	play = func(i int, r int, c int) int {
		if freq[i][c] == 0 {
			// c 不存在
			return 0
		}
		if w[i] == r {
			// r就是s[i]的长度
			return freq[i][c]
		}
		// r < w[i] holds
		if i == 0 {
			// s[i] = x, 那么
			return pref0[r][c]
		}
		if i == 1 {
			return pref1[r][c]
		}
		// i >= 2
		if w[i-1] >= r {
			return play(i-1, r, c)
		}
		// w[i-2] < r
		return freq[i-1][c] + play(i-2, r-w[i-1], c)
	}

	ans := make([]int, len(L))

	for i := range len(L) {
		c := int(C[i][0] - 'a')

		ans[i] = play(len(freq)-1, R[i], c)
		ans[i] -= play(len(freq)-1, L[i]-1, c)
	}

	return ans
}
