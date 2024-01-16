package ast

import (
	"strawberryInterpreter/token"
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

type LetStatement struct {
	Token token.Token // this is the token.LET token
	Name  *Identifier
	Value Expression
}

type Identifier struct {
	Token token.Token // this is the token.IDENT token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (ls *LetStatement) statementNode() {
}

func (i *Identifier) expressionNode() {
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
