package lexer

import "simplexSimia/token"

type Lexer struct {
	input        string
	// current position in input (points to current character)
	position     int
	// current reading position in input (after current character)
	readPosition int
	// current character under examination
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// calls for the skipWhitespace function, removing whitespaces from the input - so readable tokens
	l.skipWhitespace()

	// case statement for each token - what should it do?
	switch l.ch {
	case '=':
		// if the assign token is followed by '='
		if l.peekChar() == '=' {
			// to store the current token
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		}else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		// if the not token is followed by "="
		if l.peekChar() == '=' {
			// to store the current token
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		}else {
			tok = newToken(token.NOT, l.ch)
		}
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
		// cheap way of getting the line number in error message. May give other options for future. Stands for End Of Line
	case '\n':
		tok = newToken(token.EOL, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF

	// defaultly set - so if none of the above
	default:
		// if it is a letter
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
			// if it is a digit
		}else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
			// if it doesn't know what the bloody hell you've typed
		}else {
			// it is an illegal character
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// does exactly what is says, skips whitespaces. Well, I excluded new lines but anyhow...
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// reads identifiers. 'no shit sherlock,' i hear you cry
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	// true if the token contains valuue ranging from 0 to 9
	return '0' <= ch && ch <= '9'
}

// what does this do?
func isLetter(ch byte)bool {
	// if ch ranges from a to z - byte - capitalized or not or if it is an '_' for example 'under_score' will do just fine
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 'peeks' ahead a charcter, not incrementing steps
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}else {
		return l.input[l.readPosition]
	}
}

// reads each character, changing position after
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	}else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
