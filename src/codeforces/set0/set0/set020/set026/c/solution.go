package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := readNNums(reader, 5)
	res := solve(nums[0], nums[1], nums[2], nums[3], nums[4])
	if len(res) == 0 {
		fmt.Println("IMPOSSIBLE")
	} else {
		var buf bytes.Buffer
		for _, s := range res {
			buf.WriteString(s)
			buf.WriteByte('\n')
		}
		fmt.Print(buf.String())
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

func solve(n int, m int, a int, b int, c int) []string {
	if (n&1)+(m&1) == 2 {
		return nil
	}

	buf := make([][]byte, n)
	for i := range n {
		buf[i] = make([]byte, m)
	}

	horizontal := "abu"
	vertical := "cdv"
	square := "efg"

	var row int
	if n&1 == 1 {
		// 把最后一行填满先
		if m/2 > a {
			return nil
		}
		a -= m / 2
		for j := 0; j < m; j += 2 {
			buf[row][j] = horizontal[(j/2)&1]
			buf[row][j+1] = horizontal[(j/2)&1]
		}
		row++
	}
	var col int
	if m&1 == 1 {
		if n/2 > b {
			return nil
		}
		b -= n / 2
		for i := 0; i < n; i += 2 {
			buf[i][col] = vertical[(i/2)&1]
			buf[i+1][col] = vertical[(i/2)&1]
		}
		col++
	}
	// 尽量使用c
	for i := row; i < n; i += 2 {
		for j := col; j < m; j += 2 {
			if c > 0 {
				var x int
				if i > 0 && buf[i-1][j] == square[x] || j > 0 && buf[i][j-1] == square[x] {
					x++
				}
				if i > 0 && buf[i-1][j] == square[x] || j > 0 && buf[i][j-1] == square[x] {
					x++
				}
				buf[i][j] = square[x]
				buf[i][j+1] = square[x]
				buf[i+1][j] = square[x]
				buf[i+1][j+1] = square[x]
				c--
				continue
			}
			if a > 1 {
				var x int
				if i > 0 && buf[i-1][j] == horizontal[x] || j > 0 && buf[i][j-1] == horizontal[x] {
					x++
				}
				if i > 0 && buf[i-1][j] == horizontal[x] || j > 0 && buf[i][j-1] == horizontal[x] {
					x++
				}
				buf[i][j] = horizontal[x]
				buf[i][j+1] = horizontal[x]
				buf[i+1][j] = horizontal[(x+1)%3]
				buf[i+1][j+1] = horizontal[(x+1)%3]
				a -= 2
				continue
			}
			if b > 1 {
				var x int
				if i > 0 && buf[i-1][j] == vertical[x] || j > 0 && buf[i][j-1] == vertical[x] {
					x++
				}
				if i > 0 && buf[i-1][j] == vertical[x] || j > 0 && buf[i][j-1] == vertical[x] {
					x++
				}
				buf[i][j] = vertical[x]
				buf[i][j+1] = vertical[(x+1)%3]
				buf[i+1][j] = vertical[x]
				buf[i+1][j+1] = vertical[(x+1)%3]
				b -= 2
				continue
			}
			return nil
		}
	}
	ans := make([]string, len(buf))
	for i := range len(buf) {
		ans[i] = string(buf[i])
	}
	return ans
}
