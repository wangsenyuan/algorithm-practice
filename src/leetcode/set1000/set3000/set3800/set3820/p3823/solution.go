package p3823

import "slices"

func reverseByType(s string) string {

	arr := make([][]rune, 2)

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			arr[0] = append(arr[0], v)
		} else {
			arr[1] = append(arr[1], v)
		}
	}

	slices.Reverse(arr[0])
	slices.Reverse(arr[1])

	var buf []rune

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			buf = append(buf, arr[0][0])
			arr[0] = arr[0][1:]
		} else {
			buf = append(buf, arr[1][0])
			arr[1] = arr[1][1:]
		}
	}
	return string(buf)
}
