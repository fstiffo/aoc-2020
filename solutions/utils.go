package solutions

import (
	"io/ioutil"
	"log"
	"strings"
)

// GetFileContent read a text file into a string slice and return it
func GetFileContent(filename string) (lines []string) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(string(content), "\n")
	return
}
