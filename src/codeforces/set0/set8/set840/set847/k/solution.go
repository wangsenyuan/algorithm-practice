package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	nums := readNNums(reader, 5)
	n, a, b, k, f := nums[0], nums[1], nums[2], nums[3], nums[4]
	routes := make([]string, n)
	for i := range n {
		routes[i] = readString(reader)
	}
	return solve(n, a, b, k, f, routes)
}

func solve(n int, a int, b int, k int, f int, routes []string) int {
	city := make(map[string]int)

	addOrGet := func(name string) int {
		if v, ok := city[name]; ok {
			return v
		}
		city[name] = len(city)
		return len(city) - 1
	}

	cost := make([][]int, 2*n)
	for i := range cost {
		cost[i] = make([]int, 2*n)
	}
	var sum int
	prev := -1
	for _, route := range routes {
		s := strings.Split(route, " ")
		u := addOrGet(s[0])
		v := addOrGet(s[1])
		tmp := a
		if prev == u {
			tmp = b
		}
		cost[min(u, v)][max(u, v)] += tmp
		sum += tmp
		prev = v
	}

	var trips []int

	for i := range 2 * n {
		for j := i + 1; j < 2*n; j++ {
			if cost[i][j] > 0 {
				trips = append(trips, cost[i][j])
			}
		}
	}

	sort.Ints(trips)
	slices.Reverse(trips)

	for i := 0; i < len(trips) && i < k && trips[i] > f; i++ {
		sum -= trips[i]
		sum += f
	}

	return sum
}
