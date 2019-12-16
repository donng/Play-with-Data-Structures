package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	var words []string

	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err != nil || io.EOF == err {
			break
		}

		wordSlice := strings.Fields(line)
		for _, word := range wordSlice {
			if word = extractStr(strings.ToLower(word)); word != "" {
				words = append(words, word)
			}
		}
	}

	return words
}

func extractStr(str string) string {
	var res []rune
	for _, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			res = append(res, letter)
		}
	}
	return string(res)
}
