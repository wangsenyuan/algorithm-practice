package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, x := range drive(reader) {
		fmt.Fprintln(writer, x)
	}
}

func drive(reader *bufio.Reader) []int {
	var t int
	fmt.Fscan(reader, &t)
	ops := make([]string, t)
	args := make([]string, t)
	for i := range t {
		fmt.Fscan(reader, &ops[i], &args[i])
	}
	return solve(ops, args)
}

func solve(ops []string, args []string) []int {
	freq := make(map[string]int)

	rep := func(num int) string {
		buf := make([]byte, 18)
		for i := range 18 {
			buf[i] = '0'
		}
		for i := 17; num > 0; num /= 10 {
			if (num%10)&1 == 1 {
				buf[i] = '1'
			}
			i--
		}
		return string(buf)
	}

	add := func(num int) {
		freq[rep(num)]++
	}

	rem := func(num int) {
		freq[rep(num)]--
	}

	query := func(w string) int {
		buf := make([]byte, 18)
		for i := range 18 {
			buf[i] = '0'
		}
		m := len(w)
		copy(buf[18-m:], w)
		return freq[string(buf)]
	}

	var ans []int
	for i, op := range ops {
		if op[0] == '?' {
			ans = append(ans, query(args[i]))
		} else {
			num, _ := strconv.Atoi(args[i])
			if op[0] == '+' {
				add(num)
			} else {
				rem(num)
			}
		}
	}
	return ans
}
