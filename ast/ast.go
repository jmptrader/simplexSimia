package ast

import (
  "simplexSimia/token"
)

type Node interface {
  TokenLiteral() string
}

type Statement interface {
  Node
  statementNode()
}

type Expression interface {
  Node
  expressionNode()
}

type Program struct {
  Statements []Statement
}

type SimStatement struct {
  Token token.Token
  Name  *Identifier
  Value Expression
}

type Identifier struct {
  Token token.Token
  Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
  return i.Token.Literal
}

func (ls *SimStatement) statementNode() {}
func (ls *SimStatement) TokenLiteral() string {
  return ls.Token.Literal
}

func (p *Program) TokenLiteral() string {
  if len(p.Statements) > 0 {
    return p.Statements[0].TokenLiteral()
  }else {
    return ""
  }
}
