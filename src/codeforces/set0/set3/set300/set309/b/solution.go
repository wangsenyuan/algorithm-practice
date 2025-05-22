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
	res, _, _, _ := process(reader)
	var buf bytes.Buffer
	for _, s := range res {
		buf.WriteString(s)
		buf.WriteString("\n")
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res []string, r int, c int, s string) {
	_, r, c = readThreeNums(reader)
	s = readString(reader)
	res = solve(r, c, s)
	return
}
func solve(r int, c int, s string) []string {
	words := strings.Split(s, " ")
	n := len(words)

	h := bits.Len(uint(r))

	next := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		next[i] = make([]int, h)
		for j := range h {
			next[i][j] = n
		}
	}
	var sum int
	for i, j := n-1, n; i >= 0; i-- {
		sum += len(words[i])
		for j > i && sum+j-1-i > c {
			sum -= len(words[j-1])
			j--
		}
		// j 有可能等于i
		next[i][0] = j
		for k := 1; k < h; k++ {
			next[i][k] = next[next[i][k-1]][k-1]
		}
	}

	best := []int{0, 0, 0}

	for i := 0; i < n; i++ {
		if len(words[i]) > c {
			continue
		}
		d := r
		j := i
		for k := h - 1; k >= 0; k-- {
			if d&(1<<k) > 0 {
				j = next[j][k]
			}
		}
		if j-i > best[0] {
			best[0] = j - i
			best[1] = i
			best[2] = j
		}
	}

	var ans []string

	for i := best[1]; i < best[2]; {
		j := next[i][0]
		if j == i {
			break
		}
		ans = append(ans, strings.Join(words[i:j], " "))
		i = j
	}

	return ans
}
