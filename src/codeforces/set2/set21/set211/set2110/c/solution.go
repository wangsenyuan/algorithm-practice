package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		_, _, res := process(reader)
		if len(res) == 0 {
			buf.WriteString("-1\n")
		} else {
			for _, x := range res {
				buf.WriteString(strconv.Itoa(x))
				buf.WriteByte(' ')
			}
			buf.WriteByte('\n')
		}
	}
	fmt.Print(buf.String())
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
	bs, _ := reader.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func process(reader *bufio.Reader) (l []int, r []int, res []int) {
	n := readNum(reader)
	d := readNNums(reader, n)
	l = make([]int, n)
	r = make([]int, n)
	for i := range n {
		l[i], r[i] = readTwoNums(reader)
	}
	res = solve(d, l, r)
	return
}

func solve(d []int, l []int, r []int) []int {

	n := len(d)
	stack := make([]int, n)
	var top int

	var pos int
	for i := 0; i < n; i++ {
		if d[i] == -1 {
			stack[top] = i
			top++
		} else {
			pos += d[i]
		}

		if pos+top < l[i] || pos > r[i] {
			return nil
		}
		for pos < l[i] {
			pos++
			d[stack[top-1]] = 1
			top--
		}
		for pos+top > r[i] {
			d[stack[top-1]] = 0
			top--
		}
	}

	for top > 0 {
		d[stack[top-1]] = 0
		top--
	}
	return d
}

func solve1(d []int, l []int, r []int) []int {
	n := len(d)
	ans := make([]int, n)

	copy(ans, d)

	if l[0] > 1 {
		return nil
	}

	var lo, hi int

	for i := 0; i < n; i++ {
		if ans[i] != -1 {
			lo += ans[i]
			hi += ans[i]
		} else {
			// 尽量增加高度
			hi++
		}
		lo = max(lo, l[i])
		hi = min(hi, r[i])
		if lo > hi {
			return nil
		}
		l[i] = lo
		r[i] = hi
	}

	lo, hi = l[n-1], r[n-1]
	for i := n - 1; i >= 0; i-- {
		lo = max(lo, l[i])
		hi = min(hi, r[i])
		if lo > hi {
			return nil
		}
		if d[i] != -1 {
			lo -= d[i]
			hi -= d[i]
		} else {
			var a, b int
			if i > 0 {
				a = l[i-1]
				b = r[i-1]
			}
			ans[i] = 0
			if max(a, lo-1) <= min(b, hi-1) {
				ans[i] = 1
				lo--
				hi--
			}
		}
	}
	return ans
}
