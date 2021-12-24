package day3

import (
	"strconv"
	"strings"
)

func Run1(s string) int64 {

	rows := strings.Split(s, "\n")
	bits := len(rows[0])
	ones := make([]int, bits)
	zeros := make([]int, bits)
	for _, row := range rows {
		for p := 0; p < bits; p++ {
			if row[p] == '1' {
				ones[p] = ones[p] + 1
			} else {
				zeros[p] = zeros[p] + 1
			}
		}
	}
	var gamma string
	var epsilon string
	for p := 0; p < bits; p++ {
		if zeros[p] > ones[p] {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}
	}
	gamma_int, _ := strconv.ParseInt(gamma, 2, 64)
	epsilon_int, _ := strconv.ParseInt(epsilon, 2, 64)

	return gamma_int * epsilon_int
}
