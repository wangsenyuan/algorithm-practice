package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, ok, res := drive(reader)
		if !ok {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, len(res))
			s := fmt.Sprintf("%v", res)
			fmt.Fprintln(writer, s[1:len(s)-1])
		}
	}
}

func drive(reader *bufio.Reader) (a []int, ok bool, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	ok, res = solve(slices.Clone(a))
	return
}

func solve(a []int) (ok bool, res []int) {

	// n := len(a)

	for {
		// slices.Sort(a)
		u := slices.Min(a)
		v := slices.Max(a)
		// slices.Reverse(a)
		if (u+v)%2 == 1 {
			return false, nil
		}
		if v == 0 {
			break
		}
		x := (u + v) / 2
		res = append(res, x)
		for i := range a {
			a[i] = abs(a[i] - x)
		}
	}

	return true, res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
