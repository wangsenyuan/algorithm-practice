package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
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

func process(reader *bufio.Reader) int {
	first := readNNums(reader, 3)
	second := readNNums(reader, 3)
	return solve(first, second)
}

func solve(first []int, second []int) int {
	// a x + by + c = 0
	l1 := convert(first)
	l2 := convert(second)
	if l1.kind == -2 || l2.kind == -2 {
		return 0
	}
	if l1.kind == -1 || l2.kind == -1 {
		// 有一个是平面
		return -1
	}
	if l1.kind == l2.kind {
		// 平行线
		if l1.kind <= 1 {
			if l1.val == l2.val {
				return -1
			}
			return 0
		}
		if l1.slop == l2.slop {
			if l1.val == l2.val {
				return -1
			}
			return 0
		}
		return 1
	}

	return 1
}

type pair struct {
	first  int
	second int
}

type line struct {
	kind int // 0 for vertical, 1 for horizontal, 2 for slop, -1 for invalid
	slop pair
	val  pair
}

func convert(arr []int) line {
	a, b, c := arr[0], arr[1], arr[2]
	if a == 0 {
		if b == 0 {
			if c == 0 {
				// 平面
				return line{kind: -1}
			}
			// 无效的
			return line{kind: -2}
		}

		s := -sign(b * c)
		g := gcd(abs(c), abs(b))
		val := pair{s * abs(c) / g, abs(b) / g}

		return line{kind: 1, slop: pair{0, 1}, val: val}
	}
	if b == 0 {
		s := -sign(a * c)
		g := gcd(abs(a), abs(c))
		val := pair{s * abs(c) / g, abs(a) / g}
		return line{kind: 0, slop: pair{1, 0}, val: val}
	}
	g := gcd(abs(a), abs(b))
	s := -sign(a * b)
	slop := pair{s * abs(a) / g, abs(b) / g}
	// when x = 0, b * y + c = 0
	g = gcd(abs(c), abs(b))
	s = -sign(c * b)
	val := pair{s * abs(c) / g, abs(b) / g}
	return line{kind: 2, slop: slop, val: val}
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	return 1
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
