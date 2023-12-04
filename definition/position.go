package definition

import "fmt"

type Position interface {
	GetFile() string
	GetLine() int
	GetColumn() int
	GetPositionString() string
	Eqauals(Position) bool
}

type BasePosition struct {
	File   string
	Line   int
	Column int
}

func (p BasePosition) String() string {
	return fmt.Sprintf("%s:%d:%d", p.File, p.Line, p.Column+1)
}

func (p BasePosition) GetFile() string {
	return p.File
}

func (p BasePosition) GetLine() int {
	return p.Line
}

func (p BasePosition) GetColumn() int {
	return p.Column
}

func (p BasePosition) GetPositionString() string {
	return p.String()
}

func (p BasePosition) Eqauals(pos Position) bool {
	return p.File == pos.GetFile() && p.Line == pos.GetLine() && p.Column == pos.GetColumn()
}
