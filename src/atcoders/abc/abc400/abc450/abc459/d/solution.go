package main

import (
	"bufio"
	"bytes"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		var s string
		fmt.Fscan(reader, &s)
		res := solve(s)
		if len(res) > 0 {
			buf.WriteString("Yes\n")
			buf.WriteString(res)
			buf.WriteByte('\n')
		} else {
			buf.WriteString("No\n")
		}
	}

	buf.WriteTo(os.Stdout)
}

func solve(s string) string {
	n := len(s)
	freq := make([]int, 26)
	for i := range n {
		freq[int(s[i]-'a')]++
	}
	w := slices.Max(freq)
	if w > (n+1)/2 {
		return ""
	}

	type pair struct {
		first  int
		second int
	}

	var arr []pair
	for i, v := range freq {
		if v > 0 {
			arr = append(arr, pair{v, i})
		}
	}
	slices.SortFunc(arr, func(a, b pair) int {
		return cmp.Or(b.first-a.first, a.second-b.second)
	})

	res := make([]byte, n)
	var pos int
	for _, cur := range arr {
		c := byte(cur.second + 'a')
		for range cur.first {
			res[pos] = c
			pos += 2
			if pos >= n {
				pos = 1
			}
		}
	}

	return string(res)
}
