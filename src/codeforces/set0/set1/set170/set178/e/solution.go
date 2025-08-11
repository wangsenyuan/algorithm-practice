package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	a := make([][]int, n)
	for i := range n {
		a[i] = readNNums(reader, n)
	}
	return solve(a)
}

var dd = []int{-1, 0, 1, 0, -1}

const inf = 1 << 30

func solve(a [][]int) []int {
	n := len(a)

	b := make([][]int, n)
	for i := range n {
		b[i] = make([]int, n)
	}

	for i := 3; i+3 < n; i++ {
		for j := 3; j+3 < n; j++ {
			var sum int
			for dx := -3; dx <= 3; dx++ {
				for dy := -3; dy <= 3; dy++ {
					sum += a[i+dx][j+dy]
				}
			}
			if sum > 28 {
				b[i][j] = 1
			}
		}
	}

	a = b

	res := make([]int, 2)

	que := make([]int, n*n)

	check := func(x int, y int) int {
		a[x][y] = 0
		var head, tail int
		que[head] = x*n + y
		head++

		var sx, sy float64

		for tail < head {
			r, c := que[tail]/n, que[tail]%n
			sx += float64(r)
			sy += float64(c)
			tail++
			for i := range 4 {
				u, v := r+dd[i], c+dd[i+1]
				if u >= 0 && u < n && v >= 0 && v < n && a[u][v] == 1 {
					a[u][v] = 0
					que[head] = u*n + v
					head++
				}
			}
		}
		if head < 40 {
			return -1
		}

		sx /= float64(head)
		sy /= float64(head)

		var dist float64
		for i := range head {
			dx := float64(que[i]/n) - sx
			dy := float64(que[i]%n) - sy

			dist = max(dist, math.Sqrt(dx*dx+dy*dy))
		}

		var cnt int
		for i := range head {
			dx := float64(que[i]/n) - sx
			dy := float64(que[i]%n) - sy
			tmp := math.Sqrt(dx*dx + dy*dy)
			if tmp >= dist-3 {
				cnt++
			}
		}

		var cnt0 = int(math.Sqrt(float64(head)/3.14) * 6.8)

		if cnt >= cnt0 {
			return 0
		}
		return 1
	}

	for i := range n {
		for j := range n {
			if a[i][j] == 1 {
				shape := check(i, j)
				if shape >= 0 {
					res[shape]++
				}
			}
		}
	}

	return res
}
