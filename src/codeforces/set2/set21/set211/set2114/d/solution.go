package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
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
	n := readNum(reader)
	monsters := make([][]int, n)
	for i := 0; i < n; i++ {
		monsters[i] = readNNums(reader, 2)
	}
	return solve(monsters)
}

const inf = 1 << 60

func solve(monsters [][]int) int {
	if len(monsters) == 1 {
		return 1
	}
	// n := len(monsters)
	// (x0, y0, x1, y1)
	border := []int{inf, inf, -inf, -inf}
	for _, cur := range monsters {
		border[0] = min(border[0], cur[0])
		border[1] = min(border[1], cur[1])
		border[2] = max(border[2], cur[0])
		border[3] = max(border[3], cur[1])
	}
	h := border[2] - border[0] + 1
	w := border[3] - border[1] + 1

	// h > 0 and w > 0
	cnt := make([]int, 4)
	pos := make([]int, 4)
	for i, cur := range monsters {
		x, y := cur[0], cur[1]
		if x == border[0] {
			cnt[0]++
			pos[0] = i
		}
		if x == border[2] {
			cnt[2]++
			pos[2] = i
		}
		if y == border[1] {
			cnt[1]++
			pos[1] = i
		}
		if y == border[3] {
			cnt[3]++
			pos[3] = i
		}
	}
	best := h * w

	n := len(monsters)

	check := func(move int) int {
		b := []int{inf, inf, -inf, -inf}
		for i, cur := range monsters {
			if i == move {
				continue
			}
			b[0] = min(b[0], cur[0])
			b[1] = min(b[1], cur[1])
			b[2] = max(b[2], cur[0])
			b[3] = max(b[3], cur[1])
		}
		h1 := b[2] - b[0] + 1
		w1 := b[3] - b[1] + 1
		if h1 > w1 {
			h1, w1 = w1, h1
		}
		tmp := h1 * w1

		if tmp >= n {
			return tmp
		}
		// tmp == n - 1
		return tmp + h1
	}

	for i := range 4 {
		if cnt[i] == 1 {
			best = min(best, check(pos[i]))
		}
	}

	return best
}
