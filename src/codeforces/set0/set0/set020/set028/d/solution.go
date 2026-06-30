package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	res := drive(reader)
	fmt.Fprintln(writer, len(res))
	s := fmt.Sprintf("%v", res)
	fmt.Fprintln(writer, s[1:len(s)-1])
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	trucks := make([][]int, n)
	for i := range n {
		trucks[i] = make([]int, 4)
		fmt.Fscan(reader, &trucks[i][0], &trucks[i][1], &trucks[i][2], &trucks[i][3])
	}
	return solve(trucks)
}

func solve(trucks [][]int) []int {
	type Truck struct {
		id    int
		value int
		cnt   int
		left  int
	}
	type Node struct {
		truck int
		prev  int
	}

	var maxTotal int
	for _, cur := range trucks {
		c, l, r := cur[1], cur[2], cur[3]
		maxTotal = max(maxTotal, l+c+r)
	}

	groups := make([][]Truck, maxTotal+1)
	for i, cur := range trucks {
		v, c, l, r := cur[0], cur[1], cur[2], cur[3]
		total := l + c + r
		groups[total] = append(groups[total], Truck{i + 1, v, c, l})
	}

	dp := make([]int, maxTotal+1)
	fromNode := make([]int, maxTotal+1)
	for i := range dp {
		dp[i] = -1
		fromNode[i] = -1
	}

	bestValue := -1
	var bestPath []int

	for total, arr := range groups {
		if len(arr) == 0 {
			continue
		}

		touched := []int{0}
		nodes := []Node{}
		dp[0] = 0

		for _, cur := range arr {
			from := cur.left
			to := cur.left + cur.cnt
			if to > total {
				continue
			}
			if dp[from] >= 0 {
				val := dp[from]
				val += cur.value
				if val > dp[to] {
					if dp[to] < 0 {
						touched = append(touched, to)
					}
					dp[to] = val
					nodes = append(nodes, Node{cur.id, fromNode[from]})
					fromNode[to] = len(nodes) - 1
				}
			}
		}

		if dp[total] > bestValue {
			bestValue = dp[total]
			var path []int
			for id := fromNode[total]; id >= 0; id = nodes[id].prev {
				path = append(path, nodes[id].truck)
			}
			reverse(path)
			bestPath = path
		}

		for _, x := range touched {
			dp[x] = -1
			fromNode[x] = -1
		}
	}

	return bestPath
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
