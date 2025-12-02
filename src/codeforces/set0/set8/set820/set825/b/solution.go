package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func readString(reader *bufio.Reader) string {
	bs, _ := reader.ReadBytes('\n')
	return strings.TrimSpace(string(bs))
}

func drive(reader *bufio.Reader) bool {
	board := make([]string, 10)
	for i := range 10 {
		board[i] = readString(reader)
	}
	return solve(board)
}

func solve(board []string) bool {

	check := func(x int, y int, dx int, dy int) bool {
		var cnt int
		var hole int
		for i, j := x, y; i >= 0 && i < 10 && j >= 0 && j < 10; i, j = i+dx, j+dy {
			if board[i][j] == 'O' {
				return false
			}
			if board[i][j] == '.' {
				hole++
				if hole == 2 {
					return false
				}
			} else {
				// X
				cnt++
			}
			if cnt+hole == 5 {
				return true
			}
		}
		return false
	}

	for i := range 10 {
		for j := range 10 {
			if board[i][j] == 'X' {
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						}
						if check(i, j, dx, dy) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
