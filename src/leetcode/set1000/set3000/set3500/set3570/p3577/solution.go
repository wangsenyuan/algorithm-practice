package p3577

const mod = 1e9 + 7

func mul(a, b int) int {
	return a * b % mod
}

func countPermutations(complexity []int) int {
	// 这个条目太烂了
	n := len(complexity)
	for i := 1; i < n; i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
	}
	res := 1
	for i := range n - 1 {
		res = mul(res, i+1)
	}
	return res
}
