package parser

import (
  "simplexSimia/ast"
  "simplexSimia/lexer"
  "simplexSimia/token"
)

type Parser struct {
  l *lexer.Lexer

  // points to current token
  curToken  token.Token
  // points to next token - peeks ahead
  peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
  p := &Parser{l: l}

  // reads two two tokens so that cur and peek tokens are both set
  p.nextToken()
  p.nextToken()

  return p
}

func (p *Parser) nextToken() {
  p.curToken = p.peekToken
  p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
  return nil
}
