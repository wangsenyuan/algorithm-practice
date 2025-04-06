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
	for tc > 0 {
		tc--
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func process(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	b := readNNums(reader, n)
	return solve(k, a, b)
}

func solve(k int, a []int, b []int) int {
	// k > 0
	n := len(a)

	if sum(a) <= k {
		return 0
	}

	arr := make([]int, 1+n*2)

	for i := 0; i < 2*n; i++ {
		arr[i+1] = a[i%n] - b[i%n]
		arr[i+1] += arr[i]
	}

	pos := n
	for i := n + 1; i <= 2*n; i++ {
		if arr[i] < arr[pos] {
			pos = i
		}
	}

	stack := make([]int, n*2+1)

	check := func(x int) bool {
		var cnt int
		l, r := 1, 0
		for i := pos; i >= pos-n; i-- {
			if l <= r && stack[l]-i > x {
				l++
			}
			if pos-i >= x {
				cnt += max(0, arr[stack[l]]-arr[i])
			}
			if cnt > k {
				return false
			}
			for l <= r && arr[stack[r]] > arr[i] {
				r--
			}
			r++
			stack[r] = i
		}
		return true
	}

	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func sum(arr []int) int {
	var res int
	for _, num := range arr {
		res += num
	}
	return res
}
