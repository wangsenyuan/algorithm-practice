package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	cards := make([][]int, n)
	for i := range n {
		cards[i] = readNNums(reader, 2)
	}
	return solve(cards)
}

type pair struct {
	first  int
	second int
}

func solve(cards [][]int) int {
	n := len(cards)
	freq := make(map[int]pair)

	addFront := func(x int) {
		if v, ok := freq[x]; !ok {
			freq[x] = pair{1, 0}
		} else {
			v.first++
			freq[x] = v
		}
	}

	addBack := func(x int) {
		if v, ok := freq[x]; !ok {
			freq[x] = pair{0, 1}
		} else {
			v.second++
			freq[x] = v
		}
	}

	// 这些不能反转
	special := make(map[int]int)

	for _, c := range cards {
		f, b := c[0], c[1]
		if f != b {
			addFront(f)
			addBack(b)
		} else {
			special[f]++
		}
	}
	h := (n + 1) / 2
	for _, v := range special {
		if v >= h {
			return 0
		}
	}

	res := n
	for x, v := range freq {
		u := special[x]
		a := u + v.first
		if a >= h {
			return 0
		}
		b := h - a
		if b <= v.second {
			// 需要反转的部分
			res = min(res, b)
		}
	}

	if res > h {
		return -1
	}

	return res
}
