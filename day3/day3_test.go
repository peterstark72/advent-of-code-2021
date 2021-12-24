package day3

import "testing"

const TEST_INPUT = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func Test1(t *testing.T) {
	if Run1(TEST_INPUT) != 198 {
		t.Error()
	}
}

func Test2(t *testing.T) {
	if Run1(INPUT) != 1307354 {
		t.Error()
	}
}
