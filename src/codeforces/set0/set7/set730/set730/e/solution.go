package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	a := make([]int, n)
	d := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &a[i], &d[i])
	}
	return solve(a, d)
}

type team struct {
	id    int
	score int
}

func solve(a []int, d []int) int {
	// 先算出当前的名次
	n := len(a)
	arr := make([]team, n)
	for i := range n {
		arr[i] = team{id: i, score: a[i]}
	}

	slices.SortFunc(arr, func(x team, y team) int {
		return cmp.Or(y.score-x.score, x.id-y.id)
	})

	var ord []int
	pos := make([]int, n)
	for i, cur := range arr {
		id := cur.id
		pos[id] = i
		if d[id] <= 0 {
			ord = append(ord, id)
		}
	}

	for i := n - 1; i >= 0; i-- {
		id := arr[i].id
		if d[id] > 0 {
			ord = append(ord, id)
		}
	}

	var res int

	for _, id := range ord {
		// 现在解冻i的分数
		score := a[id] + d[id]
		x := pos[id]
		y := x
		if d[id] >= 0 {
			// 需要往前运动
			for y > 0 && (arr[y-1].score < score || arr[y-1].score == score && arr[y-1].id >= id) {
				arr[y] = arr[y-1]
				pos[arr[y-1].id] = y
				y--
			}
		} else {
			// 需要往后运动
			for y+1 < n && (score < arr[y+1].score || score == arr[y+1].score && id >= arr[y+1].id) {
				arr[y] = arr[y+1]
				pos[arr[y+1].id] = y
				y++
			}
		}

		res += abs(x - y)
		pos[id] = y
		arr[y] = team{id: id, score: score}
	}

	return res
}

func abs(num int) int {
	return max(num, -num)
}
