package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for _, v := range drive(reader) {
		fmt.Println(v)
	}
}

func drive(reader *bufio.Reader) []int {
	var n int
	fmt.Fscan(reader, &n)
	clouds := make([][]int, n)
	for i := range n {
		clouds[i] = make([]int, 4)
		fmt.Fscan(reader, &clouds[i][0], &clouds[i][1], &clouds[i][2], &clouds[i][3])
	}
	return solve(clouds)
}

func solve(clouds [][]int) []int {
	const N = 2000

	var diff [N + 2][N + 2]int

	n := len(clouds)

	for i, cur := range clouds {
		i += n
		u, d, l, r := cur[0], cur[1], cur[2], cur[3]
		diff[u][l] += i
		diff[u][r+1] -= i
		diff[d+1][l] -= i
		diff[d+1][r+1] += i
	}

	ans := make([]int, n)
	var sum int
	for i := range N {
		for j := range N {
			diff[i+1][j+1] += diff[i+1][j]
			diff[i+1][j+1] += diff[i][j+1]
			diff[i+1][j+1] -= diff[i][j]
			if diff[i+1][j+1] == 0 {
				sum++
			}
			if diff[i+1][j+1] >= n && diff[i+1][j+1] < 2*n {
				id := diff[i+1][j+1] - n
				ans[id]++
			}
		}
	}

	for i := range n {
		ans[i] = sum + ans[i]
	}

	return ans
}
