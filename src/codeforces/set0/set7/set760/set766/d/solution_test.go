package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := process(reader)
	for i := 0; i < len(res); i++ {
		expect := readString(reader)
		if res[i] != expect {
			t.Fatalf("Sample expect %s, but got %v at %d", s, res, i)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, `3 3 4
hate love like
1 love like
2 love hate
1 hate like
love like
love hate
like hate
hate like
YES
YES
NO
1
2
2
2`)
}

func TestSample2(t *testing.T) {
	runSample(t, `8 6 5
hi welcome hello ihateyou goaway dog cat rat
1 hi welcome
1 ihateyou goaway
2 hello ihateyou
2 hi goaway
2 hi hello
1 hi hello
dog cat
dog hi
hi hello
ihateyou goaway
welcome ihateyou
YES
YES
YES
YES
NO
YES
3
3
1
1
2
`)
}

func TestSample3(t *testing.T) {
	runSample(t, `26 25 53
h g x d n m r c y q u l b s v j o a z e f w t k i p
2 t z
1 b z
1 f t
2 t m
2 t g
2 j m
2 j d
1 t v
1 s t
2 h t
2 c t
1 b w
2 p t
1 t x
1 t n
1 x y
2 i t
2 q t
1 l c
2 o t
1 t u
2 r t
1 a t
1 e w
1 k t
k y
f x
e r
i v
m z
i f
z i
o d
s q
e x
r v
w s
u k
w f
p b
o x
i p
d y
q p
d x
h g
p t
a m
i d
n w
q r
i o
k h
c u
w h
p g
j i
c h
z d
e w
x s
o c
g v
b d
f w
c o
g m
n w
w t
e y
a i
n w
k v
q u
z g
r k
d q
s g
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
YES
1
1
1
2
1
2
1
1
2
2
2
2
1
2
1
2
1
2
1
2
1
2
2
1
2
1
1
2
2
1
1
2
1
1
1
1
1
2
1
2
1
1
2
2
2
2
2
1
2
1
2
1
2`)
}
