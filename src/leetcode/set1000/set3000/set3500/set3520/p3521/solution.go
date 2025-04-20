package p3521

func calculateScore(instructions []string, values []int) int64 {
	n := len(instructions)
	marked := make([]bool, n)
	var pos int
	var res int64
	for pos >= 0 && pos < n && !marked[pos] {
		marked[pos] = true
		if instructions[pos] == "add" {
			res += int64(values[pos])
			pos++
		} else {
			// jump
			pos += values[pos]
		}
	}
	return res
}
