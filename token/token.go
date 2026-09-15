package token

type TokenType string

const (
	ILLEGAL   TokenType = "ILLEGAL"
	EOF       TokenType = "EOF"
	ASSIGN    TokenType = "ASSIGN"
	PLUS      TokenType = "PLUS"
	MINUS     TokenType = "MINUS"
	BANG      TokenType = "BANG"
	ASTERISK  TokenType = "ASTERISK"
	SLASH     TokenType = "SLASH"
	LT        TokenType = "LT"
	GT        TokenType = "GT"
	COMMA     TokenType = "COMMA"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	LBRACE    TokenType = "LBRACE"
	RBRACE    TokenType = "RBRACE"
	IDENT     TokenType = "IDENT"
	INT       TokenType = "INT"
	FUNCTION  TokenType = "FUNCTION"
	LET       TokenType = "LET"
	TRUE      TokenType = "TRUE"
	FALSE     TokenType = "FALSE"
	IF        TokenType = "IF"
	ELSE      TokenType = "ELSE"
	RETURN    TokenType = "RETURN"
	EQ        TokenType = "EQ"
	NEQ       TokenType = "NEQ"
)

const (
	LOWEST = iota
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
)

type Token struct {
	Type    TokenType
	Literal string
}

var Keywords = map[string]TokenType{
	"ফাঙ্কশন": FUNCTION,
	"ধর":      LET,
	"সত্য":    TRUE,
	"মিথ্যা":  FALSE,
	"যদি":     IF,
	"নাহলে":   ELSE,
	"ফেরাও":   RETURN,
}

func (t TokenType) IsInfix() bool {
	switch t {
	case PLUS, MINUS, ASTERISK, SLASH, LT, GT, EQ, NEQ:
		return true
	default:
		return false
	}
}

func (t TokenType) Precedence() int {
	switch t {
	case EQ, NEQ:
		return EQUALS
	case LT, GT:
		return LESSGREATER
	case PLUS, MINUS:
		return SUM
	case SLASH, ASTERISK:
		return PRODUCT
	default:
		return LOWEST
	}
}
