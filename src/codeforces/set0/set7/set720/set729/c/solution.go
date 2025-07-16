package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
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
	nums := readNNums(reader, 4)
	n, k, s, t := nums[0], nums[1], nums[2], nums[3]
	cars := make([][]int, n)
	for i := 0; i < n; i++ {
		cars[i] = readNNums(reader, 2)
	}
	g := readNNums(reader, k)
	return solve(s, t, cars, g)
}

func solve(s int, t int, cars [][]int, g []int) int {
	// 在油足够的情况下，应该尽快的到达下一个地点
	g = append(g, 0)
	g = append(g, s)
	sort.Ints(g)

	slices.SortFunc(cars, func(a, b []int) int {
		return cmp.Or(a[0]-b[0], b[1]-a[1])
	})

	var arr [][]int

	for i := 0; i < len(cars); i++ {
		cur := cars[i]
		if len(arr) > 0 && (cur[0] == arr[len(arr)-1][0] || cur[1] <= arr[len(arr)-1][1]) {
			// 这个更贵，且容量更少，不划算
			continue
		}
		arr = append(arr, cur)
	}

	n := len(g)

	// c[i] < c[i+1], and v[i] < v[i+1]

	check := func(full int) bool {
		var sum int
		for i := 0; i+1 < n; i++ {
			dist := g[i+1] - g[i]
			// 当前有full的油，要移动到下一个位置，花费最小的时间
			// 如果是正常速度，那么就是在1公里，花费1升油，需要2分钟
			// 如果是加速行驶, 那么就是在1公里, 花费2升油，需要1分钟
			// 假设其中x是加速，y正常速度
			// x + y = dist
			// 2 * x + y <= full
			// x = full - dist
			if full < dist {
				return false
			}
			if 2*dist <= full {
				// 可以一直加速
				sum += dist
			} else {
				// full >= dist 始终成立，这个可以二分的左端点
				x := full - dist
				y := dist - x
				sum += x + 2*y
			}
		}
		return sum <= t
	}

	if !check(arr[len(arr)-1][1]) {
		return -1
	}

	l, r := 0, len(arr)

	for l < r {
		mid := (l + r) / 2
		if check(arr[mid][1]) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return arr[r][0]
}
