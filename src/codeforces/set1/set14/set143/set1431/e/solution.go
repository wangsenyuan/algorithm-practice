package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var tc int
	fmt.Fscan(reader, &tc)
	for range tc {
		_, _, res := drive(reader)
		s := fmt.Sprintf("%v", res)
		fmt.Sprintln(writer, s[1:len(s)-1])
	}
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) (a []int, b []int, res []int) {
	var n int
	fmt.Fscan(reader, &n)
	a = readNNums(reader, n)
	b = readNNums(reader, n)
	res = solve(a, b)
	return
}

type player struct {
	id    int
	level int
}

func solve(a []int, b []int) []int {
	n := len(a)

	check := func(x int) bool {
		// abs(a[i] - b[i]) >= x holds
		for s := range n {
			ok := true

			for i := range n {
				if abs(a[i]-b[(i+s)%n]) < x {
					ok = false
					break
				}
			}

			if ok {
				return false
			}
		}
		return true
	}

	r := max(abs(a[n-1]-b[0]), abs(a[0]-b[n-1])) + 1

	x := sort.Search(r, check) - 1

	res := make([]int, n)

	for s := range n {
		ok := true
		for i := range n {
			if abs(a[i]-b[(i+s)%n]) < x {
				ok = false
				break
			}
		}
		if ok {
			for i := range n {
				res[i] = (i+s)%n + 1
			}
			break
		}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
