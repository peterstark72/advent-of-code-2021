package day1

import "testing"

const TEST_INPUT = `199
200
208
210
200
207
240
269
260
263`

func TestPartOne(t *testing.T) {

	if RunPartOne(TEST_INPUT) != 7 {
		t.Error()
	}

}

func TestPartTwo(t *testing.T) {

	if RunPartTwo(TEST_INPUT) != 5 {
		t.Error()
	}

}
