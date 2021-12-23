package day2

import (
	"strconv"
	"strings"
)

type Cmd struct {
	Name  string
	Value int
}

func Run(s string) int {

	var h int
	var d int
	var aim int
	for _, cmd := range strings.Split(s, "\n") {

		cols := strings.Split(cmd, " ")
		name := cols[0]
		val, _ := strconv.Atoi(cols[1])
		switch name {
		case "forward":
			h = h + val
			d = d + aim*val
		case "up":
			aim = aim - val
		case "down":
			aim = aim + val
		}
	}
	return d * h
}
