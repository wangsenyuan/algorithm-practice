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
	var buf strings.Builder
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d\n", x))
	}
	fmt.Print(buf.String())
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func drive(reader *bufio.Reader) []int {
	var k int64
	var m int
	fmt.Fscan(reader, &k, &m)
	queries := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &queries[i])
	}
	return solve(k, queries)
}

const mod = 1000000007

func add(a, b int) int {
	a += b
	if a >= mod {
		a -= mod
	}
	return a
}

func solve(k int64, queries []string) []int {
	ans := make([]int, len(queries))
	mem := make(map[string]int)
	for i, s := range queries {
		if v, ok := mem[s]; ok {
			ans[i] = v
			continue
		}
		v := solveOne(k, s)
		mem[s] = v
		ans[i] = v
	}
	return ans
}

func solveOne(k int64, s string) int {
	pi := kmp(s)
	l := len(s)
	keep := max(0, l-1)

	length := []int64{0, 1, 1}
	pref := []string{"", "a", "b"}
	suff := []string{"", "a", "b"}
	cross := []int{0, 0, 0}
	cnt := []int{0, countOcc(pi, s, "a"), countOcc(pi, s, "b")}

	buildNext := func() {
		i := len(length)
		length = append(length, satAdd(length[i-1], length[i-2]))

		curPref := pref[i-1]
		if len(curPref) < keep {
			need := keep - len(curPref)
			part := pref[i-2]
			if len(part) > need {
				part = part[:need]
			}
			curPref += part
		}
		if len(curPref) > keep {
			curPref = curPref[:keep]
		}
		pref = append(pref, curPref)

		curSuff := suff[i-1] + suff[i-2]
		if len(curSuff) > keep {
			curSuff = curSuff[len(curSuff)-keep:]
		}
		suff = append(suff, curSuff)

		curCross := 0
		if l > 1 {
			curCross = countOcc(pi, s, suff[i-1]+pref[i-2])
		}
		cross = append(cross, curCross)
		cnt = append(cnt, add(add(cnt[i-1], cnt[i-2]), curCross))
	}

	p := 1
	if keep > 0 {
		for length[p] < int64(keep) {
			buildNext()
			p++
		}
	}

	need := p + 4
	for len(length)-1 < need {
		buildNext()
	}

	if k <= int64(len(cnt)-1) {
		return cnt[int(k)]
	}

	base := p + 4
	v := [3]int{cnt[base], cnt[base-1], 1}
	g := [2]int{}
	g[(p+3)&1] = cross[p+3]
	g[(p+4)&1] = cross[p+4]

	cur := int64(base)
	if cur < k {
		pair := matMul(stepMat(g[(cur+2)&1]), stepMat(g[(cur+1)&1]))
		v = matVecMul(matPow(pair, (k-cur)/2), v)
		cur += ((k - cur) / 2) * 2
	}
	if cur < k {
		v = matVecMul(stepMat(g[(cur+1)&1]), v)
	}

	return v[0]
}

func satAdd(a, b int64) int64 {
	const inf int64 = 1 << 60
	if a > inf-b {
		return inf
	}
	return a + b
}

func countOcc(pi []int, pat string, text string) int {
	var cnt int
	var j int
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pat[j] {
			j = pi[j-1]
		}
		if text[i] == pat[j] {
			j++
		}
		if j == len(pat) {
			cnt++
			if cnt >= mod {
				cnt -= mod
			}
			j = pi[j-1]
		}
	}
	return cnt
}

func kmp(s string) []int {
	res := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := res[i-1]
		for j > 0 && s[i] != s[j] {
			j = res[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		res[i] = j
	}
	return res
}

type mat [3][3]int

func stepMat(g int) mat {
	return mat{
		{1, 1, g},
		{1, 0, 0},
		{0, 0, 1},
	}
}

func matMul(a, b mat) mat {
	var c mat
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			if a[i][k] == 0 {
				continue
			}
			for j := 0; j < 3; j++ {
				c[i][j] = (c[i][j] + int(int64(a[i][k])*int64(b[k][j])%mod)) % mod
			}
		}
	}
	return c
}

func matPow(a mat, e int64) mat {
	res := mat{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}
	for e > 0 {
		if e&1 == 1 {
			res = matMul(a, res)
		}
		a = matMul(a, a)
		e >>= 1
	}
	return res
}

func matVecMul(a mat, v [3]int) [3]int {
	var res [3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			res[i] = (res[i] + int(int64(a[i][j])*int64(v[j])%mod)) % mod
		}
	}
	return res
}
