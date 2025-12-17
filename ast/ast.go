/*
Package ast
*/
package ast

import (
	"zebra/token"
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

// Program node will be the root node of every AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // The token.LET token)
	Name  *Identifier // e.g. variable name
	Value Expression  // The expression to be stored inthe variable
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

// Marker Method to declaer that a LetSTatement implements the statementNode interface
func (ls *LetStatement) statementNode() {}

// TokenLiteral Actually implement the TokenLiteral Method required for StatementNodes
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // The token.IDENT TOken
	Value string      // The variable name
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
