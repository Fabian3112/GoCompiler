package main

type Token struct {
	tokenType tokenType
	value     string
}

func (token *Token) ranking() int {
	switch token.tokenType {
	case SemikononToken:
		return -1
	case PrintToken, WhileToken, IfToken, ElseToken, DeclarationToken, AssignToken:
		return 0
	case MultToken:
		return 5
	case PlusToken:
		return 4
	case AndToken:
		return 2
	case OrToken:
		return 1
	case NotToken:
		return 6
	case EqualToken:
		return 3
	case LesserToken:
		return 3
	default:
		return -42
	}
}

type tokenType string

const (
	MultToken    tokenType = "*"
	PlusToken    tokenType = "+"
	AndToken     tokenType = "&&"
	OrToken      tokenType = "||"
	NotToken     tokenType = "!"
	EqualToken   tokenType = "=="
	LesserToken  tokenType = "<"
	BooleanToken tokenType = "boolean"

	NumberToken tokenType = "Number"
	VarToken    tokenType = "Variable"

	GroupOpenToken  tokenType = "("
	GroupCloseToken tokenType = ")"
	BlockOpenToken  tokenType = "{"
	BlockCloseToken tokenType = "}"

	PrintToken tokenType = "print"
	WhileToken tokenType = "while"
	IfToken    tokenType = "if"
	ElseToken  tokenType = "else"

	DeclarationToken tokenType = ":="
	SemikononToken   tokenType = ";"
	AssignToken      tokenType = "=" //check last so := == is already tokenized ?? Doesnt work
)
