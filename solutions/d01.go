package solutions

import "strconv"

func getInput(filename string) (input []int) {
	lines := GetFileContent(filename)
	input = make([]int, len(lines))
	for i, s := range lines {
		input[i], _ = strconv.Atoi(s)
	}
	return
}

// D01FirstHalf calculates the solution for the first half question of day 1
func D01FirstHalf() int {
	input := getInput("inputs/d01.txt")
	size := len(input)
	for first := 0; first < size; first++ {
		for second := first; second < size; second++ {
			if (input[first] + input[second]) == 2020 {
				return input[first] * input[second]
			}
		}

	}
	return -1
}

// D01SecondHalf calculates the solution for the second half question of day 1
func D01SecondHalf() int {
	input := getInput("inputs/d01.txt")
	size := len(input)
	for first := 0; first < size; first++ {
		for second := first + 1; second < size; second++ {
			for third := second + 1; third < size; third++ {
				if (input[first] + input[second] + input[third]) == 2020 {
					return input[first] * input[second] * input[third]
				}
			}
		}
	}
	return -1
}
