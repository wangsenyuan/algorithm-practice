package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.8f\n", res)
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
	nums := readNNums(reader, 5)
	return solve(nums[0], nums[1], nums[2], nums[3], nums[4])
}

func solve(l int, d int, v int, g int, r int) float64 {
	// t = d / v
	// t0 = m * (g + r)
	// t1 = t0 + g
	// 如果t >= t0 and t < t1
	t2 := (d + v - 1) / v
	// t2 * v >= d
	if t2*v > d {
		t2--
	}
	// t2 * v <= d
	// 在t2时刻到达了位置 t2 * v
	t0 := t2 / (g + r) * (g + r)
	// 在 t0时刻到达了位置 t0 * v, 且此时正好是绿灯
	t1 := t0 + g
	// 在t1时刻，变成红灯的时候，还没有通过
	if t1*v <= d {
		t3 := float64(l-d) / float64(v)
		// 要等到下次绿灯，才能通行
		return float64(t1+r) + t3
	}
	return float64(l) / float64(v)
}
