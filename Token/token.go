package token

const (
	OPEN_BRACES   Type = "{"
	CLOSED_BRACES Type = "}"
	LEFT_BRACKET  Type = "["
	RIGHT_BRACKET Type = "]"
	COMMA         Type = ","
	COLON         Type = ":"
	STRING        Type = "STRING"
	NUMBER        Type = "NUMBER"
	BOOLEAN       Type = "BOOLEAN"
	NULL          Type = "NULL"
	ILLEGAL       Type = "ILLEGAL"
)

type Type string

type Token struct {
	Type  Type
	Value string
	Line  int
	Start int
	End   int
}
