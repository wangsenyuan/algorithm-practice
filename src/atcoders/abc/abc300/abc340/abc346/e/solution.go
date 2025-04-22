package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", cur[0], cur[1]))
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

func process(reader *bufio.Reader) [][]int {
	h, w, m := readThreeNums(reader)
	commands := make([][]int, m)
	for i := 0; i < m; i++ {
		commands[i] = readNNums(reader, 3)
	}
	return solve(h, w, commands)
}

func solve(h int, w int, commands [][]int) [][]int {
	// 后面对某一列的操作，会覆盖前面对同一列的操作
	// 需要知道某一列最后一次操作的时间，同时要知道这次操作后，有多少次不同的行操作
	row := make([]int, h)
	col := make([]int, w)
	cnt := make([]int, 2)
	for i := 0; i < h; i++ {
		row[i] = -1
	}
	for i := 0; i < w; i++ {
		col[i] = -1
	}

	res := make(map[int]int)
	// 将0作为结果先放上去
	res[0] = 0

	for i := len(commands) - 1; i >= 0; i-- {
		cur := commands[i]
		if cur[0] == 1 {
			// update row
			r, x := cur[1], cur[2]
			r--
			if row[r] > 0 {
				// 这行已经被操作过了
				continue
			}
			res[x] += w - cnt[1]
			row[r] = i
			cnt[0]++
		} else {
			c, x := cur[1], cur[2]
			c--
			if col[c] > 0 {
				continue
			}
			res[x] += h - cnt[0]
			col[c] = i
			cnt[1]++
		}
	}

	var arr [][]int

	var sum int
	for k, v := range res {

		arr = append(arr, []int{k, v})
		if k != 0 {
			sum += v
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	arr[0][1] = h*w - sum
	var p int
	for i := 0; i < len(arr); i++ {
		if arr[i][1] == 0 {
			continue
		}
		arr[p] = arr[i]
		p++
	}

	return arr[:p]
}
