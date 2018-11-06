package FileOperation

import (
	"os"
	"bufio"
	"io"
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

		strArr := strings.Fields(line)
		for _, word := range strArr {
			words = append(words, strings.ToLower(word))
		}
	}

	return words
}
