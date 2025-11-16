package p3746

func minLengthAfterRemovals(s string) int {
	n := len(s)
	var level int
	for i := range n {
		if s[i] == 'a' {
			level++
		} else {
			level--
		}
	}
	return max(level, -level)
}
