package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	reader := bufio.NewReader(strings.NewReader(s))
	res := drive(reader)
	if !slices.Equal(res, expect) {
		t.Errorf("Sample expect %v, but got %v", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := `3 10 12
##########
#z1a1111b#
##########
2 3 ab 2 8
`
	expect := []int{2, 8}
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `10 3 5
###
#w#
#1#
#a#
#1#
#1#
#1#
#1#
#b#
###
3 2 abababababababab 6 2
`
	expect := []int{8, 2}
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `3 10 6
##########
#z1a1311b#
##########
2 3 ab 2 8
`
	expect := []int{2, 7}
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `20 20 244
####################
###########a96b449c#
###########5##6###1#
###########1##2###6#
###########7##1###7#
###########3##4###2#
###########1##1###5#
###########d47e679f#
####################
####################
####################
####################
####################
####################
####################
####################
####################
####################
####################
####################
4 15 bcfefcbebcfcbabebab 2 13
`
	expect := []int{2, 15}
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := `20 20 387
####################
####################
####################
####################
####################
########a7b32137c###
########7#8#####1###
########8#7#####8###
########8#1#####4###
########5#5#####8###
########d6e69515f###
########5#5#####7###
########2#5#####1###
########7#3#####8###
########6#8#####4###
########8#8#####4###
########g6h82878i###
####################
####################
####################
15 11 efifihedadebcfihgda 6 10
`
	expect := []int{16, 9}
	runSample(t, s, expect)
}


func TestSample6(t *testing.T) {
	s := `3 10 6
##########
#z1a1111b#
##########
2 3 ab 2 8
`
	expect := []int{2, 9}
	runSample(t, s, expect)
}