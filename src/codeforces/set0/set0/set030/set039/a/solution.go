package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func process(reader *bufio.Reader) int {
	a := readNum(reader)
	s := readString(reader)
	return solve(a, s)
}

func solve(a int, s string) int {
	type summand struct {
		cof int
		pos int // 0 for a++, 1 for ++a
	}
	var arr []summand

	for i := 0; i < len(s); {
		sign := 1
		if i > 0 && s[i-1] == '-' {
			sign = -1
		}
		if s[i] == 'a' {
			// a++
			arr = append(arr, summand{sign, 0})
			// a++ +
			// 把后面的+也跳过
			i += 4
			continue
		}

		if s[i] == '+' {
			// ++a
			arr = append(arr, summand{sign, 1})
			i += 4
			continue
		}
		var c int
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			c = c*10 + int(s[i]-'0')
			i++
		}
		c *= sign
		// s[i] == '*'
		i++
		if s[i] == 'a' {
			// c * a++
			arr = append(arr, summand{c, 0})
		} else {
			arr = append(arr, summand{c, 1})
		}

		i += 4
	}

	get := func(a int, x summand) int {
		if x.pos == 0 {
			return a * x.cof
		}
		return (a + 1) * x.cof
	}

	var res int

	for len(arr) > 1 {
		slices.SortFunc(arr, func(x, y summand) int {
			s1 := get(a, x) + get(a+1, y)
			s2 := get(a+1, x) + get(a, y)
			return s2 - s1
		})
		first := arr[0]
		res += get(a, first)
		a++
		arr = arr[1:]
	}

	res += get(a, arr[0])

	return res
}
