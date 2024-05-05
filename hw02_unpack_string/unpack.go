package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrInvalidString      = errors.New("invalid string")
	zeroRune         rune = 48
)

func Unpack(input string) (string, error) {
	var result strings.Builder

	rStr := []rune(input)

	for i := 0; i < len(rStr); {
		if unicode.IsDigit(rStr[i]) {
			return "", ErrInvalidString
		}

		nextIndex := i + 1
		tempCount := 1
		tempIncrement := 1

		if nextIndex < len(rStr) && unicode.IsDigit(rStr[nextIndex]) {
			tempCount = int(rStr[nextIndex] - zeroRune)
			tempIncrement = 2
		}

		s := stringByCount(string(rStr[i]), tempCount)
		result.WriteString(s)
		i += tempIncrement
	}

	return result.String(), nil
}

func stringByCount(ch string, count int) string {
	var res strings.Builder

	for i := 0; i < count; i++ {
		res.WriteString(ch)
	}

	return res.String()
}
