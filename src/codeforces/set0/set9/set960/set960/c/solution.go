package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	x, d := readTwoNums(reader)
	res := solve(x, d)
	if res == nil {
		fmt.Println(-1)
		return
	}
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for i := range res {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(x int, d int) []int {
	var res []int
	prev := -d + 1
	// 不需要不同的数字组成
	for x > 0 && len(res) < 1e4 {
		// x个非空集合，+1， 共x+1个集合
		// 1 << h >= x + 1
		h := bits.Len(uint(x + 1))
		// 1 << h >= x
		if 1<<h > x+1 {
			h--
		}
		// 1 << h <= x + 1
		cur := prev + d
		for range h {
			res = append(res, cur)
		}

		prev = cur
		x -= (1 << h) - 1
	}

	if x > 0 || len(res) > 1e4 {
		return nil
	}

	return res
}
