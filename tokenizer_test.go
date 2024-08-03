package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTokenizerValid(t *testing.T) {
	data := dataFromFile("tests/test.json")

	tokens := tokenizer(data)

	expected :=
		[]Token{
			{Type: TokenBeginObject},
			{Type: TokenString, Value: "id"},
			{Type: TokenNameSeparator},
			{Type: TokenString, Value: "647ceaf3657eade56f8224eb"},
			{Type: TokenValueSeparator},
			{Type: TokenString, Value: "index"},
			{Type: TokenNameSeparator},
			{Type: TokenNumber, Value: "0"},
			{Type: TokenValueSeparator},
			{Type: TokenString, Value: "something"},
			{Type: TokenNameSeparator},
			{Type: TokenBeginArray},
			{Type: TokenEndArray},
			{Type: TokenValueSeparator},
			{Type: TokenString, Value: "boolean"},
			{Type: TokenNameSeparator},
			{Type: TokenTrue, Value: "true"},
			{Type: TokenValueSeparator},
			{Type: TokenString, Value: "nullValue"},
			{Type: TokenNameSeparator},
			{Type: TokenNull, Value: "null"},
			{Type: TokenEndObject},
		}

	if !reflect.DeepEqual(tokens, expected) {
		for i := 0; i < len(expected); i++ {
			if tokens[i] != expected[i] {
				fmt.Println(tokens[i], expected[i], i)
			}

		}
		t.Errorf("Expected and actual tokens do not match.\nExpected: %+v\nGot: %+v", expected, tokens)
	}

}
