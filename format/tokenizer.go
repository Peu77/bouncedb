package format

import (
	"strconv"
)

// Basically all tokens, with their identifier
const (
	EndOfInput     = 0
	IntegerLiteral = 1
	FloatLiteral   = 1
	StringLiteral  = 2
	Identifier     = 3
	LeftBracket    = 4
	RightBracket   = 5
	Colon          = 6
)

// TokenMatch describes a matched token.
type TokenMatch struct {
	tokenType int
	raw       string
}

// IsNumeric Checking is a string is numeric.
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// isParenthesis checks if a string is a "{"
func isLeftBracket(char string) bool {
	return char == "{"
}

// isParenthesis checks if a string is a "}"
func isRightBracket(char string) bool {
	return char == "}"
}

// lex breaks an input into tokens
func tokenize(input string) []TokenMatch {
	var results []TokenMatch
	index := 0

	for index < len(input) {
		char := string([]rune(input)[index])

		// We have a number. That means, increase until the number ends.
		if IsNumeric(char) {
			raw := char

			for index+1 < len(input) {
				currentChar := string([]rune(input)[index+1])

				if IsNumeric(currentChar) {
					index++
					raw += currentChar
				} else {
					break
				}
			}

			results = append(results, TokenMatch{tokenType: IntegerLiteral, raw: raw})
		} else if isLeftBracket(char) {
			results = append(results, TokenMatch{tokenType: LeftBracket, raw: char})
		} else if isRightBracket(char) {
			results = append(results, TokenMatch{tokenType: RightBracket, raw: char})
		} else if char == ":" {
			results = append(results, TokenMatch{tokenType: Colon, raw: char})
		} else if char == "\"" {
			raw := string([]rune(input)[index+1])
			index++

			// All char is relevant until we have another "
			for index+1 < len(input) && string([]rune(input)[index+1]) != "\"" {
				currentChar := string([]rune(input)[index+1])
				raw += currentChar
				index++
			}

			index++

			results = append(results, TokenMatch{tokenType: StringLiteral, raw: raw})
		} else {

			// Identifier
			if !isLeftBracket(char) && !isRightBracket(char) && char != ":" && !IsNumeric(char) && char != " " && char != "\n" {
				raw := char
				advance := true

				for index+1 < len(input) {
					currentChar := string([]rune(input)[index+1])

					if isLeftBracket(currentChar) || isRightBracket(currentChar) || currentChar == ":" || currentChar == " " || char == "\n" {
						advance = false
						break
					}

					index++
					raw += currentChar
				}
				if advance {
					index++
				}
				results = append(results, TokenMatch{tokenType: Identifier, raw: raw})
			}
		}

		index++
	}

	return append(results, TokenMatch{tokenType: EndOfInput})
}
