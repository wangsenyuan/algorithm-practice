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

func process(reader *bufio.Reader) []int {
	n, m, r := readThreeNums(reader)
	sets := make([][][]int, m)
	for i := range m {
		sets[i] = make([][]int, n)
		for j := range n {
			sets[i][j] = readNNums(reader, n)
		}
	}
	rounds := make([][]int, r)
	for i := range r {
		rounds[i] = readNNums(reader, 3)
	}
	return solve(n, m, r, sets, rounds)
}

const inf = 1 << 60

func solve(n int, m int, r int, sets [][][]int, rounds [][]int) []int {
	// k 超过n,就相当于没有限制
	// 这个时候，任意两点间，可以使用最小值
	// 但是 k < n 的时候，如果使用dp，似乎会严重超时
	fp := make([][]int, n)
	for i := range n {
		fp[i] = make([]int, n)
		for j := range n {
			fp[i][j] = inf
		}
	}
	for _, set := range sets {
		for k := range n {
			for i := range n {
				for j := range n {
					set[i][j] = min(set[i][j], set[i][k]+set[k][j])
				}
			}
		}
		for i := range n {
			for j := range n {
				if i != j {
					fp[i][j] = min(fp[i][j], set[i][j])
				}
			}
		}
	}

	arr := make([]int, n)
	next := make([]int, n)
	bfs := func(s int) [][]int {
		dist := make([][]int, n)
		for i := range n {
			dist[i] = make([]int, n)
			for j := range n {
				dist[i][j] = inf
			}
		}
		for i := range n {
			dist[i][0] = fp[s][i]
			arr[i] = fp[s][i]
		}
		for i := range n - 1 {
			for j := range n {
				next[j] = inf
				for k := range n {
					next[j] = min(next[j], arr[k]+fp[k][j])
				}
				dist[j][i+1] = min(dist[j][i], next[j])
			}
			for j := range n {
				arr[j] = min(arr[j], next[j])
			}
		}
		return dist
	}
	dist := make([][][]int, n)
	for i := range n {
		dist[i] = bfs(i)
	}

	ans := make([]int, r)
	for i, cur := range rounds {
		s, t, k := cur[0], cur[1], cur[2]
		s--
		t--
		k = min(k, n-1)
		ans[i] = dist[s][t][k]
	}
	return ans
}
