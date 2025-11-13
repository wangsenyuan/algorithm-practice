package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, best, swap := drive(reader)
	fmt.Println(best)
	fmt.Println(swap[0], swap[1])
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) (s string, t string, best int, swap []int) {
	readString(reader)
	s = readString(reader)
	t = readString(reader)
	best, swap = solve(s, t)
	return
}

func solve(s string, t string) (int, []int) {
	// 如果交换i, j, s[i] = t[j] & s[j] = t[i], 那么得到2
	// 如果交换i, j, s[i] = t[j] & s[j] != t[i] 且 s[j] != t[j], 则得到1
	// 否则没有必要交换

	var pos [26][26]int
	for i := range 26 {
		for j := range 26 {
			pos[i][j] = -1
		}
	}

	var sum int
	for i := range s {
		if s[i] != t[i] {
			sum++
			x := int(s[i] - 'a')
			y := int(t[i] - 'a')
			pos[x][y] = i
		}
	}

	if sum == 0 {
		return 0, []int{-1, -1}
	}

	for x := range 26 {
		for y := range 26 {
			if x != y && pos[x][y] >= 0 && pos[y][x] >= 0 {
				i := pos[x][y]
				j := pos[y][x]
				return sum - 2, []int{i + 1, j + 1}
			}
		}
	}

	// 尝试减少1
	for i := range s {
		if s[i] != t[i] {
			x := int(t[i] - 'a')
			for y := range 26 {
				if x != y && pos[x][y] >= 0 {
					// 从其他地方交换x过来
					j := pos[x][y]
					if s[j] != t[j] {
						return sum - 1, []int{i + 1, j + 1}
					}
				}
			}
		}
	}
	return sum, []int{-1, -1}
}
