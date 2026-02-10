package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	tc := readNums(reader)[0]

	for range tc {
		tmp := readNums(reader)
		n, m := tmp[0], tmp[1]
		ans := solve(n, m)
		for _, x := range ans {
			fmt.Fprintln(writer, x)
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s, _ := reader.ReadString('\n')
	fields := strings.Fields(s)
	nums := make([]int, len(fields))
	for i, field := range fields {
		nums[i], _ = strconv.Atoi(field)
	}
	return nums
}

func solve(n int, m int) []string {
	buf := make([][]byte, n)

	ans := make([]string, n)

	for i := range n {
		buf[i] = make([]byte, m)
		for j := range m {
			if i == 0 || j == m-1 {
				buf[i][j] = 'B'
			} else {
				buf[i][j] = 'W'
			}
		}
		ans[i] = string(buf[i])
	}

	return ans
}
