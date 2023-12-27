package Parser

import (
	"JsonParser/Token"
	"fmt"
)

func invalidJSONFormat(invalidToken token.Token) {
	fmt.Println("Invalid JSON format at ", invalidToken.Line, " at value ", invalidToken.Value)
}

func arrayParser(tokens []token.Token, start int) (bool, int) {
	valid := true
	i := start
	for valid == true && i < len(tokens) {
		if token.RIGHT_BRACKET == tokens[i].Type {
			i++
			return valid, i
		} else if token.STRING == tokens[i].Type || token.BOOLEAN == tokens[i].Type ||
			token.NUMBER == tokens[i].Type || token.NULL == tokens[i].Type {
			i++
			if token.COMMA == tokens[i].Type {
				i++
			}
		} else {
			invalidJSONFormat(tokens[i])
			valid = false
		}
	}
	return valid, i
}

func objectParser(tokens []token.Token, start int) (bool, int) {
	var valid = true
	i := start
	if start >= len(tokens) {
		valid = false
		fmt.Println("Invalid JSON Format")
	}
	for valid == true && i < len(tokens) {
		if token.CLOSED_BRACES == tokens[i].Type {
			i++
			if i < len(tokens) && token.COMMA == tokens[i].Type {
				i++
			}
			return valid, i
		} else if token.STRING == tokens[i].Type {
			i++
			if i < len(tokens) && token.COLON == tokens[i].Type {
				i++
				if i < len(tokens) {
					if token.OPEN_BRACES == tokens[i].Type {
						valid, i = objectParser(tokens, i+1)
					} else if token.LEFT_BRACKET == tokens[i].Type {
						valid, i = arrayParser(tokens, i+1)
					} else if token.STRING == tokens[i].Type || token.BOOLEAN == tokens[i].Type ||
						token.NUMBER == tokens[i].Type || token.NULL == tokens[i].Type {
						i++
						if i < len(tokens) && token.COMMA == tokens[i].Type {
							i++
						}
					} else {
						valid = false
						invalidJSONFormat(tokens[i])
					}
				}
			} else {
				valid = false
				invalidJSONFormat(tokens[i])
			}
		} else {
			valid = false
			invalidJSONFormat(tokens[i])
		}

	}
	return valid, i
}

func ParseJSON(tokens []token.Token) {
	valid := true
	for i := 0; i < len(tokens) && valid == true; i++ {
		switch tokens[i].Type {
		case token.OPEN_BRACES:
			valid, i = objectParser(tokens, i+1)
		case token.LEFT_BRACKET:
			valid, i = arrayParser(tokens, i+1)
		case token.COMMA:
			i++
		default:
			valid = false
			invalidJSONFormat(tokens[i])
		}
	}
	if valid == true {
		fmt.Println("JSON is Valid")
	}
}
