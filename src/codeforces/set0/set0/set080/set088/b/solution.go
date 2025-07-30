package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func process(reader *bufio.Reader) int {
	var n, m, x int
	fmt.Fscan(reader, &n, &m, &x)
	keyboard := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &keyboard[i])
	}
	var k int
	fmt.Fscan(reader, &k)
	var word string
	fmt.Fscan(reader, &word)
	return solve(x, keyboard, word)
}

func solve(mx int, keyboard []string, word string) int {
	// 找到所有离A最近的S
	n := len(keyboard)
	m := len(keyboard[0])

	var shifts [][2]int
	for i := range n {
		for j := range m {
			if keyboard[i][j] == 'S' {
				shifts = append(shifts, [2]int{i, j})
			}
		}
	}

	marked := make([]bool, 26)

	vis := make([]bool, 26)
	for i := range n {
		for j := range m {
			if keyboard[i][j] != 'S' {
				x := int(keyboard[i][j] - 'a')
				vis[x] = true
				if marked[x] {
					continue
				}
				for _, cur := range shifts {
					dx := cur[0] - i
					dy := cur[1] - j
					if dx*dx+dy*dy <= mx*mx {
						marked[x] = true
						break
					}
				}
			}
		}
	}

	var res int

	for i := range word {
		if word[i] >= 'A' && word[i] <= 'Z' {
			x := int(word[i] - 'A')
			if !marked[x] {
				if !vis[x] || len(shifts) == 0 {
					return -1
				}
				res++
			}
		} else {
			x := int(word[i] - 'a')
			if !vis[x] {
				return -1
			}
		}
	}

	return res
}

func abs(x int) int {
	return max(x, -x)
}
