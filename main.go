package roman_go

import (
	"fmt"
	s "strings"
)

var rs = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func Decode(r string) int {
	var t int
	var a []int
	for _, l := range r {
		lc := fmt.Sprintf("%c", l)
		t += rs[lc]
		if len(a) > 0 {
			if a[len(a)-1] < rs[lc] {
				t -= a[len(a)-1] * 2
			}
		}
		a = append(a, rs[lc])
	}
	fmt.Println(t)
	return t
}

func Encode(i int) string {

	if i < 1 || i > 4000 {
		return "ROMAN_OUT_OF_RANGE"
	}

	m := i / 1000
	ms := ""
	if m > 0 {
		ms = s.Repeat("M", m)
	}

	d := i / 500
	ds := ""
	if i > 0 && d%2 == 1 {
		ds = s.Repeat("D", 1)
	}

	c := i / 100
	cs := ""
	if c > 0 && d%2 == 1 {
		cs = s.Repeat("C", c%10-5)
	} else if c > 0 && d%2 == 0 {
		cs = s.Repeat("C", c%10)
	}

	l := i % 100
	ls := ""
	if l > 50 {
		ls = s.Repeat("L", 1)
	}

	x := i % 50
	xs := ""
	if x > 10 {
		xs = s.Repeat("X", x/10)
	}

	v := i % 10
	vs := ""
	if v > 5 {
		vs = s.Repeat("V", 1)
	}

	o := i % 5
	os := ""
	if o > 0 {
		os = s.Repeat("I", o)
	}

	rm := ms + ds + cs + ls + xs + vs + os

	if len(rm) == 4 && rm[len(rm)-4:] == "IIII" {
		return rm[:len(rm)-4] + "IV"
	}

	if len(rm) == 5 && rm[len(rm)-5:] == "VIIII" {
		return rm[:len(rm)-5] + "IX"
	}

	if len(rm) > 5 && rm[len(rm)-5:] == "XIIII" {
		return rm[:len(rm)-5] + "XIV"
	}

	if len(rm) > 5 && rm[len(rm)-5:] == "VIIII" {
		return rm[:len(rm)-5] + "IX"
	}

	return rm
}
