package main

import (
	"bufio"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	process(reader, writer)
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

func process(reader *bufio.Reader, writer *bufio.Writer) {
	n := readNum(reader)
	queries := make([][]int, n)
	for i := range n {
		s, _ := reader.ReadBytes('\n')
		if s[0] != '3' {
			queries[i] = make([]int, 3)
			queries[i][0] = int(s[0] - '0')
		} else {
			queries[i] = make([]int, 2)
			queries[i][0] = 3
		}
		pos := 2
		for j := 1; j < len(queries[i]); j++ {
			pos = readInt(s, pos, &queries[i][j]) + 1
		}
	}
	solve(queries, writer)
}

const H = 64

// Custom int to string conversion to avoid fmt.Sprintf overhead
func intToString(n int) string {
	if n == 0 {
		return "0"
	}

	// Handle negative numbers
	negative := false
	if n < 0 {
		negative = true
		n = -n
	}

	// Count digits
	digits := 0
	temp := n
	for temp > 0 {
		digits++
		temp /= 10
	}

	// Allocate string
	start := 0
	if negative {
		digits++
		start = 1
	}

	result := make([]byte, digits)
	if negative {
		result[0] = '-'
	}

	// Convert to string
	for i := digits - 1; i >= start; i-- {
		result[i] = byte('0' + n%10)
		n /= 10
	}

	return string(result)
}

func solve(queries [][]int, writer *bufio.Writer) {
	getLevel := func(x int) int {
		return bits.Len(uint(x))
	}

	offset := make([]int, H)

	shift := func(i int, k int) {
		m := 1 << (i - 1)
		k %= m
		if k < 0 {
			k += m
		}
		offset[i] += k
		offset[i] %= m
	}

	getOriginal := func(v int, i int) int {
		sz := 1 << (i - 1)
		p := v - sz
		return sz + ((p-offset[i])%sz+sz)%sz
	}

	for _, cur := range queries {
		if cur[0] == 1 {
			x, k := cur[1], cur[2]
			shift(getLevel(x), k)
		} else if cur[0] == 2 {
			x, k := cur[1], cur[2]
			i := getLevel(x)
			for i < H {
				shift(i, k)
				k <<= 1
				i++
			}
		} else {
			x := cur[1]
			i := getLevel(x)
			sz := 1 << (i - 1)
			p := x - sz
			v := sz + ((p+offset[i])%sz+sz)%sz

			// Build output array first
			for i >= 1 {
				z := getOriginal(v, i)
				writer.WriteString(intToString(z))
				writer.WriteByte(' ')
				v >>= 1
				i--
			}

			writer.WriteByte('\n')
		}
	}
}
