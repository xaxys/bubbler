package compiler

import "github.com/xaxys/bubbler/parser"

type PrintVisitor struct {
	parser.BasebubblerVisitor
}

func NewPrintVisitor() *PrintVisitor {
	return &PrintVisitor{}
}
