package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) string {
	s := readString(reader)
	t := readString(reader)
	return solve(s, t)
}

func solve(s string, t string) string {
	cnt := make([]int, 26)
	for _, v := range []byte(s) {
		cnt[v-'a']++
	}

	n := len(t)
	sum := make([]int, 26)
	for i := 0; i < n; i++ {
		sum[t[i]-'a']++
	}

	check := func(i int, x int) bool {
		if cnt[x] == 0 {
			return false
		}
		cnt[x]--
		for j := 0; j < 26; j++ {
			if cnt[j] < sum[j] {
				cnt[x]++
				return false
			}
		}
		cnt[x]++
		return true
	}
	m := len(s)
	construct := func(i int, x int) string {
		buf := make([]byte, m)
		for j := 0; j < i; j++ {
			x := int(t[j] - 'a')
			cnt[x]--
			buf[j] = t[j]
		}
		buf[i] = byte(x + 'a')
		cnt[x]--
		for j, y := i+1, 0; j < m; j++ {
			for y < 26 && cnt[y] == 0 {
				y++
			}
			cnt[y]--
			buf[j] = byte(y + 'a')
		}
		return string(buf)
	}

	// 如果t是x的一个前缀，这个要特殊处理一下
	if n < m {
		ok := true
		for i := range 26 {
			if sum[i] > cnt[i] {
				ok = false
				break
			}
		}
		if ok {
			buf := make([]byte, m)
			for j := 0; j < n; j++ {
				x := int(t[j] - 'a')
				cnt[x]--
				buf[j] = t[j]
			}
			for j, y := n, 0; j < m; j++ {
				for y < 26 && cnt[y] == 0 {
					y++
				}
				cnt[y]--
				buf[j] = byte(y + 'a')
			}
			return string(buf)
		}
	}

	for i := n - 1; i >= 0; i-- {
		// 如果能够由s的部分，组成t的前缀
		x := int(t[i] - 'a')
		sum[x]--
		for y := x + 1; y < 26; y++ {
			if check(i, y) {
				return construct(i, y)
			}
		}
	}

	return "-1"
}
