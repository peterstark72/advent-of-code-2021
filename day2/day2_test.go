package day2

import "testing"

const TEST_INPUT = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func Test1(t *testing.T) {

	if Run(TEST_INPUT) != 150 {
		t.Error()
	}

}

func Test2(t *testing.T) {

	if Run(TEST_INPUT) != 900 {
		t.Error()
	}

}
