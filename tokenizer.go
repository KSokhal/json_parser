package main

import (
	"fmt"
	"strconv"
	"unicode"
)

// JSON Spec: https://www.rfc-editor.org/rfc/rfc8259.txt
// Visual JSON flowchart: https://www.json.org/json-en.html

type TokenType int

const (
	TokenBeginArray TokenType = iota
	TokenBeginObject
	TokenEndArray
	TokenEndObject
	TokenNameSeparator  // Colon
	TokenValueSeparator // Comma
	TokenString
	TokenNumber
	TokenTrue
	TokenFalse
	TokenNull
)

type Token struct {
	Type  TokenType
	Value string
}

func tokenizer(data string) []Token {
	tokens := []Token{}
	index := 0

	for index < len(data) {
		char := rune(data[index])

		if char == '[' {
			tokens = append(tokens, Token{Type: TokenBeginArray})
			index++
			continue
		}
		if char == '{' {
			tokens = append(tokens, Token{Type: TokenBeginObject})
			index++
			continue
		}
		if char == ']' {
			tokens = append(tokens, Token{Type: TokenEndArray})
			index++
			continue
		}
		if char == '}' {
			tokens = append(tokens, Token{Type: TokenEndObject})
			index++
			continue
		}
		if char == ':' {
			tokens = append(tokens, Token{Type: TokenNameSeparator})
			index++
			continue
		}

		if char == ',' {
			tokens = append(tokens, Token{Type: TokenValueSeparator})
			index++
			continue
		}

		// If " is found it is start of string so need to collect the whole string
		if char == '"' {
			// Start the string
			value := ""
			// Get the first char of the string
			index++
			char = rune(data[index])
			// While the char is not a closing quote keep adding the chars to value
			for char != '"' {
				value += string(char)
				index++
				char = rune(data[index])
			}
			tokens = append(tokens, Token{Type: TokenString, Value: value})
			index++
			continue
		}

		if isAlphaNumeric(char) {
			value := ""
			for isAlphaNumeric(char) {
				value += string(char)
				index++
				char = rune(data[index])
			}
			tokenType, err := getTokenType(value)
			check(err)

			tokens = append(tokens, Token{Type: tokenType, Value: value})

			continue
		}

		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			index++

			continue

		}

		panic("Invalid character")

	}

	return tokens
}

// Checks if the given character is alphanumeric or underscore.
// It returns true if the character is a letter, digit, or underscore; otherwise, it returns false.
func isAlphaNumeric(char rune) bool {
	return unicode.IsLetter(char) || unicode.IsDigit(char) || char == '_'
}

func isNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

func isTrueValue(value string) bool {
	return value == "true"
}

func isFalseValue(value string) bool {
	return value == "false"
}

func isNullValue(value string) bool {
	return value == "null"
}

func getTokenType(value string) (TokenType, error) {
	if isNumber(value) {
		return TokenNumber, nil
	} else if isTrueValue(value) {
		return TokenTrue, nil
	} else if isFalseValue(value) {
		return TokenFalse, nil
	} else if isNullValue(value) {
		return TokenNull, nil
	} else {
		return TokenString, fmt.Errorf("unexpected value: %s", value)
	}
}
