package format

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

/*
my_value: "test",
object: {
	my_object_value: 1,
}
*/

// JsonEntry describes an entry in a json-object.
type JsonEntry struct {
	name  string
	value reflect.Value
}

// JsonObject describes a basic json object, with entries.
type JsonObject struct {
	items []JsonEntry
}

// All tokens
var tokens []TokenMatch

// The current token index
var currentIndex = -1

// The current token
var currentToken TokenMatch

// advancing the current token
func nextToken() (TokenMatch, error) {
	if currentIndex+1 >= len(tokens) {
		return TokenMatch{}, errors.New("the index is bigger than the length of the tokens")
	}

	currentIndex++

	return tokens[currentIndex], nil
}

// Expecting a certain token
func expect(tokenType int) TokenMatch {
	if currentToken.tokenType != tokenType {
		fmt.Printf("Expected token with type of %i, got %i", tokenType, currentToken.tokenType)
		os.Exit(2)
	}

	token, err := nextToken()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	currentToken = token

	return currentToken
}

// value of an entry.
func value() (reflect.Value, error) {
	switch currentToken.tokenType {
	case IntegerLiteral:
		asInt, err := strconv.Atoi(currentToken.raw)

		// Somehow, the int isn't an int
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		expect(IntegerLiteral)

		return reflect.ValueOf(asInt), nil

	case StringLiteral:
		expect(StringLiteral)

		return reflect.ValueOf(currentToken.raw), nil
	}

	return reflect.Value{}, errors.New("this token wasn't found")
}

// We probably have found an entry
func entry() JsonEntry {
	entryName := currentToken.raw

	expect(Identifier)
	expect(Colon)

	value, err := value()

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	return JsonEntry{name: entryName, value: value}
}

// ToJsonObject Transforming a string to a JsonObject
func ToJsonObject(input string) (JsonObject, error) {
	var entries []JsonEntry
	tokens = tokenize(input)

	fmt.Println(tokens)

	currentToken, _ = nextToken()

	for currentToken.tokenType != EndOfInput {
		entries = append(entries, entry())
	}

	return JsonObject{items: entries}, nil
}
