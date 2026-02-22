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
		ok, res := drive(reader)
		if !ok {
			fmt.Fprintln(writer, "-1")
		} else {
			fmt.Fprintln(writer, res[0], res[1])
		}
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (bool, []int) {
	s := readString(reader)
	return solve(s)
}

func solve(s string) (bool, []int) {
	n := len(s)
	// level < 0, 基本就是把第一个 - 和 最后一个+交换
	var level int
	ok := true
	for i := range n {
		if s[i] == '+' {
			level++
		} else {
			level--
		}
		if level < 0 {
			ok = false
		}
	}

	if level < 0 {
		return false, nil
	}

	if ok {
		return true, []int{1, 1}
	}

	res := []int{-1, n}
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			res[1] = i
		} else {
			if res[0] < 0 {
				res[0] = i
			}
		}
	}

	buf := []byte(s)
	buf[res[0]], buf[res[1]] = buf[res[1]], buf[res[0]]

	for i := range n {
		if buf[i] == '+' {
			level++
		} else {
			level--
		}
		if level < 0 {
			return false, nil
		}
	}

	res[0]++
	res[1]++

	return true, res
}
