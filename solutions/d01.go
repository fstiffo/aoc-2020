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
	return len(input)
}

// D01SecondHalf calculates the solution for the second half question of day 1
func D01SecondHalf() int {
	return 0
}
