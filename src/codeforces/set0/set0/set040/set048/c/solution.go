package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	unique, ans := process(reader)
	if unique {
		fmt.Println("unique")
		fmt.Println(ans)
	} else {
		fmt.Println("not unique")
	}
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

func process(reader *bufio.Reader) (unique bool, ans int) {
	n := readNum(reader)
	a := readNNums(reader, n)
	return solve(a)
}

func solve(a []int) (unique bool, ans int) {
	n := len(a)
	if n == 1 {
		return false, 0
	}
	check := func(alpha float64) int {
		cur := alpha
		for i := 0; i < n; i++ {
			dist := a[i]
			if i > 0 {
				dist -= a[i-1]
			}
			dist *= 10
			cur -= float64(dist)
			if cur >= 10 || cur < 0 {
				return -1
			}
			cur += alpha
		}
		return int(cur)/10 + a[n-1]
	}

	low := float64(a[0]) * 10
	hi := low + 10

	ans = -1

	for alpha := low; alpha < hi; alpha += 0.00001 {
		tmp := check(alpha)
		if tmp < 0 {
			continue
		}
		if ans >= 0 && ans != tmp {
			return false, 0
		}
		ans = tmp
	}
	if ans < 0 {
		return false, 0
	}
	return true, ans
}
