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

// solveOne counts occurrences of pattern s in the k-th Fibonacci string, mod 1e9+7.
//
// Recurrence:  cnt[n] = cnt[n-1] + cnt[n-2] + cross[n]
//
// cross[n] = occurrences of s that straddle the join in f_{n-1}·f_{n-2}.
// A straddling occurrence can only touch the last (|s|-1) chars of f_{n-1}
// and the first (|s|-1) chars of f_{n-2}, so we only track those windows.
//
// Once the Fibonacci strings are long enough, pref[n] becomes constant and
// suff[n] becomes period-2.  Therefore cross[n] is also period-2, and the
// recurrence can be advanced to large k with 3×3 matrix exponentiation.
func solveOne(k int64, s string) int {
	pi := kmp(s)
	keep := len(s) - 1 // width of the boundary window on each side

	// length[n], pref[n], suff[n], cross[n], cnt[n] indexed from n=0.
	length := []int64{0, 1, 1}
	pref := []string{"", "a", "b"}
	suff := []string{"", "a", "b"}
	cross := []int{0, 0, 0}
	cnt := []int{0, countOcc(pi, s, "a"), countOcc(pi, s, "b")}

	buildNext := func() {
		n := len(length) // index being appended

		length = append(length, satAdd(length[n-1], length[n-2]))

		// pref[n] = first `keep` chars of f_n = f_{n-1}·f_{n-2}.
		// Once |f_{n-1}| >= keep, pref[n] == pref[n-1] (constant).
		curPref := pref[n-1]
		if len(curPref) < keep {
			extra := pref[n-2]
			if need := keep - len(curPref); len(extra) > need {
				extra = extra[:need]
			}
			curPref += extra
		}
		if len(curPref) > keep {
			curPref = curPref[:keep]
		}
		pref = append(pref, curPref)

		// suff[n] = last `keep` chars of f_n = f_{n-1}·f_{n-2}.
		// Once |f_{n-2}| >= keep, suff[n] == suff[n-2] (period-2).
		curSuff := suff[n-1] + suff[n-2]
		if len(curSuff) > keep {
			curSuff = curSuff[len(curSuff)-keep:]
		}
		suff = append(suff, curSuff)

		// cross[n]: run KMP over the join window (length ≤ 2*keep).
		curCross := 0
		if keep > 0 {
			curCross = countOcc(pi, s, suff[n-1]+pref[n-2])
		}
		cross = append(cross, curCross)
		cnt = append(cnt, add(add(cnt[n-1], cnt[n-2]), curCross))
	}

	// Phase 1: advance until length[stableAt] >= keep.
	// At that point pref has reached its full width and suff will become
	// period-2 two steps later.
	stableAt := 1
	if keep > 0 {
		for length[stableAt] < int64(keep) {
			buildNext()
			stableAt++
		}
	}

	// Phase 2: build 4 more steps so that cross[stableAt+3] and
	// cross[stableAt+4] are the two stable period-2 values.
	base := stableAt + 4
	for len(length)-1 < base {
		buildNext()
	}

	if k <= int64(len(cnt)-1) {
		return cnt[int(k)]
	}

	// Phase 3: matrix exponentiation.
	// cross[n] == stableCross[n&1] for all n >= stableAt+3.
	var stableCross [2]int
	stableCross[(stableAt+3)&1] = cross[stableAt+3]
	stableCross[(stableAt+4)&1] = cross[stableAt+4]

	// State vector v = [cnt[n], cnt[n-1], 1].
	// One step n → n+1: v = stepMat(stableCross[(n+1)&1]) · v
	// Two consecutive steps share the same parity pattern, giving a fixed
	// 2-step matrix we can raise to a power.
	v := [3]int{cnt[base], cnt[base-1], 1}
	cur := int64(base)

	if cur < k {
		twoStep := matMul(stepMat(stableCross[(cur+2)&1]), stepMat(stableCross[(cur+1)&1]))
		v = matVecMul(matPow(twoStep, (k-cur)/2), v)
		cur += ((k - cur) / 2) * 2
	}
	if cur < k {
		v = matVecMul(stepMat(stableCross[(cur+1)&1]), v)
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

// countOcc counts non-overlapping occurrences of pat (given its KMP table pi) in text.
func countOcc(pi []int, pat, text string) int {
	var cnt, j int
	for i := 0; i < len(text); i++ {
		for j > 0 && text[i] != pat[j] {
			j = pi[j-1]
		}
		if text[i] == pat[j] {
			j++
		}
		if j == len(pat) {
			cnt++
			j = pi[j-1]
		}
	}
	return cnt
}

// kmp builds the KMP failure (prefix) function for s.
func kmp(s string) []int {
	pi := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

type mat [3][3]int

// stepMat returns the transition matrix for one recurrence step with cross-count g:
//
//	[cnt[n+1]]   [1 1 g] [cnt[n]  ]
//	[cnt[n]  ] = [1 0 0] [cnt[n-1]]
//	[   1    ]   [0 0 1] [   1    ]
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
	res := mat{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
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
