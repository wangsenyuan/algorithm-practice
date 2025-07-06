package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.6f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n := readNum(reader)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 2)
	}
	return solve(points)
}

type point struct {
	x int
	y int
}

func ccw(a, b, c []int) int {
	return (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
}

func solve(points [][]int) float64 {
	n := len(points)
	var best int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x, y := -1, -1
			for k := 0; k < n; k++ {
				if k == i || k == j {
					continue
				}
				tmp := ccw(points[i], points[j], points[k])
				if tmp < 0 {
					x = max(x, -tmp)
				} else {
					y = max(y, tmp)
				}
			}
			if x >= 0 && y >= 0 {
				best = max(best, x+y)
			}
		}
	}

	return float64(best) / 2
}
