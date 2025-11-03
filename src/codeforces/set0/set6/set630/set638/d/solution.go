package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) int {
	n, m, _ := readThreeNums(reader)
	layers := make([][]string, n)
	for i := 0; i < n; i++ {
		layers[i] = make([]string, m)
		for j := 0; j < m; j++ {
			layers[i][j] = readString(reader)
		}
		readString(reader)
	}
	return solve(layers)
}

func solve(layers [][]string) int {
	n := len(layers)
	m := len(layers[0])
	k := len(layers[0][0])

	var res int

	check := func(x, y, z int) bool {
		// 如果(x, y, z)的下方有一个1，如果它的上方是1，yes
		if x > 0 && layers[x-1][y][z] == '1' {
			if x+1 < n && layers[x+1][y][z] == '1' {
				return true
			}
			if y+1 < m && layers[x][y+1][z] == '1' && layers[x-1][y+1][z] == '0' {
				return true
			}
			if z+1 < k && layers[x][y][z+1] == '1' && layers[x-1][y][z+1] == '0' {
				return true
			}
		}
		if y > 0 && layers[x][y-1][z] == '1' {
			if y+1 < m && layers[x][y+1][z] == '1' {
				return true
			}

			if z+1 < k && layers[x][y][z+1] == '1' && layers[x][y-1][z+1] == '0' {
				return true
			}
			if x+1 < n && layers[x+1][y][z] == '1' && layers[x+1][y-1][z] == '0' {
				return true
			}
		}
		if z > 0 && layers[x][y][z-1] == '1' {
			if z+1 < k && layers[x][y][z+1] == '1' {
				return true
			}
			if y+1 < m && layers[x][y+1][z] == '1' && layers[x][y+1][z-1] == '0' {
				return true
			}
			if x+1 < n && layers[x+1][y][z] == '1' && layers[x+1][y][z-1] == '0' {
				return true
			}
		}
		return false
	}

	for i, layer := range layers {
		for j, row := range layer {
			for u := range k {
				if row[u] == '1' {
					if check(i, j, u) {
						res++
					}
				}
			}
		}
	}
	return res
}
