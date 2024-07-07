package object

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}

}

// 現在の環境を包含する新しい環境を作成する
// 関数呼び出しなどの変数の束縛は、この新しい環境に保存することで、元の環境が変更されないようにする
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
