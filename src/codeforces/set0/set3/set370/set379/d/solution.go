package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ok, s1, s2, _ := process(reader)
	if !ok {
		fmt.Println("Happy new year!")
	} else {
		fmt.Println(s1)
		fmt.Println(s2)
	}
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) (bool, string, string, []int) {
	nums := readNNums(reader, 4)

	ok, s1, s2 := solve(nums[0], nums[1], nums[2], nums[3])

	return ok, s1, s2, nums
}
func solve(k int, x int, n int, m int) (bool, string, string) {
	if x == 0 {
		return true, strings.Repeat("B", n), strings.Repeat("A", m)
	}

	// x > 0

	check := func(u int, du int, v int, dv int) bool {
		a := du / 2
		b := du % 2
		// CA
		if u*2+a+b > n {
			return false
		}
		a = dv / 2
		b = dv % 2
		if v*2+a+b > m {
			return false
		}

		s1 := u
		s2 := v

		for i := 3; i <= k; i++ {
			s3 := s1 + s2
			if du%2 == 1 && dv/2 == 1 {
				s3++
			}
			if s3 > x {
				return false
			}
			dw := du/2*2 + dv%2
			du, dv = dv, dw
			// s3 <= x
			s1, s2 = s2, s3
		}

		return s2 == x
	}

	for u := 0; u*2 <= n; u++ {
		for v := 0; v*2 <= m; v++ {
			for du := range 4 {
				for dv := range 4 {
					if check(u, du, v, dv) {
						return true, create(n, u, du), create(m, v, dv)
					}
				}
			}
		}
	}

	return false, "", ""
}

func create(n int, u int, flag int) string {
	a := flag / 2
	b := flag % 2
	buf := make([]byte, n)
	var i int
	if a == 1 {
		buf[0] = 'C'
		i++
	}

	for i+1 < n && u > 0 {
		buf[i] = 'A'
		buf[i+1] = 'C'
		i += 2
		u--
	}
	for i < n {
		buf[i] = 'B'
		i++
	}
	if b == 1 {
		buf[n-1] = 'A'
	}
	return string(buf)
}
