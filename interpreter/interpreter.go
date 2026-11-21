package interpreter

import (
	"errors"
	"strings"

	"strawberryInterpreter/evaluator"
	"strawberryInterpreter/lexer"
	"strawberryInterpreter/parser"
)

type ParseError struct {
	Messages []string
}

func (e *ParseError) Error() string {
	return strings.Join(e.Messages, "\n")
}

func Run(source string) (string, error) {
	p := parser.New(lexer.New(source))
	program := p.ParseProgram()
	if messages := p.Errors(); len(messages) > 0 {
		return "", &ParseError{Messages: messages}
	}

	if len(program.Statements) == 0 {
		return "", nil
	}

	result := evaluator.Eval(program)
	if result == nil {
		return "", errors.New("this statement is not implemented by the evaluator yet")
	}

	return result.Inspect(), nil
}
