package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Printf("%.10f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n, T := readTwoNums(reader)
	songs := make([][]int, n)
	for i := range n {
		songs[i] = readNNums(reader, 2)
	}
	return solve(T, songs)
}

func solve(T int, songs [][]int) float64 {
	// n := len(songs)
	f := make([]float64, T+1)
	f[0] = 1
	nf := make([]float64, T+1)
	var ans float64
	for i, cur := range songs {
		i++
		p := float64(cur[0]) / 100
		t := cur[1]
		np := math.Pow(1-p, float64(t-1))
		np2 := (1 - p) * np
		// nf := make([]float64, T+1)
		clear(nf)
		for j := i; j <= T; j++ {
			var d float64
			if j > t {
				d = f[j-t-1] * np
			}
			nf[j] = (nf[j-1]-d)*(1-p) + f[j-1]*p
			if j >= t {
				nf[j] += f[j-t] * np2
			}
			ans += nf[j]
		}
		copy(f, nf)
	}

	return ans
}
