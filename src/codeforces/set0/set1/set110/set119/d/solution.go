package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res[0], res[1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	n := len(s)
	for n > 0 && (s[n-1] == '\n' || s[n-1] == '\r') {
		n--
	}
	return s[:n]
}

func drive(reader *bufio.Reader) []int {
	a := readString(reader)
	b := readString(reader)
	return solve(a, b)
}

const inf = 1 << 60

func solve(a string, b string) (res []int) {
	if len(a) != len(b) {
		return []int{-1, -1}
	}
	n := len(a)
	dp1 := zFunction(b + "#" + a)[n+1:]
	// dp1[i] = a中是b的前缀的长度
	br := reverse(b)
	dp2 := make([]int, n+1)
	for i := range n + 1 {
		dp2[i] = -inf
	}
	for i := 0; i+1 < n && a[i] == br[i]; i++ {
		dp2[i+1] = dp1[i+1]
	}

	res = []int{-1, -1}
	ar := reverse(a)
	p := kmp(ar)

	update := func(i int, j int) {
		if res[0] < i || res[0] == i && res[1] > j {
			res[0] = i
			res[1] = j
		}
	}

	for i, j := 0, 0; i < n; i++ {
		for j > 0 && ar[j] != b[i] {
			j = p[j-1]
		}
		if ar[j] == b[i] {
			j++
		}
		// j == n这个情况算什么呢？
		if j == n {
			// update(0, n-1)
			j = p[j-1]
			continue
		}
		// j是当前位置，匹配的ar的，那么 i = n - i
		r := n - j
		l := n - i - 1
		if l+dp2[l] >= r {
			update(l-1, r)
		}
	}

	return
}

func reverse(s string) string {
	buf := []byte(s)
	slices.Reverse(buf)
	return string(buf)
}

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	for i, l, r := 1, 0, 0; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}
	return z
}

func kmp(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := p[i-1]
		for j > 0 && s[i] != s[j] {
			j = p[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		p[i] = j
	}
	return p
}
