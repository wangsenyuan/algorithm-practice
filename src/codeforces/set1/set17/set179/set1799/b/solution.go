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
		var n int
		fmt.Fscan(reader, &n)
		a := make([]int, n)
		for i := range n {
			fmt.Fscan(reader, &a[i])
		}
		ok, res := solve(a)
		if !ok {
			fmt.Fprintln(writer, -1)
			continue
		}
		fmt.Fprintln(writer, len(res))
		for _, x := range res {
			fmt.Fprintln(writer, x[0], x[1])
		}
	}
}

func solve(a []int) (bool, [][]int) {
	// 目标是 a[?] = 2
	mn := slices.Min(a)
	mx := slices.Max(a)

	if mn == mx {
		return true, nil
	}
	if mn == 1 {
		return false, nil
	}
	pos := -1
	for i, v := range a {
		if v == mn {
			pos = i
			break
		}
	}

	var ans [][]int

	for {
		for i, v := range a {
			if v != mn {
				for v > mn {
					ans = append(ans, []int{i + 1, pos + 1})
					v = (v + mn - 1) / mn
				}
				mn = v
				pos = i
				a[i] = v
			}
		}
		if slices.Max(a) == mn {
			break
		}
	}

	return true, ans
}
