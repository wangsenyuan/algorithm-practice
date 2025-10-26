package p3723

import "bytes"

func maxSumOfSquares(num int, sum int) string {
	var buf bytes.Buffer
	for range num {
		x := min(sum, 9)
		buf.WriteByte('0' + byte(x))
		sum -= x
	}
	if sum > 0 {
		return ""
	}
	return buf.String()
}
