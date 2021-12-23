package day1

import (
	"strconv"
	"strings"
)

func atioarray(s string) (a []int) {
	for _, row := range strings.Split(s, "\n") {
		num, _ := strconv.Atoi(row)
		a = append(a, num)
	}
	return
}

func cumsum(a []int) (cs int) {
	for _, v := range a {
		cs = cs + v
	}
	return
}

func RunPartOne(s string) (result int) {
	a := atioarray(s)
	prev := a[0]
	for _, itm := range a[1:] {
		if itm > prev {
			result = result + 1
		}
		prev = itm
	}

	return
}

func RunPartTwo(s string) (result int) {
	a := atioarray(s)
	prev := cumsum(a[0:3])
	for i := 1; i < len(a)-2; i++ {
		ws := cumsum(a[i : i+3])
		if ws > prev {
			result = result + 1
		}
		prev = ws
	}
	return
}
