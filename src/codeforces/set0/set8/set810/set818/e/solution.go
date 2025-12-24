package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a, k)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, k int) int {
	var fs []pair

	for i := 2; i <= k; i++ {
		if k%i == 0 {
			var cnt int
			for k%i == 0 {
				cnt++
				k /= i
			}
			fs = append(fs, pair{i, cnt})
		}
	}

	if k > 1 {
		fs = append(fs, pair{k, 1})
	}

	m := len(fs)

	freq := make([]int, m)

	update := func(num int, d int) {
		for i := range m {
			v := fs[i].first
			var c int
			for num%v == 0 {
				c++
				num /= v
			}
			freq[i] += c * d
		}
	}

	check := func() bool {
		for i := range m {
			if freq[i] < fs[i].second {
				return false
			}
		}
		return true
	}

	n := len(a)
	var l int
	var res int
	for r := range n {
		update(a[r], 1)
		for l < r && check() {
			update(a[l], -1)
			if !check() {
				update(a[l], 1)
				break
			}
			l++
		}
		if check() {
			res += l + 1
		}
	}

	return res
}
