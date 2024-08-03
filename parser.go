package main

import (
	"fmt"
	"strconv"
)

type ASTNodeType string

const (
	NodeObject  = "object"
	NodeArray   = "array"
	NodeString  = "string"
	NodeNumber  = "number"
	NodeBoolean = "boolean"
	NodeNull    = "null"
)

// Define ASTNode as an interface in Go
type ASTNode interface {
}

// ObjectNode represents an object node in the AST
type ObjectNode struct {
	Type  ASTNodeType
	Value map[string]ASTNode
}

// ArrayNode represents an array node in the AST
type ArrayNode struct {
	Type  ASTNodeType
	Value []ASTNode
}

// StringNode represents a string node in the AST
type StringNode struct {
	Type  ASTNodeType
	Value string
}

// NumberNode represents a number node in the AST
type NumberNode struct {
	Type  ASTNodeType
	Value float64 // Go uses float64 for JSON numbers
}

// BooleanNode represents a boolean node in the AST
type BooleanNode struct {
	Type  ASTNodeType
	Value bool
}

// NullNode represents a null node in the AST
type NullNode struct {
	Type string
}

func parser(tokens []Token) ASTNode {
	if len(tokens) == 0 {
		fmt.Println("No tokens to parse")
		return nil
	}

	index := 0
	return parseValue(&tokens, &index)
}
func parseValue(tokens *[]Token, index *int) ASTNode {
	token := (*tokens)[*index]

	switch token.Type {
	case TokenString:
		return StringNode{Type: NodeString, Value: token.Value}
	case TokenNumber:
		value, err := strconv.ParseFloat(token.Value, 64)
		check(err)
		return NumberNode{Type: NodeNumber, Value: value}
	case TokenTrue:
		return BooleanNode{Type: NodeBoolean, Value: true}
	case TokenFalse:
		return BooleanNode{Type: NodeBoolean, Value: false}
	case TokenNull:
		return NullNode{}
	case TokenBeginObject:
		return parseObject(tokens, index)
	case TokenBeginArray:
		return parseArray(tokens, index)
	default:
		fmt.Println("Unexpected token type")
	}
	return nil
}

func parseObject(tokens *[]Token, index *int) ASTNode {
	node := ObjectNode{Type: NodeObject, Value: make(map[string]ASTNode)}

	// Skip open brace {
	*index++
	token := (*tokens)[*index]

	// While not the end of object
	for token.Type != TokenEndObject {
		if token.Type == TokenString {
			// Read the key
			key := token.Value
			*index++
			token = (*tokens)[*index]

			// Skip colon :
			if token.Type != TokenNameSeparator {
				panic("Expected :")
			}
			*index++
			token = (*tokens)[*index]

			value := parseValue(tokens, index)
			node.Value[key] = value

		} else {
			panic("Expected string key")
		}

		*index++
		token = (*tokens)[*index]
		// Skip comma,
		if token.Type == TokenValueSeparator {
			*index++
			token = (*tokens)[*index]
		}
	}

	return node
}

func parseArray(tokens *[]Token, index *int) ASTNode {
	node := ArrayNode{Type: NodeArray, Value: make([]ASTNode, 0)}
	// Skip open brace [
	*index++
	token := (*tokens)[*index]
	// While not the end of object
	for token.Type != TokenEndArray {
		value := parseValue(tokens, index)
		node.Value = append(node.Value, value)
		*index++
		token = (*tokens)[*index]
	}
	return node
}
