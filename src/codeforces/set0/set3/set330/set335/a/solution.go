package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) string {
	var s string
	var n int
	fmt.Fscan(reader, &s, &n)
	return solve(s, n)
}

func solve(s string, n int) string {
	// 一共m个sheet, 假设sheet上n长度的x, 那么从 m个x中能够得到s
	freq := make([]int, 26)
	for _, c := range s {
		freq[c-'a']++
	}
	// m * freq_x[i] >= freq[i]

	check := func(m int) bool {
		if m == 0 {
			return false
		}
		// m个sheet
		var sum int
		for _, v := range freq {
			if v > 0 {
				sum += (v + m - 1) / m
			}
		}
		return sum <= n
	}

	m := sort.Search(1<<20, check)

	if m == 1<<20 {
		return "-1"
	}

	var buf []byte
	for i, v := range freq {
		if v > 0 {
			k := (v + m - 1) / m
			for range k {
				buf = append(buf, byte('a'+i))
			}
		}
	}
	for len(buf) < n {
		buf = append(buf, 'x')
	}

	return fmt.Sprintf("%d\n%s", m, string(buf))
}
