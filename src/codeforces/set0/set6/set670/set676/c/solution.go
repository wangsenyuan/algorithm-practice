package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func drive(reader *bufio.Reader) int {
	_, k := readTwoNums(reader)
	s := readString(reader)
	return solve(k, s)
}

func solve(k int, s string) int {
	res1 := solve1(k, s, 'a')
	res2 := solve1(k, s, 'b')
	return max(res1, res2)
}

func solve1(k int, s string, x byte) int {
	var pref int
	n := len(s)
	pos := make([]int, n+1)
	pos[0] = -1
	w := 0
	var best int
	for r := range n {
		if s[r] == x {
			pref++
		}
		tmp := r + 1 - pref
		for w < tmp {
			w++
			pos[w] = r
		}
		tmp -= k
		if tmp < 0 {
			best = r + 1
		} else {
			l := pos[tmp]
			best = max(best, r-l)
		}
	}
	return best
}
