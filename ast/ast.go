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
// x + 10; など
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

// IntegerLiteral は整数リテラルを表す構造体
// 5, 10, 993322 など
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// IntegerLiteral は Expression Interface を満たす
func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression は前置演算子を表す構造体
type PrefixExpression struct {
	Token    token.Token // 前置トークン、例えば '!'
	Operator string
	Right    Expression
}

// PrefixExpression は Expression Interface を満たす
func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

// InfixExpression は中置演算子を表す構造体
type InfixExpression struct {
	Token    token.Token // 中置トークン、例えば '+'
	Left     Expression
	Operator string
	Right    Expression
}

// InfixExpression は Expression Interface を満たす
func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
  var out bytes.Buffer
  out.WriteString("(")
  out.WriteString(ie.Left.String())
  out.WriteString(" " + ie.Operator + " ")
  out.WriteString(ie.Right.String())
  out.WriteString(")")
  return out.String()
}