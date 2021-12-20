package day1

import (
	"strconv"
	"strings"
)

func Run(s string) (result int) {

	rows := strings.Split(s, "\n")
	prev, _ := strconv.Atoi(rows[0])

	for _, row := range rows[1:] {
		number, _ := strconv.Atoi(row)
		if number > prev {
			result = result + 1
		}
		prev = number
	}

	return
}
