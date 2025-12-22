package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		grid[i] = readString(reader)[:m]
	}
	res := solve(grid)
	if res < 0 {
		fmt.Println("NO SOLUTION")
	} else {
		fmt.Println(res)
	}
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			return s[:i]
		}
	}
	return s
}

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

var dir [][]int = [][]int{
	{-1, -1},
	{-1, 1},
	{1, -1},
	{1, 1},
}

var cd [][]int = [][]int{
	{-1, -1}, {-1, 0}, {0, -1}, {0, 0},
}

var ab []int = []int{2, 1, 1, 2}

const inf = 1 << 60

func solve(grid []string) int {
	que := list.New()

	n := len(grid)
	m := len(grid[0])
	dist := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dist[i] = make([]int, m+2)
		for j := 0; j <= m+1; j++ {
			dist[i][j] = inf
		}
	}

	dist[1][1] = 0
	que.PushBack(Lamp{1, 1})

	getValue := func(x int, y int) int {
		if x <= 0 || x > n || y <= 0 || y > m {
			return 0
		}
		if grid[x-1][y-1] == '/' {
			return 1
		}
		return 2
	}

	for que.Len() > 0 {
		front := que.Front()
		u := front.Value.(Lamp)
		que.Remove(front)
		for i := 0; i <= 3; i++ {
			x, y := u.x+dir[i][0], u.y+dir[i][1]
			var d int
			if getValue(u.x+cd[i][0], u.y+cd[i][1]) != ab[i] {
				d = 1
			}
			if x > 0 && x <= n+1 && y > 0 && y <= m+1 && dist[x][y] > dist[u.x][u.y]+d {
				dist[x][y] = dist[u.x][u.y] + d
				if d == 0 {
					que.PushFront(Lamp{x, y})
				} else {
					que.PushBack(Lamp{x, y})
				}
			}
		}
	}
	if dist[n+1][m+1] >= inf {
		return -1
	}
	return dist[n+1][m+1]
}

type Lamp struct {
	x int
	y int
}
