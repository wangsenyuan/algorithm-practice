package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (t string, res string) {
	s := readString(reader)
	t = readString(reader)
	res = solve(s, t)
	return
}

func solve(s string, t string) string {
	a := getFreq(s)
	b := getFreq(t)

	c := make([]int, 27)
	check := func(w int) bool {
		copy(c, a)
		for i := range 26 {
			need := w * b[i]
			if c[i]+c[26] < need {
				return true
			}
			c[26] -= max(need-c[i], 0)
		}
		return false
	}

	k := sort.Search(max(len(s), len(t))+1, check)
	k--
	// k是最小的满足条件的copy数
	buf := []byte(s)
	var cur int
	for i := range s {
		if buf[i] == '?' {
			for cur < 26 && a[cur] >= k*b[cur] {
				cur++
			}
			if cur == 26 {
				buf[i] = 'a'
				continue
			}
			if cur < 26 {
				buf[i] = byte('a' + cur)
				a[cur]++
				a[26]--
			}
		}
	}

	return string(buf)
}

func getFreq(s string) []int {
	res := make([]int, 27)
	for i := range s {
		if s[i] == '?' {
			res[26]++
		} else {
			res[int(s[i]-'a')]++
		}
	}
	return res
}
