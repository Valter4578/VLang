package ast

import "github.com/valter4578/vlang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal value of the token it’s associated with Node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statementNode() {

}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {

}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) statementNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
