package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%.10f\n", drive(reader))
}

func drive(reader *bufio.Reader) float64 {
	var n, k int
	fmt.Fscan(reader, &n, &k)
	t := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &t[i])
	}
	return solve(k, t)
}

func solve(k int, t []int) float64 {
	n := len(t)
	// s 是 t 的前缀和
	// a[i] = Σ s[j]/t[j]，b[i] = Σ 1/t[j]
	// 区间 (p, i] 的期望是 a[i]-a[p] - s[p]*(b[i]-b[p])
	s := make([]int, n+1)
	a := make([]float64, n+1)
	b := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + t[i-1]
		a[i] = a[i-1] + float64(s[i])/float64(t[i-1])
		b[i] = b[i-1] + 1/float64(t[i-1])
	}

	// fp[i] 是前 i 关切成当前段数的最小期望。一段时就是 a[i]
	fp := make([]float64, n+1)
	copy(fp, a)

	type pt struct{ x, y float64 }
	for seg := 2; seg <= k; seg++ {
		// 点 (s[p], y[p])，y[p] = fp[p] - a[p] + s[p]*b[p]
		// fp[i] = a[i] + min (y[p] - b[i]*s[p])，维护下凸壳
		q := []pt{{0, 1e99}}
		for j := seg - 1; j <= n-(k-seg); j++ {
			for len(q) > 1 && q[0].y-b[j]*q[0].x >= q[1].y-b[j]*q[1].x {
				q = q[1:]
			}
			old := fp[j]
			fp[j] = q[0].y - b[j]*q[0].x + a[j]
			cur := pt{float64(s[j]), old - a[j] + float64(s[j])*b[j]}
			for len(q) > 1 {
				prev, last := q[len(q)-2], q[len(q)-1]
				x1, y1 := last.x-prev.x, last.y-prev.y
				x2, y2 := cur.x-last.x, cur.y-last.y
				if x1*y2-y1*x2 > 0 {
					break
				}
				q = q[:len(q)-1]
			}
			q = append(q, cur)
		}
	}
	return fp[n]
}
