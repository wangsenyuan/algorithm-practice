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
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	var n int
	fmt.Fscan(reader, &n)
	d := make([]int64, n)
	for i := range d {
		fmt.Fscan(reader, &d[i])
	}
	return solve(d)
}

func solve(d []int64) int64 {
	slices.Sort(d)
	var ans int64
	var sum int64
	for i, v := range d {
		ans -= int64(i)*v - sum
		sum += v
	}
	ans += int64(d[len(d)-1])
	return ans
}
