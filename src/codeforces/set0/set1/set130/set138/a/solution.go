package main

import (
	"bufio"
	"fmt"
	"os"
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
	res := make([]int, len(ss))
	for i := range len(ss) {
		res[i], _ = strconv.Atoi(ss[i])
	}
	return res
}

func drive(reader *bufio.Reader) string {
	w := readNums(reader)
	n, k := w[0], w[1]
	poem := make([]string, n*4)
	for i := range poem {
		poem[i] = readString(reader)
	}
	return solve(poem, k)
}

func solve(poem []string, k int) string {
	n := len(poem)
	// n % 4 == 0

	checkRhythm := func(a string, b string) bool {
		// 后缀要相同
		var cnt int
		for i := 0; i < len(a) && i < len(b); i++ {
			x := a[len(a)-1-i]
			y := b[len(b)-1-i]
			if x != y {
				return false
			}
			if isVowel(x) {
				cnt++
			}
			if cnt == k {
				return true
			}
		}
		return false
	}

	pairs := [][4]int{
		{0, 1, 2, 3},
		{0, 2, 1, 3},
		{0, 3, 1, 2},
	}

	findSchema := func(quatrain []string) [4]int {
		for _, cur := range pairs {
			ok := checkRhythm(quatrain[cur[0]], quatrain[cur[1]]) && checkRhythm(quatrain[cur[2]], quatrain[cur[3]])

			if !ok {
				continue
			}
			if checkRhythm(quatrain[cur[0]], quatrain[cur[2]]) {
				return [4]int{-2, -2, -2, -2}
			}
			return cur
		}
		return [4]int{-1, -1, -1, -1}
	}

	ans := findSchema(poem[0:4])
	if ans[0] == -1 {
		return "NO"
	}

	for i := 4; i < n; i += 4 {
		tmp := findSchema(poem[i : i+4])

		if tmp[0] == -1 {
			return "NO"
		}

		if tmp == ans || tmp[0] == -2 {
			continue
		}
		if ans[0] == -2 {
			ans = tmp
			continue
		}

		return "NO"
	}

	if ans[0] == -2 {
		return "aaaa"
	}
	if pairs[0] == ans {
		return "aabb"
	}
	if pairs[1] == ans {
		return "abab"
	}
	if pairs[2] == ans {
		return "abba"
	}
	return "NO"
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
