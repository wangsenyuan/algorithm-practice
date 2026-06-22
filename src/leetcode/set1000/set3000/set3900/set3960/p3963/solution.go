package p3963

func createGrid(m int, n int) []string {
	res := make([]string, m)
	buf := make([]byte, n)
	for i := range m {
		for j := range n {
			if i == 0 || j == n-1 {
				buf[j] = '.'
			} else {
				buf[j] = '#'
			}
		}
		res[i] = string(buf)
	}
	return res
}
