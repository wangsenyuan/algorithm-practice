package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		readNum(reader)
		s := readString(reader)
		res := solve(s)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

type data struct {
	r1  int
	r2  int
	cnt int
}

func solve(a string) int {
	n := len(a)
	var res int
	var prev []data

	for i := 0; i < n; i++ {
		var next []data
		if a[i] == '1' {
			// 只有 g[i][i] = 0
			for _, p := range prev {
				if p.r1 <= i && i <= p.r2 {
					next = append(next, data{i, i, p.cnt + 1})
					res = max(res, p.cnt+1)
					break
				}
			}
			if len(next) == 0 {
				res = max(res, 1)
				next = append(next, data{i, i, 1})
			}
		} else {
			// 除了 g[i][i], 其他都是0
			tmp := i
			for _, p := range prev {
				if p.r1 <= i-1 {
					tmp += p.cnt
				}
			}
			if tmp > 0 {
				next = append(next, data{0, i - 1, tmp})
				res = max(res, tmp)
			}

			tmp = n - 1 - i
			for _, p := range prev {
				if i+1 <= p.r2 {
					tmp += p.cnt
				}
			}
			if tmp > 0 {
				next = append(next, data{i + 1, n - 1, tmp})
				res = max(res, tmp)
			}
		}
		prev = next
	}

	return res
}
