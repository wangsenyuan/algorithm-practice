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
	buf := []byte(s)
	buf = slices.Compact(buf)
	s = string(buf)
	n := len(s)

	t := readString(reader)
	m, _ := strconv.Atoi(t)

	set := NewBitSet(1 << 26)

	a := make([]int, m)
	for i := range m {
		w := readString(reader)
		for _, ch := range w {
			a[i] |= 1 << (ch - 'a')
		}
		set.Set(a[i])
	}

	res := make(map[int]int)

	for i := range n {
		var mask int
		for j := i; j < n && (i == 0 || s[i-1] != s[j]); j++ {
			x := int(s[j] - 'a')
			if mask&(1<<x) == 0 {
				mask |= 1 << x
				if set.IsSet(mask | (1 << x)) {
					res[mask]++
				}
			}
		}
	}

	for _, v := range a {
		fmt.Fprintln(writer, res[v])
	}
}

type BitSet struct {
	sz  int
	arr []uint64
}

func NewBitSet(n int) *BitSet {
	sz := (n + 63) / 64
	arr := make([]uint64, sz)
	return &BitSet{sz, arr}
}

func (bs *BitSet) Set(p int) {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	bs.arr[i] |= 1 << uint64(j)
}

func (bs *BitSet) IsSet(p int) bool {
	i, j := p/64, p%64
	i = bs.sz - 1 - i
	return (bs.arr[i]>>uint64(j))&1 == 1
}
