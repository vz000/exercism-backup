package cryptosquare

import (
	"unicode"
	"strings"
)

func Normalize(pt string) string {
	text := []rune(pt)
	var newText strings.Builder
	for _, letter := range text {
		if unicode.IsLetter(letter){
			newText.WriteRune(unicode.ToLower(letter))
		} else if unicode.IsNumber(letter) {
			newText.WriteRune(letter)
		}
	}
	return newText.String()
}

func FindSize(pt string) (int, int) {
	length := len([]rune(pt))
	c := 0
	r := 0
	for {
		if c >= r {
			if (c - r) <= 1 {
				if (c * r) >= length {
					break
				} else {
					if (c - r) >= 1 {
						r++
					} else {
						c++
						r = 1
					}
				}
			} else {
				r++
			}
		} else {
			c++
		}
	}
	return c, r
}

func CreateRectangle(pt string, c int, r int) []string {
	normalizedText := []rune(pt)
	length := len(normalizedText)
	var chunk strings.Builder
	rectangle := make([]string, r)
	a := 0
	for i := 0; i < r; i++ {
		b := a
		for j := a; j < c+b; j++ {
			if j >= length {
				chunk.WriteString(" ")
			} else {
				chunk.WriteRune(normalizedText[j])
				a++
			}
		}
		b = a
		rectangle[i] = chunk.String()
		chunk.Reset()
	}
	return rectangle
}

func Encoded(rectangle []string, c int) string {
	var encoded strings.Builder
	for i := 0; i < c; i++ {
		for _, row := range rectangle {
			runeRow := []rune(row)
			if unicode.IsSpace(runeRow[i]) {
				encoded.WriteString(" ")
			} else {
				encoded.WriteRune(runeRow[i])
			}
		}

		if i + 1 < c {
			encoded.WriteString(" ")
		}
	}
	return encoded.String()
}

func Encode(pt string) string {
	normalized := Normalize(pt)
	c, r := FindSize(normalized)
	rectangle := CreateRectangle(normalized, c, r)
	encoded := Encoded(rectangle, c)
	return encoded
}
