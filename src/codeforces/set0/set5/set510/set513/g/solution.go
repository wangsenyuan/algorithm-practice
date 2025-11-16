package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Printf("%.12f\n", res)
}

func drive(reader *bufio.Reader) float64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	p := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &p[i])
	}
	return solve(k, p)
}

func solve(k int, p []int) float64 {
	k = min(k, 1000)

	n := len(p)
	f := make([][]float64, n)
	nf := make([][]float64, n)
	for i := range n {
		f[i] = make([]float64, n)
		nf[i] = make([]float64, n)
		for j := i + 1; j < n; j++ {
			if p[i] > p[j] {
				f[i][j] = 1
			}
		}
	}

	segs := n * (n + 1) / 2
	for range k {
		for i := range n {
			for j := i + 1; j < n; j++ {
				nf[i][j] = f[i][j] * float64(((i+1)*i/2)+((j-i)*(j-i-1)/2)+((n-j)*(n-j-1)/2))
				for sum := j; sum <= i+n-1; sum++ {
					from := max(0, sum-n+1)
					to := min(i, sum-j)
					nf[i][j] += (1.0 - f[sum-j][sum-i]) * float64(to-from+1)
				}
				for sum := i; sum <= i+j-1; sum++ {
					from := max(0, sum-j+1)
					to := min(i, sum-i)
					nf[i][j] += f[sum-i][j] * float64(to-from+1)
				}
				for sum := i + j + 1; sum <= j+(n-1); sum++ {
					from := max(i+1, sum-n+1)
					to := min(j, sum-j)
					nf[i][j] += f[i][sum-j] * float64(to-from+1)
				}
				nf[i][j] /= float64(segs)
			}
		}
		for i := range n {
			for j := i + 1; j < n; j++ {
				f[i][j] = nf[i][j]
			}
		}
	}
	var ans float64
	for i := range n {
		for j := i + 1; j < n; j++ {
			ans += f[i][j]
		}
	}
	return ans
}
