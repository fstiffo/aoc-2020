package solutions

import (
	"fmt"
	"strings"
)

type line struct {
	policyNum1   int
	policyNum2   int
	policyLetter byte
	password     string
}

func newLine(s string) (line, error) {
	var policyNum1 int
	var policyNum2 int
	var policyLetter byte
	var password string
	_, err := fmt.Sscanf(s, "%d-%d %c: %s", &policyNum1, &policyNum2, &policyLetter, &password)
	return line{
		policyNum1,
		policyNum2,
		policyLetter,
		password,
	}, err
}

func pwIsValidFirst(l line) bool {
	count := strings.Count(l.password, string(l.policyLetter))
	return l.policyNum1 <= count && count <= l.policyNum2
}

func pwIsValidSecond(l line) bool {
	return (l.password[l.policyNum1-1] == l.policyLetter) != (l.password[l.policyNum2-1] == l.policyLetter)
}

func getD02Input(filename string) (input []line) {
	lines := GetFileContent(filename)
	input = make([]line, len(lines))
	for i, s := range lines {
		input[i], _ = newLine(s)
	}
	return
}

// D02FirstHalf calculates the solution for the first half question of day 2
func D02FirstHalf() int {
	input := getD02Input("inputs/d02.txt")
	count := 0
	for _, l := range input {
		if pwIsValidFirst(l) {
			count++
		}
	}
	return count
}

// D02SecondHalf calculates the solution for the second half question of day 2
func D02SecondHalf() int {
	input := getD02Input("inputs/d02.txt")
	count := 0
	for _, l := range input {
		if pwIsValidSecond(l) {
			count++
		}
	}
	return count
}
