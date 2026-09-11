package lexer

import (
	"strings"
	"lipi/token"
)

type Lexer struct {
	input        string
	position     int 
	ch           byte 
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, ch: input[0]}
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NEQ, Literal: "!="}
			l.readChar()
		} else {
			tok = newToken(token.BANG, l.ch)
		}
		tok = newToken(token.BANG, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			tok = l.nextWord()
		} else if isDigit(l.ch) {
			tok = l.nextInt()
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	l.position++
	if l.position >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.position]
	}
}

func (l *Lexer) peekChar() byte {
	if l.position+1 >= len(l.input) {
		return 0
	}
	return l.input[l.position+1]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func (l *Lexer) nextWord() token.Token {
	var ident strings.Builder
	for {
		ident.WriteString(string(l.ch))
		if !isLetter(l.peekChar()) && !isDigit(l.peekChar()) {
			break
		}
		l.readChar()
	}
	if tokType, ok := token.Keywords[ident.String()]; ok {
		return token.Token{Type: tokType, Literal: ident.String()}
	}
	return token.Token{Type: token.IDENT, Literal: ident.String()}
}

func (l *Lexer) nextInt() token.Token {
	var intVal strings.Builder
	for {
		intVal.WriteString(string(l.ch))
		if !isDigit(l.peekChar()) {
			break
		}
		l.readChar()
	}
	tok := token.Token{Type: token.INT, Literal: intVal.String()}
	return tok
}