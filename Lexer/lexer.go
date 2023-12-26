package Lexer

import (
	"JsonParser/Token"
	"strings"
)

func isNumber(char rune) bool {
	return '0' <= char && char <= '9'
}

func isLetter(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func parseNumber(input string, pos int) (string, int) {
	var (
		value string
		end   = pos + 1
	)

	for i := pos; i < len(input) && isNumber(rune(input[i])); i++ {
		value += string(input[i])
		end = i
	}
	return value, end + 1
}

func parseLetter(input string, pos int) (string, int) {
	var (
		value string
		end   = pos + 1
	)

	for i := pos; i < len(input) && isLetter(rune(input[i])); i++ {
		value += string(input[i])
		end = i
	}
	return value, end + 1
}

func parseString(input string, pos int) (string, int) {
	var (
		value = ""
		end   = pos + 1
	)

	for i := pos + 1; input[i] != '"'; i++ {
		value += string(input[i])
		end = i
	}
	return value, end + 2
}

func newToken(tokenType token.Type, value string, line int, start int, end int) token.Token {
	return token.Token{
		Type:  tokenType,
		Value: value,
		Line:  line,
		Start: start,
		End:   end,
	}
}

func createToken(input string, pos int, line int) token.Token {
	var t token.Token

	switch input[pos] {
	case '{':
		t = newToken(token.OPEN_BRACES, "{", line, pos, pos+1)
	case '}':
		t = newToken(token.CLOSED_BRACES, "}", line, pos, pos+1)
	case '[':
		t = newToken(token.LEFT_BRACKET, "[", line, pos, pos+1)
	case ']':
		t = newToken(token.RIGHT_BRACKET, "]", line, pos, pos+1)
	case ':':
		t = newToken(token.COLON, ":", line, pos, pos+1)
	case ',':
		t = newToken(token.COMMA, ",", line, pos, pos+1)
	case '"':
		str, end := parseString(input, pos)
		t = newToken(token.STRING, str, line, pos, end)
	default:
		if isNumber(rune(input[pos])) {
			num, end := parseNumber(input, pos)
			return newToken(token.NUMBER, num, line, pos, end)
		} else if isLetter(rune(input[pos])) {
			letter, end := parseLetter(input, pos)
			if letter == "true" {
				return newToken(token.NUMBER, "true", line, pos, end)
			} else if letter == "false" {
				return newToken(token.BOOLEAN, "false", line, pos, end)
			} else if letter == "null" {
				return newToken(token.NULL, "null", line, pos, end)
			}
		}
		t = newToken(token.ILLEGAL, string(input[pos]), line, pos, pos+1)

	}
	return t
}

func isWhiteSpace(input string) bool {
	return input == " " || input == "\t" || input == "\n" || input == "\r"
}

func Tokenization(input string) []token.Token {
	lines := strings.Split(input, "\n")
	var tokens []token.Token
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		charArray := strings.Split(line, "")
		for j := 0; j < len(charArray); {
			if isWhiteSpace(charArray[j]) {
				j++
			} else {
				var t = createToken(line, j, i)
				j = t.End
				tokens = append(tokens, t)
			}
		}
	}
	return tokens
}
