package temperature

import (
	"strconv"
)

/**
 * https://www.codingame.com/ide/puzzle/temperatures
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func FindClosestTo0(input []string) string {
	var lower int64 = 100
	var index int

	if len(input) == 0 {
		return "0"
	}

	for it, inputs := range input {
		t, _ := strconv.ParseInt(inputs, 10, 32)
		if Abs(t) < lower {
			lower = Abs(t)
			index = it
		} else if Abs(t) == lower && t > 0 {
			lower = Abs(t)
			index = it
		}
	}
	return input[index]
}
