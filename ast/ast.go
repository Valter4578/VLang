package ast

import "github.com/valter4578/vlang/token"

type Node interface {
	TokenLiteral() string
}

type Statment interface {
	Node
	statmentNode()
}

type Expression interface {
	Node
	statmentNode()
}

type Program struct {
	Statments []Statment
}

type LetStatement struct {
	Token token.Token // LET token
	Name  *Identifier
	Value Expression
}

// func (ls *LetStatement) statementNode() {

// }

// func (ls *LetStatement) TokenLiteral() string {
// 	return ls.Token.Literal
// }

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) statmentNode() {

}

type Identifier struct {
	Token token.Token // IDENT token
	Value string
}

func (i *Identifier) statmentNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// TokenLiteral returns the literal value of the token it’s associated with Node
func (p *Program) TokenLiteral() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteral()
	} else {
		return ""
	}
}
