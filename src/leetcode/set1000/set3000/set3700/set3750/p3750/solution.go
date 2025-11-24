package p3750

func minimumFlips(n int) int {
	var ds []int
	for i := n; i > 0; i >>= 1 {
		ds = append(ds, i&1)
	}
	var res int
	for l, r := 0, len(ds)-1; l < r; l, r = l+1, r-1 {
		res += ds[l] ^ ds[r]
	}
	return res * 2
}
