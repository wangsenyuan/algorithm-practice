package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	n, k := readTwoNums(reader)
	t := readString(reader)
	return solve(n, k, t)
}

const mod = 1e9 + 7

func add(a, b int) int {
	return (a + b) % mod
}

func sub(a, b int) int {
	return add(a, mod-b)
}

func solve(n int, k int, t string) int {
	m := len(t)

	seen := make([]bool, k)
	var arr []byte
	for i := m - 1; i >= 0; i-- {
		x := int(t[i] - 'a')
		if !seen[x] {
			seen[x] = true
			arr = append(arr, t[i])
		}
	}

	for j := range k {
		if !seen[j] {
			arr = append(arr, byte('a'+j))
		}
	}

	slices.Reverse(arr)

	buf := make([]byte, m+n)
	copy(buf, t)

	for i := m; i < len(buf); i++ {
		buf[i] = arr[(i-m)%len(arr)]
	}

	fp := make([]int, k)

	for i := 0; i < m+n; i++ {
		x := int(buf[i] - 'a')
		var sum int
		for j := range k {
			sum = add(sum, fp[j])
		}
		cur := sub(add(sum, 1), fp[x])
		fp[x] = add(fp[x], cur)
	}

	// empty
	res := 1
	for i := range k {
		res = add(res, fp[i])
	}
	return res
}
