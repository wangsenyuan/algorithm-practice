package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var m int
	fmt.Fscan(reader, &m)
	p := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &p[i])
	}
	return solve(p)
}

const mod = 1e9 + 7

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

type pair struct {
	first  int
	second int
}

func solve(p []int) int {
	sort.Ints(p)
	// 小于200 000的质数，大概有10000个？
	// 如果每个质数都出现，每个质数大概出现20次，这样子 20 * 20 * .... 10000次左右，似乎有问题
	// 因为会很大，所以要用% (mod - 1)来计算

	var arr []pair

	for i := range p {
		if len(arr) == 0 || arr[len(arr)-1].first != p[i] {
			arr = append(arr, pair{p[i], 1})
		} else {
			arr[len(arr)-1].second++
		}
	}

	n := len(arr)
	pref := make([]int, n+1)
	pref[0] = 1

	for i, cur := range arr {
		pref[i+1] = pref[i] * (cur.second + 1) % (mod - 1)
	}

	res := 1
	suf := 1
	for i := n - 1; i >= 0; i-- {
		tmp := suf * pref[i] % (mod - 1)

		sum := (1 + arr[i].second) * arr[i].second / 2

		cur := pow(arr[i].first, sum)

		res = mul(res, pow(cur, tmp))

		suf = suf * (arr[i].second + 1) % (mod - 1)
	}

	return res
}
