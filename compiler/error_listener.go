package compiler

import (
	"errors"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/xaxys/bubbler/definition"
)

type errorListener struct {
	*antlr.DefaultErrorListener
	File string
	Err  error
}

func (l *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	err := &definition.SyntaxError{
		Position: definition.BasePosition{
			File:   l.File,
			Line:   line,
			Column: column,
		},
		Err: fmt.Errorf("%s", msg),
	}
	l.Err = errors.Join(l.Err, err)
}
