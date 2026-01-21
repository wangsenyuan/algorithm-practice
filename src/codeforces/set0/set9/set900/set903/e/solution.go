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
	_, res := drive(reader)
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

func drive(reader *bufio.Reader) (words []string, res string) {
	nums := readNums(reader)
	k := nums[0]
	n := nums[1]
	words = make([]string, k)
	for i := range k {
		words[i] = readString(reader)
	}
	res = solve(n, words)
	return
}

func solve(n int, s []string) string {
	k := len(s)

	f0 := getFreq(s[0])

	distinct := true

	for i := range 26 {
		if f0[i] > 1 {
			distinct = false
			break
		}
	}

	for i := 1; i < k; i++ {
		f := getFreq(s[i])
		for j := range 26 {
			if f0[j] != f[j] {
				return "-1"
			}
		}
	}

	findDiff := func(i int, j int) []int {
		var res []int
		for p := range n {
			if s[i][p] != s[j][p] {
				res = append(res, p)
			}
		}
		return res
	}

	check := func(t string) bool {
		for i := range k {
			var cnt int
			for j := range n {
				if t[j] != s[i][j] {
					cnt++
				}
			}
			if cnt == 1 || cnt > 2 {
				return false
			}
			if cnt == 0 && distinct {
				return false
			}
		}
		return true
	}

	play := func(pos []int, buf []byte) string {
		if len(pos) >= 3 {
			for i := range len(pos) {
				for j := i + 1; j < len(pos); j++ {
					buf[pos[i]], buf[pos[j]] = buf[pos[j]], buf[pos[i]]
					if check(string(buf)) {
						return string(buf)
					}
					buf[pos[i]], buf[pos[j]] = buf[pos[j]], buf[pos[i]]
				}
			}
			return ""
		}
		// len(pos) <= 2, len(pos) == 1， 有没有可能？
		// s = aba   si = baa, sj = aab
		for _, i := range pos {
			for j := range n {
				buf[i], buf[j] = buf[j], buf[i]
				if check(string(buf)) {
					return string(buf)
				}
				buf[i], buf[j] = buf[j], buf[i]
			}
		}
		return ""
	}

	for i := range k {
		same := true
		for j := i + 1; j < k; j++ {
			// n * k <= 5000
			pos := findDiff(i, j)
			if len(pos) == 0 {
				continue
			}
			same = false
			if len(pos) > 4 {
				return "-1"
			}
			buf := []byte(s[i])
			res := play(pos, buf)
			if len(res) > 0 {
				return res
			}
		}
		if i == 0 && same {
			buf := []byte(s[0])
			buf[0], buf[1] = buf[1], buf[0]
			return string(buf)
		}
	}

	return "-1"
}

func getFreq(s string) []int {
	res := make([]int, 26)
	for _, v := range s {
		res[int(v-'a')]++
	}
	return res
}
