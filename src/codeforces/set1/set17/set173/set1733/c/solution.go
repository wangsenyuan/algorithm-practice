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
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
		}
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
func solve(A []int) (res [][]int) {
	n := len(A)
	if n == 1 {
		return
	}

	if A[0] != A[n-1] {
		res = append(res, []int{1, n})
		if (A[0]+A[n-1])&1 == 1 {
			A[n-1] = A[0]
		} else {
			A[0] = A[n-1]
		}
	}
	// A[0] = A[n-1]

	for i := 1; i < n-1; i++ {
		if A[i] != A[0] {
			if (A[i]+A[0])&1 == 1 {
				res = append(res, []int{1, i + 1})
			} else {
				res = append(res, []int{i + 1, n})
			}
		}
	}

	return
}
