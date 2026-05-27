package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_007

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(drive(reader))
}

func drive(reader *bufio.Reader) int {
	var n, h, w int
	fmt.Fscan(reader, &n, &h, &w)
	var pattern string
	fmt.Fscan(reader, &pattern)
	return solve(n, h, w, pattern)
}

type state struct {
	offsetTop    int
	offsetRight  int
	offsetBottom int
	offsetLeft   int
}

func (s state) area() int {
	dx := max(0, s.offsetRight-s.offsetLeft+1)
	dy := max(0, s.offsetTop-s.offsetBottom+1)
	return dx * dy
}

func solve(n int, h int, w int, pattern string) int {
	cur := []int{0, 0}
	lo := []int{0, 0}
	hi := []int{0, 0}
	dim := []int{w, h}

	var ans int
	play := func(moves int, cheating bool) bool {
		step := pattern[moves%n]
		switch step {
		case 'L':
			if cheating {
				cur[0] = hi[0] + 1
			} else {
				cur[0]++
			}
		case 'R':
			if cheating {
				cur[0] = lo[0] - 1
			} else {
				cur[0]--
			}
		case 'U':
			if cheating {
				cur[1] = hi[1] + 1
			} else {
				cur[1]++
			}
		default:
			if cheating {
				cur[1] = lo[1] - 1
			} else {
				cur[1]--
			}
		}

		var changed bool
		for i := range 2 {
			if cur[i] < lo[i] || cur[i] > hi[i] {
				ans = (ans + (moves+1)%mod*dim[i^1]) % mod
				dim[i]--
				changed = true
			}
			lo[i] = min(lo[i], cur[i])
			hi[i] = max(hi[i], cur[i])
		}

		return changed
	}

	inGame := func() bool {
		return dim[0] > 0 && dim[1] > 0
	}

	for moves := 0; inGame() && moves < n; moves++ {
		if moves > 0 && moves%n == 0 && cur[0] == 0 && cur[1] == 0 {
			return -1
		}
		play(moves, false)
	}

	var todo []int

	for moves := n; inGame() && moves < 2*n; moves++ {
		if moves%n == 0 && cur[0] == 0 && cur[1] == 0 {
			return -1
		}
		if play(moves, false) {
			todo = append(todo, moves%n)
		}
	}

	for k := 2; inGame(); k++ {
		for _, moves := range todo {
			if inGame() {
				if moves%n == 0 && cur[0] == 0 && cur[1] == 0 {
					return -1
				}
				play(moves+k*n, true)
			}
		}
	}

	return ans
}
