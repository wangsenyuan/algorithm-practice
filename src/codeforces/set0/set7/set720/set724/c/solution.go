package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	var buf bytes.Buffer
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
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

func process(reader *bufio.Reader) []int {
	n, m, k := readThreeNums(reader)
	sensors := make([][]int, k)
	for i := 0; i < k; i++ {
		sensors[i] = readNNums(reader, 2)
	}
	return solve(n, m, sensors)
}

func solve(n int, m int, sensors [][]int) []int {
	k := len(sensors)
	ans := make([]int, k)

	dia := make([][]int, 2*(n+m))
	anti := make([][]int, 2*(n+m))

	offset := n + m

	for i := range k {
		ans[i] = -1
		x, y := sensors[i][0], sensors[i][1]
		// y = x + a
		a := y - x
		dia[a+offset] = append(dia[a+offset], i)
		// y = -x + a
		a = y + x
		anti[a] = append(anti[a], i)
	}

	row := make([][2][2]bool, n+1)
	col := make([][2][2]bool, m+1)

	update := func(i int, pos int, cur int) {
		if ans[i] != -1 {
			return
		}
		switch pos {
		case 0:
			ans[i] = cur + sensors[i][1]
		case 2:
			ans[i] = cur + m - sensors[i][1]
		case 3:
			ans[i] = cur + sensors[i][0]
		default:
			ans[i] = cur + n - sensors[i][0]
		}
	}

	var x, y, d int
	var cur int
	for {
		var pos int
		if y == 0 {
			pos = 0
		} else if x == n {
			pos = 1
		} else if y == m {
			pos = 2
		} else {
			pos = 3
		}
		// 在上下两层的时候
		if pos%2 == 0 {
			if row[x][pos/2][d] {
				break
			}
			row[x][pos/2][d] = true
		} else {
			if col[y][pos/2][d] {
				break
			}
			col[y][pos/2][d] = true
		}

		// 在逆时针时，且处于底部或者顶部
		// 或者在顺时针时，切处于左右两边
		if d == 0 && (pos == 0 || pos == 2) || d == 1 && (pos == 1 || pos == 3) {
			for _, i := range dia[y-x+offset] {
				update(i, pos, cur)
			}
		} else {
			// 在左右两边
			for _, i := range anti[x+y] {
				update(i, pos, cur)
			}
		}

		// 代码不好看且容易错, d = 0, 表示逆时针运动, d = 1, 表示顺时针运动
		// 现在清晰多了
		if d == 0 {
			switch pos {
			case 0:
				w := min(n-x, m)
				if w == m {
					// 到达顶部了
					d ^= 1
				}
				cur += w
				x += w
				y += w

			case 1:
				w := min(n, m-y)
				if w == n {
					d ^= 1
				}
				cur += w
				x -= w
				y += w

			case 2:
				w := min(x, m)
				if w == m {
					d ^= 1
				}
				cur += w
				x -= w
				y -= w

			default:
				w := min(n, y)
				if w == n {
					d ^= 1
				}
				cur += w
				x += w
				y -= w

			}
		} else {
			// 顺时针
			switch pos {
			case 0:
				w := min(x, m)
				if m == w {
					d ^= 1
				}
				cur += w
				x -= w
				y += w

			case 1:
				w := min(n, y)
				if n == w {
					d ^= 1
				}
				cur += w
				x -= w
				y -= w

			case 2:
				w := min(n-x, m)
				if m == w {
					d ^= 1
				}
				cur += w
				x += w
				y -= w

			default:
				w := min(n, m-y)
				if n == w {
					d ^= 1
				}

				cur += w
				x += w
				y += w
			}
		}

		if x+y == 0 || x == 0 && y == m || x == n && y == 0 || x == n && y == m {
			break
		}
	}

	return ans
}
