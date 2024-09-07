package compiler

import (
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

// CompilerはASTを受け取り、それをバイトコードにコンパイルする
type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
	}
}

func (c *Compiler) Bytecode() *Bytecode {
	return &Bytecode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

// Compilerが生成したバイトコードを返すための構造体
type Bytecode struct {
	Instructions code.Instructions
	Constants    []object.Object
}
