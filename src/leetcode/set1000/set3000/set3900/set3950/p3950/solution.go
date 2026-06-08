package p3950

func consecutiveSetBits(n int) bool {
	w := n & (n >> 1)
	if w == 0 {
		return false
	}
	// w > 0
	return w&(w-1) == 0
}
