package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func unpack(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	var result strings.Builder
	result.Grow(len(s))

	runes := []rune(s)
	var lastChar rune
	isEscaped := false

	for i, r := range runes {
		if isEscaped {
			lastChar = r
			result.WriteRune(lastChar)
			isEscaped = false
			continue
		}

		if r == '\\' {
			isEscaped = true
			continue
		}

		if unicode.IsDigit(r) {
			// Цифра не может быть первым символом
			if i == 0 {
				return "", errors.New("invalid string: starts with a digit")
			}

			prevIsDigit := unicode.IsDigit(runes[i-1])
			prevWasEscaped := i > 1 && runes[i-2] == '\\'

			if prevIsDigit && !prevWasEscaped {
				return "", errors.New("invalid string: digit after digit")
			}

			// Нет символа для повторения
			if lastChar == 0 {
				return "", errors.New("invalid string: digit with no preceding character")
			}

			count := int(r - '0')
			if count > 0 {
				result.WriteString(strings.Repeat(string(lastChar), count-1))
			} else {
				currentStr := result.String()
				if len(currentStr) > 0 {
					result.Reset()
					result.WriteString(currentStr[:len(currentStr)-1])
				}
			}
			lastChar = 0
			continue
		}

		lastChar = r
		result.WriteRune(lastChar)
	}

	if isEscaped {
		return "", errors.New("invalid string: ends with an escape character")
	}

	return result.String(), nil
}

func main() {
	var str string
	if len(os.Args) < 2 {
		str = "45"
	} else if len(os.Args) == 2 {
		str = os.Args[1]
	}
	res, err := unpack(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
