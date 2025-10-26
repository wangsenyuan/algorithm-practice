package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := drive(reader)
	fmt.Println(res)
}

func drive(reader *bufio.Reader) int {
	var n int
	fmt.Fscan(reader, &n)
	f := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &f[i])
	}
	return solve(n, f)
}

func solve(n int, f []int) int {
	deg := make([]int, n)
	for i := 0; i < n; i++ {
		if f[i]-1 != i {
			deg[f[i]-1]++
		}
	}

	dist := make([]int, n)
	que := make([]int, n)
	var head, tail int
	for i := 0; i < n; i++ {
		if deg[i] == 0 {
			que[head] = i
			head++
		}
	}

	for tail < head {
		u := que[tail]
		tail++
		v := f[u] - 1
		if u != v {
			dist[v] = max(dist[v], dist[u]+1)
			deg[v]--
			if deg[v] == 0 {
				que[head] = v
				head++
			}
		}
	}
	ans2 := 1
	var ans1 int
	for u := range n {
		ans1 = max(ans1, dist[u])
		if deg[u] > 0 {
			// 这个在换上
			v := u
			var cnt int
			for deg[v] > 0 {
				cnt++
				deg[v] = 0
				v = f[v] - 1
			}
			ans2 = lcm(ans2, cnt)
		}
	}

	if ans1 == 0 {
		return ans2
	}

	return max(1, (ans1+ans2-1)/ans2*ans2)
}

func lcm(a, b int) int {
	g := gcd(a, b)
	return a / g * b
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
