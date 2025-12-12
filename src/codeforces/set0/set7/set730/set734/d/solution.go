package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func drive(reader *bufio.Reader) bool {
	var n, x0, y0 int
	fmt.Fscan(reader, &n, &x0, &y0)
	pieces := make([][]int, n)
	for i := range n {
		pieces[i] = make([]int, 3)
		var c string
		var x, y int
		fmt.Fscan(reader, &c, &x, &y)
		switch c {
		case "B":
			pieces[i][0] = 0
		case "R":
			pieces[i][0] = 1
		default:
			pieces[i][0] = 2
		}
		pieces[i][1] = x
		pieces[i][2] = y
	}
	return solve(x0, y0, pieces)
}

type piece struct {
	w int
	x int
	y int
}

func (this piece) distanceFrom(x int, y int) int {
	return abs(this.x-x) + abs(this.y-y)
}
func abs(num int) int {
	return max(num, -num)
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}

func solve(x0 int, y0 int, pieces [][]int) bool {
	n := len(pieces)
	arr := make([]piece, n)
	for i := range n {
		arr[i] = piece{pieces[i][0], pieces[i][1], pieces[i][2]}
	}
	// 根据距离排序
	slices.SortFunc(arr, func(a, b piece) int {
		d1 := a.distanceFrom(x0, y0)
		d2 := b.distanceFrom(x0, y0)
		return d1 - d2
	})

	p1 := []int{-1, 0, 1}
	p2 := []int{-1, 0, 1}

	decode := func(dx int, dy int) int {
		for i, u := range p1 {
			for j, v := range p2 {
				if u == dx && v == dy {
					return i*3 + j
				}
			}
		}
		return -1
	}

	// 8个方向上，是否有棋子
	var flag int
	for _, cur := range arr {
		dx := x0 - cur.x
		dy := y0 - cur.y
		id := decode(sign(dx), sign(dy))

		switch cur.w {
		case 0:
			// BISHOP
			if abs(dx) == abs(dy) && (flag>>id)&1 == 0 {
				// 可以攻击到
				return true
			}
		case 1:
			// ROCK
			if (dx == 0 || dy == 0) && (flag>>id)&1 == 0 {
				return true
			}
		default:
			// cur.w = 2, a queue
			if (abs(dx) == abs(dy) || dx == 0 || dy == 0) && (flag>>id)&1 == 0 {
				return true
			}
		}
		if abs(dx) == abs(dy) || dx == 0 || dy == 0 {
			flag |= 1 << id
		}
	}

	return false
}
