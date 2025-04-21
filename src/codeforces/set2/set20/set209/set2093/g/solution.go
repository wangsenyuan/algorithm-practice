package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	buf.WriteTo(os.Stdout)
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
	n, k := readTwoNums(reader)
	a := readNNums(reader, n)
	return solve(a, k)
}

const H = 30

func solve(a []int, k int) int {
	// l...r 中，找到 a[i] ^ a[j] >= k
	res := -1

	update := func(arr []int, r int) {
		if len(arr) == 0 {
			return
		}
		l := arr[len(arr)-1]
		if res < 0 || r-l+1 < res {
			res = r - l + 1
		}
	}

	t := &trie{}

	for r, num := range a {
		t.add(num, H-1, r)
		tmp := t
		for i := H - 1; i >= 0 && tmp != nil; i-- {
			x := (num >> i) & 1
			y := (k >> i) & 1
			// y如果等于0， 那么就是希望找到1那个路径
			if y == 0 && tmp.children[1^x] != nil {
				update(tmp.children[1^x^y].arr, r)
			}

			tmp = tmp.children[x^y]
		}

		if tmp != nil {
			update(tmp.arr, r)
		}

	}

	return res
}

type trie struct {
	children [2]*trie
	arr      []int
}

func (t *trie) add(num int, d int, p int) {
	t.arr = append(t.arr, p)
	if d < 0 {
		return
	}
	x := (num >> d) & 1
	if t.children[x] == nil {
		t.children[x] = &trie{}
	}
	t.children[x].add(num, d-1, p)
}
