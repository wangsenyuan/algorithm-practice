package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		res, _ := drive(reader)
		fmt.Fprintln(writer, ouput(res))
	}
}

func ouput(ops []int) string {
	if len(ops) == 0 {
		return "0"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "%d\n", len(ops))
	for i, x := range ops {
		if i > 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(&b, "%d", x)
	}
	return b.String()
}

func drive(reader *bufio.Reader) ([]int, []int) {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i])
	}
	return solve(a), a
}

func solve(a []int) []int {
	n := len(a)
	absPrefix := make([]int64, n+1)
	suffix := make([]int64, n+1)
	for i, x := range a {
		absPrefix[i+1] = absPrefix[i] + int64(abs(x))
	}
	for i := n - 1; i >= 0; i-- {
		suffix[i] = suffix[i+1] + int64(a[i])
	}

	best := suffix[0]
	idx := -1
	for i, x := range a {
		if x <= 0 {
			continue
		}
		candidate := absPrefix[i] - int64(x) + suffix[i+1]
		if candidate > best {
			best = candidate
			idx = i
		}
	}
	if idx < 0 {
		return nil
	}

	// First make the prefix before idx entirely negative. The final operation
	// at idx then turns that prefix positive while making a[idx] negative.
	res := make([]int, 0, idx+1)
	flipped := false
	for i := idx - 1; i >= 0; i-- {
		positive := a[i] > 0
		if flipped {
			positive = !positive
		}
		if positive {
			res = append(res, i+1)
			flipped = !flipped
		}
	}
	res = append(res, idx+1)

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
