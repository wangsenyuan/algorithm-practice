package p3612

import "slices"

func processStr(s string) string {
	var buf []byte

	for i := range len(s) {
		x := s[i]
		if x >= 'a' && x <= 'z' {
			buf = append(buf, x)
		} else if x == '*' {
			if len(buf) > 0 {
				buf = buf[:len(buf)-1]
			}
		} else if x == '#' {
			buf = append(buf, buf...)
		} else if x == '%' {
			slices.Reverse(buf)
		}
	}
	return string(buf)
}
