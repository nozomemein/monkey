package ast

import (
	"bytes"
	"monkey/token"
)

// AST Node
type Node interface {
	TokenLiteral() string
	String() string
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

// Program はASTのルートノード
type Program struct {
	Statements []Statement
}

// Program は Node Interface を満たす
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement は let文を表す構造体
type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 変数名
	Value Expression  // 式
}

// LetStatement は Statement Interface を満たす
func (l *LetStatement) StatementNode()       {}
func (l *LetStatement) TokenLiteral() string { return l.Token.Literal }
func (l *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(l.TokenLiteral() + " ")
	out.WriteString(l.Name.String())
	out.WriteString(" = ")

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

// Identifier は識別子を表す構造体
type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string      // 変数名
}

// Identifier は Expression Interface を満たす
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// ReturnStatement は return文を表す構造体
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// ReturnStatement は Statement Interface を満たす
func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement は式文を表す構造体
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression  // 式
}

// ExpressionStatement は Statement Interface を満たす
func (es *ExpressionStatement) StatementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
