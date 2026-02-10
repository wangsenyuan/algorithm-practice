package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	vp := readNum(reader)
	vd := readNum(reader)
	t := readNum(reader)
	f := readNum(reader)
	c := readNum(reader)

	res := solve(vp, vd, t, f, c)
	fmt.Println(res)
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func solve(vp, vd, t, f, c int) int {
	if vd <= vp {
		return 0
	}

	pv := float64(vp)
	dv := float64(vd)
	pt := float64(t)
	pf := float64(f)
	pc := float64(c)

	d := pv * pt
	if d >= pc {
		return 0
	}

	const eps = 1e-9
	var ans int
	for d+eps < pc {
		catchDist := d * dv / (dv - pv)
		if catchDist+eps >= pc {
			break
		}
		ans++
		backTime := catchDist / dv
		d = catchDist + pv*(backTime+pf)
	}
	return ans
}
