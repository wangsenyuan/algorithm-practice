package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	similarIndex, res := drive(reader)
	if similarIndex > 0 {
		fmt.Println(similarIndex)
	}
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
	for i := range len(ss) {
		nums[i], _ = strconv.Atoi(ss[i])
	}
	return nums
}

func drive(reader *bufio.Reader) (similarIndex int, res string) {
	readString(reader)
	s := readString(reader)
	m := readNums(reader)[0]
	words := make([]string, m)
	for i := range m {
		words[i] = readString(reader)
	}
	return solve(s, words)
}

func solve(s string, words []string) (similarIndex int, res string) {
	ss := strings.Split(s, " ")
	n := len(ss)

	var perms [][]int
	var f func(flag int, cur []int)

	f = func(flag int, cur []int) {
		if len(cur) == n {
			perms = append(perms, slices.Clone(cur))
			return
		}
		for i := range n {
			if flag&(1<<i) == 0 {
				cur = append(cur, i)
				f(flag|(1<<i), cur)
				cur = cur[:len(cur)-1]
			}
		}
	}

	f(0, nil)

	check := func(words []string, arr []int) bool {
		var j int
		for i := 0; i < len(words) && j < n; i++ {
			if words[i] == ss[arr[j]] {
				j++
			}
		}
		return j == n
	}

	play := func(word string) int {
		xx := strings.Split(word, " ")
		xx = xx[1:]

		w := -1
		for _, p := range perms {
			if check(xx, p) {
				x := countInversions(p)
				if w == -1 || x < w {
					w = x
				}
			}
		}
		if w < 0 {
			return 0
		}

		return n*(n-1)/2 - w + 1
	}

	var p int
	for i, cur := range words {
		p1 := play(cur)

		if p1 > p {
			similarIndex = i + 1
			p = p1
		}
	}
	if p == 0 {
		return 0, "Brand new problem!"
	}
	return similarIndex, "[:" + strings.Repeat("|", p) + ":]"
}

func countInversions(arr []int) int {
	var res int
	for i, x := range arr {
		for j := range i {
			if arr[j] > x {
				res++
			}
		}
	}
	return res
}
