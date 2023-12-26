package main

import (
	"JsonParser/Lexer"
	"JsonParser/Parser"
)

func validateJSON(input string) {
	var tokens = Lexer.Tokenization(input)
	Parser.ParseJSON(tokens)
}

func main() {
	input := "{[]}"
	validateJSON(input)
}
