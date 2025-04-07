package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	a, b, c := readThreeNums(reader)
	n, roots := solve(a, b, c)
	fmt.Println(n)
	if n <= 0 {
		return
	}
	for _, f := range roots {
		fmt.Printf("%.10f\n", f)
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

func solve(a int, b int, c int) (int, []float64) {
	// -b +/- sqrt(b^2 - 4ac)
	if b*b < 4*a*c {
		return 0, nil
	}

	if a == 0 {
		// b * x + c = 0
		if b == 0 {
			if c == 0 {
				return -1, nil
			}
			return 0, nil
		}
		return 1, []float64{-float64(c) / float64(b)}
	}

	if b*b == 4*a*c {
		return 1, []float64{-float64(b) / float64(2*a)}
	}

	w := math.Sqrt(float64(b*b - 4*a*c))
	res := []float64{(float64(-b) - w) / float64(2*a), (float64(-b) + w) / float64(2*a)}
	if res[0] > res[1] {
		res[0], res[1] = res[1], res[0]
	}
	return 2, res
}
