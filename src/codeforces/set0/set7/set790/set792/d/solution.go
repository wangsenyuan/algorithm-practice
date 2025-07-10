package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	buf.WriteTo(os.Stdout)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) []int {
	n, q := readTwoNums(reader)
	vertices := make([]int, q)
	moves := make([]string, q)
	for i := 0; i < q; i++ {
		vertices[i] = readNum(reader)
		moves[i] = readString(reader)
	}
	return solve(n, vertices, moves)
}

func solve(n int, vertices []int, moves []string) []int {

	root := n/2 + 1

	play := func(u int, s string) int {
		for i := range s {
			if s[i] == 'U' {
				if u == root {
					continue
				}
				lo := bits.TrailingZeros(uint(u))
				if (u>>(lo+1))&1 == 0 {
					// 左子树
					u ^= 1<<(lo+1) ^ (1 << lo)
				} else {
					// 右子树
					u ^= 1 << lo
				}
			} else if s[i] == 'L' {
				if u&1 == 1 {
					continue
				}
				lo := bits.TrailingZeros(uint(u))
				u ^= (1 << lo) ^ (1 << (lo - 1))
			} else {
				if u&1 == 1 {
					continue
				}
				lo := bits.TrailingZeros(uint(u))
				u ^= (1 << (lo - 1))
			}
		}
		return u
	}

	ans := make([]int, len(moves))

	for i, s := range moves {
		ans[i] = play(vertices[i], s)
	}
	return ans
}
