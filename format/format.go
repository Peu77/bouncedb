package format

import (
	"errors"
	"fmt"
	"go/types"
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
	rtype types.BasicKind
	value interface{}
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

// peeking the next token
func peekToken() (TokenMatch, error) {
	if currentIndex+1 >= len(tokens) {
		return TokenMatch{}, errors.New("the index is bigger than the length of the tokens")
	}

	return tokens[currentIndex+1], nil
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
func value() (interface{}, error) {
	switch currentToken.tokenType {
	case IntegerLiteral:
		asInt, err := strconv.Atoi(currentToken.raw)

		// Somehow, the int isn't an int
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}

		expect(IntegerLiteral)

		return asInt, nil

	case StringLiteral:
		raw := currentToken.raw

		expect(StringLiteral)

		return raw, nil
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
func ToJsonObject(input string) JsonObject {
	var entries []JsonEntry
	tokens = tokenize(input)

	currentToken, _ = nextToken()

	expect(LeftBracket)

	for currentToken.tokenType != EndOfInput {
		entries = append(entries, entry())
		peek, err := peekToken()

		if err == nil || peek.tokenType == Comma {
			expect(Comma)
		}
	}

	expect(RightBracket)

	return JsonObject{items: entries}
}

// FromJsonObject converts our json object to a string
func FromJsonObject(object JsonObject) string {
	raw := "{\n"

	for _, item := range object.items {
		switch reflect.ValueOf(item.value).Type().Name() {
		case "string":
			raw += item.name + ": " + fmt.Sprint(item.value) + ", \n"
			break
		case "int":
			raw += fmt.Sprintf(item.name+": %d, \n", item.value.(int))
			break
		}
	}

	return raw + "\n}"
}

// FromStructToJson converts a struct to a string
func FromStructToJson(s interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(s))
	raw := "{\n"

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		ending := ",\n"

		if i+1 == v.NumField() {
			ending = ""
		}

		raw += "\t" + v.Type().Field(i).Name + ": " + fmt.Sprint(fieldValue) + ending
	}

	return raw + "\n}"
}
