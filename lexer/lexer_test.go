package lexer

import (
	"fmt"
	"testing"
	"io/ioutil"
	"simplexSimia/token"
	"strings"
)

func TestNextToken(t *testing.T) {
	fileNameDir := "test.sim"
	// tests whether file has .sim extension
	if strings.Contains(fileNameDir, ".sim") == true {
	}else {
		// if it doesn't, note that the wrong file type is present, giving advice, not significant, on possible solutions, two pathetic solutions
		t.Fatalf("%s is the wrong file type, either change the file extension or change speicifed file", fileNameDir)
	}
	inp, err := ioutil.ReadFile(fileNameDir)
  if err != nil {
      fmt.Print(err)
  }

	input := string(inp)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
		// expected tokens (type and literal)
	}{
		{token.SIM, "sim"},
		{token.IDENT, "test"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.SIM, "sim"},
		{token.IDENT, "hello"},
		{token.ASSIGN, "="},
		{token.INT, "43110"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.EOL, "\n"},
		{token.SIM, "sim"},
		{token.IDENT, "tot"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fnc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.EOL, "\n"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.EOL, "\n"},
		{token.SIM, "sim"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "tot"},
		{token.LPAREN, "("},
		{token.IDENT, "test"},
		{token.COMMA, ","},
		{token.IDENT, "hello"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.EOL, "\n"},
		{token.NOT, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.EOL, "\n"},
		{token.SIM, "sim"},
		{token.IDENT, "simia"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fnc"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.COMMA, ","},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.EOL, "\n"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.ASTERISK, "*"},
		{token.IDENT, "b"},
		{token.LT, "<"},
		{token.IDENT, "a"},
		{token.ASTERISK, "*"},
		{token.IDENT, "a"},
		{token.MINUS, "-"},
		{token.IDENT, "b"},
		{token.ASTERISK, "*"},
		{token.IDENT, "b"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.EOL, "\n"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.EOL, "\n"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.EOL, "\n"},
		{token.RBRACE, "}"},
		{token.EOL, "\n"},
		{token.RBRACE, "}"},
		{token.EOL, "\n"},
		{token.EOL, "\n"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "simia"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.NOT_EQ, "!="},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOL, "\n"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.IDENT, "simia"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.EQ, "=="},
		{token.TRUE, "true"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOL, "\n"},
		{token.EOF, ""},
	}
	// used to keep count of which line number it is on
  line := 1
	// used to keep count of how many tokens are on a single line
	tokCount := 0

	l := New(input)

	// actual token - type nor literal
	// actual := l.ch

	for i, tt := range tests {
		tok := l.NextToken()
		tokCount++

		// tests whether there is a new line or not
		if tok.Type == token.EOL || tokCount == 0 {
			line++
			tokCount = 0
		}

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q, on line:%d, token:%d",
				i, 										tt.expectedType, 				tok.Type, 							 line, 				tokCount)
			// current token position, expected token type, actual value of token, line number, token position on line
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q, on line:%d, token:%d",
				i, tt.expectedLiteral, tok.Literal, line, tokCount)
			// current token position, expected token type, actual value of token, line number, token position on line
		}
		// to update actual value - when the NextToken is called
		// actual = l.ch
	}
}
