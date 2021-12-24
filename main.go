package main

import (
	"fmt"

	"github.com/peterstark72/advent-of-code-2021/day1"
	"github.com/peterstark72/advent-of-code-2021/day2"
	"github.com/peterstark72/advent-of-code-2021/day3"
)

func main() {

	fmt.Println(day1.RunPartOne(day1.INPUT))
	fmt.Println(day1.RunPartTwo(day1.INPUT))

	fmt.Println(day2.Run(day2.INPUT))

	fmt.Println(day3.Run1(day3.INPUT))

}
