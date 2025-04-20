package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, move := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", move[0], move[1]))
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res [][]int, n int, files [][]int) {
	n, m := readTwoNums(reader)
	files = make([][]int, m)
	arr := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		var k int
		pos := readInt(s, 0, &k) + 1
		files[i] = make([]int, k)
		for j := range k {
			pos = readInt(s, pos, &files[i][j]) + 1
		}
		arr[i] = make([]int, k)
		copy(arr[i], files[i])
	}

	// files 不能被修改
	res = solve(n, arr)
	return
}

func solve(n int, files [][]int) [][]int {
	// 先把cluster n空出来
	var res [][]int

	belong := make([]int, n+1)
	for i := range n + 1 {
		belong[i] = -1
	}
	m := len(files)
	first := make([]int, m)

	for i, file := range files {
		first[i] = n
		for _, c := range file {
			belong[c] = i
			first[i] = min(first[i], c)
		}
	}

	move := func(from int, to int) {
		// to must be free
		res = append(res, []int{from, to})
		id := belong[from]
		belong[to] = id
		belong[from] = -1
		for i, c := range files[id] {
			if c == from {
				files[id][i] = to
				break
			}
		}
		for _, c := range files[id] {
			first[id] = min(first[id], c)
		}
	}

	findFree := func() int {
		for i := 1; i <= n; i++ {
			if belong[i] == -1 {
				return i
			}
		}
		return -1
	}

	marked := make([]bool, m)
	pos := 1
	for pos < n {
		found := -1
		for i := range m {
			if !marked[i] && (found == -1 || first[i] < first[found]) {
				found = i
			}
		}
		if found == -1 {
			break
		}

		for y := pos; y < pos+len(files[found]); y++ {
			x := files[found][y-pos]
			if x != y {
				// x要迁移到y的位置
				if belong[y] != -1 {
					z := findFree()
					move(y, z)
				}
				// y is free now
				move(x, y)
			}
		}

		marked[found] = true
		pos += len(files[found])
	}

	return res
}
