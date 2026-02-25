package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	drive(reader, writer)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader, writer *bufio.Writer) {
	solve(reader, writer)
}

func solve(reader *bufio.Reader, writer *bufio.Writer) {
	s := readString(reader)
	n := len(s)

	wp := make(map[int32]int32)
	var last [26]int32
	for i := range 26 {
		last[i] = -1
	}

	var order [26]struct{ d, last int }

	cmp := func(a, b struct{ d, last int }) int {
		return b.last - a.last
	}
	for r := range n {
		last[s[r]-'a'] = int32(r)

		size := 0
		for d := range 26 {
			if last[d] >= 0 {
				order[size].d = d
				order[size].last = int(last[d])
				size++
			}
		}

		slices.SortFunc(order[:size], cmp)

		var C int32
		for i := 0; i < size; i++ {
			C |= 1 << order[i].d
			if r < n-1 && (C>>(s[r+1]-'a'))&1 == 1 {
				break
			}
			wp[C]++
		}
	}

	t := readString(reader)
	m, _ := strconv.Atoi(t)

	for range m {
		cur := readString(reader)
		var C int32
		for _, ch := range cur {
			C |= 1 << (ch - 'a')
		}
		fmt.Fprintln(writer, wp[C])
	}
}
