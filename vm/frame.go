package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	cl          *object.Closure
	ip          int
	basePointer int // Points to the bottom of ths stack for the current frame, a.k.a. frame pointer. It will be initialized with the current stack pointer.
}

func NewFrame(cl *object.Closure, basePointer int) *Frame {
	return &Frame{
		cl:          cl,
		ip:          -1,
		basePointer: basePointer,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
