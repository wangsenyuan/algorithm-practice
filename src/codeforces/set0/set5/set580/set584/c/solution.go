package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	_, _, _, res := drive(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func drive(reader *bufio.Reader) (s1 string, s2 string, t int, res string) {
	_, t = readTwoNums(reader)
	s1 = readString(reader)
	s2 = readString(reader)
	res = solve(s1, s2, t)
	return
}

func solve(s1 string, s2 string, t int) string {
	var diff []int
	var same []int

	n := len(s1)
	for i := range n {
		if s1[i] != s2[i] {
			diff = append(diff, i)
		} else {
			same = append(same, i)
		}
	}
	m := len(diff)
	w := len(same)

	const abc = "abc"

	play := func(x int, z int) string {
		buf := []byte(s1)

		// 剩下的y个位置，必须和s1不同，但不一定和s2不同
		for i := range z {
			buf[same[i]] = 'a'
			if s1[same[i]] == 'a' {
				buf[same[i]] = 'b'
			}
		}
		y := t - x - z
		// 必须有y个和s2不同，
		for i := x; i < m; i++ {
			// 剩下的必须和s1不同
			if y > 0 {
				for j := range 3 {
					if s1[diff[i]] != abc[j] && s2[diff[i]] != abc[j] {
						buf[diff[i]] = abc[j]
						break
					}
				}
				y--
			} else {
				// 它和s2不同
				buf[diff[i]] = s2[diff[i]]
			}
		}

		return string(buf)
	}

	// abc vs def
	// 如果有x个和abc相同，那么就有 m - x个不同的位置，
	for x := 0; x <= m && x <= t; x++ {
		if m-x > t {
			continue
		}
		// 这x个肯定和def不同
		// 剩下y个位置 s3和s1不同
		y := m - x
		z := t - y
		if z <= w && z+x <= t && t-(z+x) <= y {
			return play(x, z)
		}
	}

	return "-1"
}
