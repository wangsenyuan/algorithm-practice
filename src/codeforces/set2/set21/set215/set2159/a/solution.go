package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var tc int
	fmt.Fscan(reader, &tc)

	ask := func(arr []int) int {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("? %d", len(arr)))
		for _, i := range arr {
			buf.WriteString(fmt.Sprintf(" %d", i))
		}
		buf.WriteByte('\n')
		buf.WriteTo(os.Stdout)
		var res int
		fmt.Fscan(reader, &res)
		return res
	}

	for range tc {
		var n int
		fmt.Fscan(reader, &n)
		res := solve(n, ask)
		var buf bytes.Buffer
		buf.WriteString("!")
		for _, v := range res {
			buf.WriteString(fmt.Sprintf(" %d", v))
		}
		buf.WriteByte('\n')
		buf.WriteTo(os.Stdout)
	}
}

func solve(n int, ask func(arr []int) int) []int {
	ids := make([]int, 2*n)
	for i := range 2 * n {
		ids[i] = i + 1
	}
	res := make([]int, 2*n+1)

	var s []int

	x := ask(ids[:2])
	if x == 0 {
		s = ids[:2]
		ids = ids[2:]
	} else {
		s = []int{ids[0], ids[2]}
		ids = ids[2:]
		ids[0] = 2
	}

	for len(ids) > 0 {
		s1 := append(s, ids[0])
		x := ask(s1)
		if x != 0 {
			res[ids[0]] = x
		} else {
			s = s1
		}
		ids = ids[1:]
	}
	s = s[:0]
	var unknown []int
	for i := 1; i <= 2*n; i++ {
		if res[i] != 0 {
			s = append(s, i)
		} else {
			unknown = append(unknown, i)
		}
	}

	for len(unknown) > 0 {
		s1 := append(s, unknown[0])
		x := ask(s1)
		res[unknown[0]] = x
		unknown = unknown[1:]
	}

	return res[1:]
}
