package ast

import (
	"monkey/token"
)

// AST Node
type Node interface {
	TokenLiteral() string
}

// Statement(文) Node
type Statement interface {
	Node
	StatementNode()
}

// Expression(式) Node
type Expression interface {
	Node
	expressionNode()
}

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

// LetStatement は let文を表す構造体
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 変数名
	Value Expression  // 式
}

// Statement Interface の実装
func (l *LetStatement) StatementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }

// Identifier は識別子を表す構造体
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string      // 変数名
}

// Expression Interface の実装
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
