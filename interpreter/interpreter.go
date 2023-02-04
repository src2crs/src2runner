package interpreter

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// TODO: Add documentation

type Interpreter struct {
	i          *interp.Interpreter
	lastError  error
	lastResult interface{}
}

func New() *Interpreter {
	i := &Interpreter{interp.New(interp.Options{}), nil, nil}
	i.i.Use(stdlib.Symbols)
	return i
}

func (interpreter *Interpreter) Evaluate(source string) {
	res, err := interpreter.i.Eval(source)
	if err == nil {
		interpreter.lastResult = res.Interface()
	}
	interpreter.lastError = err
}

func (interpreter *Interpreter) LastResult() interface{} {
	return interpreter.lastResult
}

func (interpreter *Interpreter) LastError() interface{} {
	return interpreter.lastError
}

func (interpreter *Interpreter) Ok() bool {
	return interpreter.lastError == nil
}
