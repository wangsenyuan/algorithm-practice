package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)
	var buf bytes.Buffer
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
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
	buf, _ := reader.ReadBytes('\n')
	buf = bytes.TrimSpace(buf)
	return string(buf)
}

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	a := readNNums(reader, n)

	if solve(a) {
		return "YES"
	}
	return "NO"
}

func solve(a []int) bool {
	n := len(a)
	if n == 1 {
		return true
	}
	// if x != y and y > 0
	// a[i] = (x - y) * i + y * (n + 1)
	// a[2] - a[1] = x - y
	// a[3] - a[2] = x - y
	d := a[1] - a[0]
	if a[0] < d || (a[0]-d)%(n+1) != 0 {
		return false
	}
	y := (a[0] - d) / (n + 1)
	x := d + y
	if x < 0 {
		return false
	}
	for i := 1; i <= n; i++ {
		if a[i-1] != (x-y)*i+y*(n+1) {
			return false
		}
	}

	return true
}

func check1(a []int) bool {
	n := len(a)
	x := a[0]
	for i := 1; i <= n; i++ {
		if a[i-1] != x*i {
			return false
		}
	}
	return true
}

func check2(a []int) bool {
	n := len(a)
	y := a[n-1]

	for i := 1; i <= n; i++ {
		if a[i] != y*(n-i+1) {
			return false
		}
	}
	return true
}
