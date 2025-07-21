package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(res)
		buf.WriteByte('\n')
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
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) string {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, m)
	if solve(a, b) {
		return "YES"
	}
	return "NO"
}

func solve(a []int, b []int) bool {
	sort.Ints(b)
	n := len(a)
	m := len(b)
	// a[i] = b[j] - a[i] and a[i] <= a[i+1]
	// => b[j] - a[i] <= a[i+1] => b[j] <= a[i] + a[i+1]
	if a[n-1] < b[m-1]-a[n-1] {
		a[n-1] = b[m-1] - a[n-1]
	}
	for i := n - 2; i >= 0; i-- {
		j := sort.Search(m, func(j int) bool {
			return b[j] >= a[i]+a[i+1]
		})
		// b[j] >= a[i] + a[i+1]
		if j == m || b[j] > a[i]+a[i+1] {
			j--
		}
		// b[j] <= a[i] + a[i+1]
		// b[j] - a[i] <= a[i+1]
		if j < 0 {
			// 不存在 b[j] 使的， b[j] - a[i] <= a[i+1]
			if a[i] <= a[i+1] {
				// 可以不改变
				continue
			}
			return false
		}
		// b[j] - a[i] <= a[i+1] 成立
		if a[i] > a[i+1] || a[i] < b[j]-a[i] {
			a[i] = b[j] - a[i]
		}
	}
	return true
}
