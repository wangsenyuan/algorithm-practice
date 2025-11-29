package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	buf, _ := reader.ReadString('\n')
	return strings.TrimSpace(buf)
}

func drive(reader *bufio.Reader) string {
	a := readString(reader)
	b := readString(reader)
	return solve(a, b)
}

func solve(oleg string, igor string) string {
	a := []byte(oleg)
	b := []byte(igor)
	n := len(a)

	slices.SortFunc(a, func(i, j byte) int {
		return int(i-'a') - int(j-'a')
	})
	slices.SortFunc(b, func(i, j byte) int {
		return int(i-'a') - int(j-'a')
	})

	buf := make([]byte, n)

	p0 := []int{0, (n+1)/2 - 1}
	p1 := []int{n - n/2, n - 1}

	l, r := 0, n-1

	var turn int

	for l <= r {
		if turn == 0 {
			if b[p1[1]] > a[p0[0]] {
				buf[l] = a[p0[0]]
				l++
				p0[0]++
			} else {
				buf[r] = a[p0[1]]
				r--
				p0[1]--
			}
		} else {
			if a[p0[0]] < b[p1[1]] {
				// 最小的，比自己最大的还大
				buf[l] = b[p1[1]]
				p1[1]--
				l++
			} else {
				buf[r] = b[p1[0]]
				p1[0]++
				r--
			}
		}

		turn ^= 1
	}

	return string(buf)
}
