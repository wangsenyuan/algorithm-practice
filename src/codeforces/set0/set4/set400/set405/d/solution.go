package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, y := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(y)))
	for _, i := range y {
		buf.WriteString(fmt.Sprintf("%d ", i))
	}
	buf.WriteByte('\n')
	buf.WriteTo(os.Stdout)
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

func process(reader *bufio.Reader) (x []int, y []int) {
	n := readNum(reader)
	x = readNNums(reader, n)
	y = solve(x)
	return
}

func solve(x []int) []int {
	// n := len(x)
	s := 1_000_000
	// 一共这么多个block
	marked := make([]bool, s+1)

	for _, u := range x {
		marked[u] = true
	}

	var y []int

	var sum int

	for i := 1; i <= s; i++ {
		if marked[i] {
			if !marked[s-i+1] {
				y = append(y, s-i+1)
				marked[s-i+1] = true
			} else if i*2 <= s {
				if i*2 < s {
					sum += i - 1
					sum += s - i
				} else {
					sum += i - 1
				}
			}
		}
	}
	// i - 1 + s - i = s - 1
	// sum % (s - 1) == 0
	for i := 1; i*2 <= s && sum > 0; i++ {
		if !marked[i] {
			if i*2 < s {
				y = append(y, i)
				y = append(y, s-i+1)
				sum -= s - 1
			} else {
				y = append(y, i+1)
				sum -= s - (i + 1)
			}
		}
	}
	return y
}
