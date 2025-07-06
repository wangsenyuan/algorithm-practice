package p3602

import (
	"strconv"
	"strings"
)

func concatHex36(n int) string {
	s1 := strconv.FormatInt(int64(n * n), 16)
	s2 := strconv.FormatInt(int64(n * n * n), 36)
	res := s1 + s2
	return strings.ToUpper(res)
}
