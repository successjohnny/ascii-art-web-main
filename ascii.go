package main

import (
	"errors"
	"os"
	"strings"
)

func GenerateAscii(text, banner string) (string, error) {
	text = strings.ReplaceAll(text, "\r\n", "\n")

	for _, r := range text {
		if r != '\n' && (r < 32 || r > 126) {
			return "", errors.New("invalid character")
		}
	}

	content, err := os.ReadFile("banners/" + banner + ".txt")
	if err != nil {
		return "", err
	}

	lines := strings.Split(
		strings.ReplaceAll(string(content), "\r\n", "\n"),
		"\n",
	)

	inputLines := strings.Split(text, "\n")
	var result strings.Builder

	allEmpty := true
	for _, line := range inputLines {
		if line != "" {
			allEmpty = false
			break
		}
	}

	if allEmpty && len(inputLines) > 1 {
		for i := 0; i < len(inputLines)-1; i++ {
			result.WriteByte('\n')
		}
		return result.String(), nil
	}

	for _, line := range inputLines {
		if line == "" {
			result.WriteByte('\n')
			continue
		}

		for row := 0; row < 8; row++ {
			for _, r := range line {
				offset := 1 + (int(r)-32)*9 + row
				result.WriteString(lines[offset])
			}
			result.WriteByte('\n')
		}
	}

	return result.String(), nil
}
