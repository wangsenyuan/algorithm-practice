package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString(reader)
	s := readString(reader)
	res := solve(s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(s string) string {
	buf := []byte(s)
	n := len(buf)
	h := n / 3

	cnt := make([]int, 3)
	for i := range n {
		cnt[buf[i]-'0']++
	}

	cur := []int{h, h, h}

	check := func() int {
		var res int
		for i := range 3 {
			res += abs(cur[i] - cnt[i])
		}
		return res / 2
	}

	need := check()
	var w int
	for i := range n {
		x := int(buf[i] - '0')
		cnt[x]--

		for j := range 3 {
			if cur[j] == 0 {
				continue
			}
			var rep int
			if j != x {
				rep++
			}
			cur[j]--
			if w+rep+check() == need {
				buf[i] = byte(j + '0')
				w += rep
				break
			}
			cur[j]++
		}

	}

	return string(buf)
}

func solve1(s string) string {
	buf := []byte(s)
	n := len(buf)
	h := n / 3

	cnt := make([]int, 3)
	for i := range n {
		cnt[buf[i]-'0']++
	}

	if cnt[0] < h {
		for i := range n {
			x := int(buf[i] - '0')
			if cnt[x] > h && cnt[0] < h {
				buf[i] = '0'
				cnt[0]++
				cnt[x]--
			}
		}
	} else if cnt[0] > h {
		var pos int
		for i := 0; i < n && cnt[1] < h && cnt[0] > h; i++ {
			if buf[i] == '0' {
				pos++
				if pos > h {
					buf[i] = '1'
					cnt[1]++
					cnt[0]--
				}
			}
		}

		pos = 0
		for i := 0; i < n && cnt[2] < h && cnt[0] > h; i++ {
			if buf[i] == '0' {
				pos++
				if pos > h {
					buf[i] = '2'
					cnt[2]++
					cnt[0]--
				}
			}
		}
	}

	for i := 0; i < n && cnt[1] < h; i++ {
		if buf[i] == '2' {
			buf[i] = '1'
			cnt[1]++
			cnt[2]--
		}
	}

	for i := n - 1; i >= 0 && cnt[1] > h; i-- {
		if buf[i] == '1' {
			buf[i] = '2'
			cnt[1]--
			cnt[2]++
		}
	}

	return string(buf)
}

func abs(a int) int {
	return max(a, -a)
}
