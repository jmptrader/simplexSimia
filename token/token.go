package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	// EOL - standing for End Of Line, for detecting the end of a line. Used for line number and stricter syntac
	EOL 		= "EOL"
	EOF     = "EOF"

	// identifiers & literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS		 = "-"
	NOT			 = "!"
	ASTERISK = "*"
	SLASH		 = "/"

	EQ			 = "=="
	NOT_EQ	 = "!="

	LT			 = "<"
	GT 			 = ">"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	SIM      = "SIM"
	TRUE 		 = "TRUE"
	FALSE 	 = "FALSE"
	IF 			 = "IF"
	ELSE  	 = "ELSE"
	RETURN   = "RETURN"
)

// our keywords such as our funky function and our simple sim - for variables, something different
var keywords = map[string]TokenType {
	"fnc": 	  FUNCTION,
	"sim": 	  SIM,
	"true":   TRUE,
	"false":  FALSE,
	"if":		  IF,
	"else":   ELSE,
	"return": RETURN,
}



func LookupIdent(ident string) TokenType {
	// checks our keywords table to check whether the given identifier really is one of our keywords. I mean why wouldn't it be?
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}
