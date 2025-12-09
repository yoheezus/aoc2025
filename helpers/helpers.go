package helpers

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(path string) []string {
	var content []string
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	return content
}

func ReadInputAll(path string) string {
	b, _ := os.ReadFile(path)

	return string(b)
}

func Sum(s []int) int {
	total := 0
	for _, x := range s {
		total += x
	}

	return total
}
