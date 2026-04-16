package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readNums(reader *bufio.Reader) []int {
	s := readString(reader)
	ss := strings.Split(s, " ")
	nums := make([]int, len(ss))
	for i, ss := range ss {
		nums[i], _ = strconv.Atoi(ss)
	}
	return nums
}

func drive(reader *bufio.Reader) int {
	k := readNums(reader)[1]
	s := readString(reader)
	return solve(k, s)
}

func solve(k int, s string) int {
	n := len(s)
	k -= 2
	// first and last mins must be used

	check := func(expect int) bool {
		prev := 0
		lastAvalible := 0
		x := k
		for i := 1; i < n-1; i++ {
			if s[i] == '0' {
				lastAvalible = i
			}
			if i-prev > expect {
				if n-1-i <= expect {
					return true
				}
				x--
				if x < 0 || i-lastAvalible > expect {
					return false
				}
				prev = lastAvalible
			}
		}
		return true
	}

	return sort.Search(n-1, check)
}
