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
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		if x {
			buf.WriteString("Yes\n")
		} else {
			buf.WriteString("No\n")
		}
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
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return s
}

func process(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	a := readNNums(reader, n)
	queries := make([][]int, m)
	for i := 0; i < m; i++ {
		queries[i] = readNNums(reader, 2)
	}
	return solve(a, queries)
}

type pair struct {
	first  int
	second int
}

func solve(a []int, queries [][]int) []bool {
	n := len(a)

	var arr []pair
	for i := 0; i < n; i++ {
		if i == 0 || a[i] != arr[len(arr)-1].first {
			arr = append(arr, pair{a[i], i})
		}
	}

	prev := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = -2
	}
	prev[0] = -1

	if len(arr) > 1 {
		prev[arr[1].second] = -1
	}
	for i := 2; i < len(arr); i++ {
		prev[arr[i].second] = arr[i-1].second
		if sign(arr[i].first-arr[i-1].first) == sign(arr[i-1].first-arr[i-2].first) {
			prev[arr[i].second] = prev[arr[i-1].second]
		}
	}

	for i := 1; i < n; i++ {
		if prev[i] == -2 {
			prev[i] = prev[i-1]
		}
	}

	check := func(l int, r int) bool {
		if l >= prev[r] {
			return true
		}
		// l < prev[r]
		if l >= prev[prev[r]] && a[prev[r]] > a[r] {
			return true
		}
		return false
	}

	ans := make([]bool, len(queries))

	for i, cur := range queries {
		ans[i] = check(cur[0]-1, cur[1]-1)
	}
	return ans
}

func sign(num int) int {
	if num > 0 {
		return 1
	}
	if num < 0 {
		return -1
	}
	return 0
}
