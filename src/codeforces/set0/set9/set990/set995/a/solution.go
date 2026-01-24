package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, res := drive(reader)
	if len(res) == 0 {
		fmt.Println(-1)
		return
	}
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Fprintln(writer, len(res))
	for _, cur := range res {
		fmt.Fprintln(writer, cur[0], cur[1], cur[2])
	}
}

func drive(reader *bufio.Reader) (k int, park [][]int, res [][]int) {
	var n int
	fmt.Fscan(reader, &n, &k)
	park = make([][]int, 4)
	for i := range 4 {
		park[i] = make([]int, n)
		for j := range n {
			fmt.Fscan(reader, &park[i][j])
		}
	}
	res = solve(k, park)
	return
}

func solve(k int, park [][]int) [][]int {
	n := len(park[0])
	a := make([][]int, 4)
	for i := range 4 {
		a[i] = slices.Clone(park[i])
	}

	var ans [][]int
	var tot int

	move := func(id int, x1 int, y1 int, x2 int, y2 int) {
		ans = append(ans, []int{id, x2 + 1, y2 + 1})
		a[x1][y1] = 0
		a[x2][y2] = id
	}

	var cycle [][]int

	for i := range n {
		cycle = append(cycle, []int{1, i})
		if a[1][i] != 0 {
			if park[0][i] == a[1][i] {
				move(a[1][i], 1, i, 0, i)
				continue
			}
			tot++
		}
	}

	for i := n - 1; i >= 0; i-- {
		cycle = append(cycle, []int{2, i})
		if a[2][i] != 0 {
			if park[3][i] == a[2][i] {
				move(a[2][i], 2, i, 3, i)
				continue
			}
			tot++
		}
	}

	if tot == 2*n {
		return nil
	}

outerLoop:
	for tot > 0 {
		for i := range n {
			if a[1][i] != 0 && park[0][i] == a[1][i] {
				move(a[1][i], 1, i, 0, i)
				tot--
				continue outerLoop
			}
			if a[2][i] != 0 && park[3][i] == a[2][i] {
				move(a[2][i], 2, i, 3, i)
				tot--
				continue outerLoop
			}
		}

		for i := range cycle {
			r1, c1 := cycle[i][0], cycle[i][1]
			r2, c2 := cycle[(i+1)%len(cycle)][0], cycle[(i+1)%len(cycle)][1]
			if a[r1][c1] == 0 && a[r2][c2] != 0 {
				move(a[r2][c2], r2, c2, r1, c1)
				continue outerLoop
			}
		}
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
