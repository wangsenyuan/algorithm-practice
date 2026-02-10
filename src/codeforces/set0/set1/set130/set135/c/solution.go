package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s := readString(reader)
	res := solve(s)
	for _, s := range res {
		fmt.Println(s)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

type data struct {
	val int
	pos int
}

func solve(s string) []string {
	n := len(s)

	cnt := make([]int, 3)
	for i := range n {
		if s[i] == '0' {
			cnt[0]++
		} else if s[i] == '1' {
			cnt[1]++
		} else {
			cnt[2]++
		}
	}

	var res []string

	if cnt[0]+cnt[2] > cnt[1] {
		res = append(res, "00")
	}
	if cnt[0]+1 < cnt[1]+cnt[2] {
		res = append(res, "11")
	}
	if s[n-1] != '0' {
		if s[n-1] == '?' {
			cnt[2]--
			cnt[1]++
		}
		x := (cnt[0] + cnt[2] - cnt[1] + n&1) / 2
		if x >= 0 && x <= cnt[2] {
			res = append(res, "01")
		}
		if s[n-1] == '?' {
			cnt[2]++
			cnt[1]--
		}
	}
	if s[n-1] != '1' {
		if s[n-1] == '?' {
			cnt[2]--
			cnt[0]++
		}
		// 如果将x变成1, cnt[2] - x 变成0
		// cnt[1] + x = cnt[0] + cnt[2] - x
		x := (cnt[1] + cnt[2] - cnt[0] - n&1) / 2
		if x >= 0 && x <= cnt[2] {
			res = append(res, "10")
		}
		if s[n-1] == '?' {
			cnt[2]++
			cnt[0]--
		}
	}

	slices.Sort(res)

	return res
}
