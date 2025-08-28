package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	reader.ReadBytes('\n')
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	return solve(s)
}

func solve(s string) int {
	var packmen []int
	var asterisks []int
	for i := range s {
		if s[i] == 'P' {
			packmen = append(packmen, i)
		} else if s[i] == '*' {
			asterisks = append(asterisks, i)
		}
	}

	n := len(packmen)

	check := func(T int) bool {
		var j int
		for i := 0; i < n && j < len(asterisks); i++ {
			// packmen是否可以在T时间内，吃到asterisks[j]
			// 先左还是先右，其实不一样的，假设左边长度为x，右边为y
			// 如果先左后右，2 * x + y, 否则就是 2 * y + x
			if asterisks[j] < packmen[i] {
				// j在i的左边
				if asterisks[j] < packmen[i]-T {
					return false
				}
				x := packmen[i] - asterisks[j]
				for j < len(asterisks) && asterisks[j] < packmen[i] {
					j++
				}
				for j < len(asterisks) {
					y := asterisks[j] - packmen[i]
					if min(2*x+y, 2*y+x) > T {
						break
					}
					j++
				}
			} else {
				// ast[j] > packment[i]
				for j < len(asterisks) && asterisks[j] <= packmen[i]+T {
					j++
				}
			}
		}
		return j == len(asterisks)
	}

	return sort.Search(len(s)*2, check)
}
