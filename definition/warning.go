package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

type TopLevelWarning interface {
	error
	IsTopLevelWarning()
}

// ==================== Top Level Warning ====================

type GeneralWarning struct {
	TopLevelWarning
	Position
	Warning error
}

func (w GeneralWarning) String() string {
	return fmt.Sprintf("%s: warning: %s", w.GetPositionString(), w.Warning.Error())
}

func (w GeneralWarning) Error() string {
	return w.String()
}

// --------------------

type CompileWarning struct {
	TopLevelWarning
	Position
	Warning error
}

func (w CompileWarning) String() string {
	return fmt.Sprintf("%s: compile warning: %s", w.GetPositionString(), w.Warning.Error())
}

func (w CompileWarning) Error() string {
	return w.String()
}

// ==================== Specific Warning ====================

type ArithTruncationWarning struct {
	Expr string
	From string
	To   string
}

func (w ArithTruncationWarning) String() string {
	return fmt.Sprintf("'%s' may truncated from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.", w.Expr, w.From, w.To)
}

func (w ArithTruncationWarning) Error() string {
	return w.String()
}

// --------------------

type ArithOverflowWarning struct {
	Expr string
	From string
	To   string
}

func (w ArithOverflowWarning) String() string {
	return fmt.Sprintf("'%s' may overflowed from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.", w.Expr, w.From, w.To)
}

func (w ArithOverflowWarning) Error() string {
	return w.String()
}

// --------------------

type ArithSignToUnsignWarning struct {
	Expr string
	From string
	To   string
}

func (w ArithSignToUnsignWarning) String() string {
	return fmt.Sprintf("'%s' may lost data from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.", w.Expr, w.From, w.To)
}

func (w ArithSignToUnsignWarning) Error() string {
	return w.String()
}

// --------------------

type ImportingWarning struct {
	File    *FileIdentifer
	Warning error
}

func (w ImportingWarning) String() string {
	str := "\n" + util.IndentSpace8(w.Warning)
	return fmt.Sprintf("warnings occurred while importing '%s':%s", w.File.Name, str)
}

func (w ImportingWarning) Error() string {
	return w.String()
}

// --------------------

type OptionUnknownWarning struct {
	OptionName string
}

func (w OptionUnknownWarning) String() string {
	return fmt.Sprintf("unknown option '%s'", w.OptionName)
}

func (w OptionUnknownWarning) Error() string {
	return w.String()
}
