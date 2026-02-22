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

	tc := readString(reader)
	tc_no, _ := strconv.Atoi(tc)
	for range tc_no {
		res := drive(reader)
		writer.WriteString(fmt.Sprintf("%d\n", res))
	}

}

func drive(reader *bufio.Reader) int {
	readString(reader)
	a := make([]string, 2)
	for i := range 2 {
		a[i] = readString(reader)
	}
	return solve(a)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func solve(a []string) int {
	n := len(a[0])

	var res int

	for i := 0; i < n; {
		j := i
		for i < n && a[0][i] == a[0][j] {
			i++
		}
		res += i - j - 1
	}

	return res
}
