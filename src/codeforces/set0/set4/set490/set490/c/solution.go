package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, ok, s1, s2 := drive(reader)
	if !ok {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(s1)
	fmt.Println(s2)
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

func drive(reader *bufio.Reader) (s string, a int, b int, ok bool, x string, y string) {
	s = readString(reader)
	a, b = readTwoNums(reader)
	ok, x, y = solve(s, a, b)
	return
}

func add(x int, y int, mod int) int {
	x += y
	if x >= mod {
		x -= mod
	}
	return x
}

func mul(x int, y int, mod int) int {
	// x <= 10
	var res int

	for x > 0 {
		if x&1 == 1 {
			res = add(res, y, mod)
		}
		y = add(y, y, mod)
		x >>= 1
	}

	return res
}

func solve(s string, a int, b int) (bool, string, string) {
	n := len(s)

	bases := make([]int, n+1)
	bases[0] = 1
	for i := 1; i <= n; i++ {
		bases[i] = mul(10, bases[i-1], b)
	}

	suf := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - '0')
		suf[i] = add(mul(x, bases[n-i-1], b), suf[i+1], b)
	}

	if s[0] == '0' {
		return false, "", ""
	}

	var pref int
	for i := 0; i+1 < n; i++ {
		x := int(s[i] - '0')
		pref = add(mul(10, pref, a), x, a)
		if pref == 0 && suf[i+1] == 0 && s[i+1] != '0' {
			return true, s[:i+1], s[i+1:]
		}
	}

	return false, "", ""
}
