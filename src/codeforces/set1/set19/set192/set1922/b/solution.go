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
	for i := 0; i < tc; i++ {
		fmt.Fprintln(writer, drive(reader))
	}
}

func drive(reader *bufio.Reader) int64 {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func readNum(reader *bufio.Reader) int {
	var x int
	fmt.Fscan(reader, &x)
	return x
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := range res {
		res[i] = readNum(reader)
	}
	return res
}

func solve(a []int) int64 {
	var res int
	slices.Sort(a)
	var sum int
	for i := 0; i < len(a); {
		j := i
		for i < len(a) && a[i] == a[j] {
			i++
		}
		res += count3(i - j)
		res += sum * count2(i-j)
		sum += i - j
	}

	return int64(res)
}

func count2(n int) int {
	return n * (n - 1) / 2
}

func count3(n int) int {
	return n * (n - 1) * (n - 2) / 6
}
