package p3561

func resultingString(s string) string {
	buf := make([]byte, len(s))
	var top int
	for i := 0; i < len(s); i++ {
		if top > 0 && check(buf[top-1], s[i]) {
			top--
		} else {
			buf[top] = s[i]
			top++
		}
	}
	return string(buf[:top])
}

func check(a, b byte) bool {
	x := int(a - 'a')
	y := int(b - 'a')
	x, y = min(x, y), max(x, y)
	return x+1 == y || x+25 == y
}
