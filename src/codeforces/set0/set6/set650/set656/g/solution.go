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
	nums := readNums(reader)
	n, T := nums[0], nums[2]
	a := make([]string, n)
	for i := range n {
		a[i] = readString(reader)
	}
	res := solve(T, a)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	vals := strings.Split(s, " ")
	res := make([]int, len(vals))
	for i := range len(vals) {
		res[i], _ = strconv.Atoi(vals[i])
	}
	return res
}

func solve(T int, a []string) int {
	n := len(a)
	m := len(a[0])
	var res int
	for j := range m {
		var cnt int
		for i := range n {
			if a[i][j] == 'Y' {
				cnt++
			}
		}
		if cnt >= T {
			res++
		}
	}
	return res
}
