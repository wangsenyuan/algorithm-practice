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
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
}
func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(a []int) int {
	freq := make(map[[3]int]int)

	n := len(a)

	var res int
	for i := 0; i+3 <= n; i++ {
		cur := [3]int{a[i], a[i+1], a[i+2]}
		w := freq[cur]
		for j := range 3 {
			v := cur[j]
			cur[j] = -1
			res += freq[cur] - w
			cur[j] = v
		}
		for j := range 3 {
			v := cur[j]
			cur[j] = -1
			freq[cur]++
			cur[j] = v
		}
		freq[cur]++
	}
	return res
}

func solve1(a []int) int {
	var res int

	mem := make(map[Record]int)

	n := len(a)
	for i := 2; i < n; i++ {
		rec := [3]int{a[i-2], a[i-1], 0}
		res += mem[rec]
		rec = [3]int{a[i-2], 0, a[i]}
		res += mem[rec]
		rec = [3]int{0, a[i-1], a[i]}
		res += mem[rec]
		rec = [3]int{a[i-2], a[i-1], a[i]}
		res -= 3 * mem[rec]

		mem[[3]int{a[i-2], a[i-1], 0}]++
		mem[[3]int{a[i-2], 0, a[i]}]++
		mem[[3]int{0, a[i-1], a[i]}]++
		mem[[3]int{a[i-2], a[i-1], a[i]}]++
	}

	return res
}

type Record [3]int
