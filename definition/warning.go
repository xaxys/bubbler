package definition

import (
	"fmt"

	"github.com/gookit/color"
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
	return fmt.Sprintf(
		"%s %s %s",
		color.FgLightWhite.Sprintf("%s:", w.GetPositionString()),
		color.FgYellow.Render("warning:"),
		w.Warning,
	)
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
	return fmt.Sprintf(
		"%s %s %s",
		color.FgLightWhite.Sprintf("%s:", w.GetPositionString()),
		color.FgYellow.Render("compile warning:"),
		w.Warning,
	)
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
	return fmt.Sprintf("'%s' may truncated from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.",
		color.FgLightWhite.Render(w.Expr),
		color.FgLightWhite.Render(w.From),
		color.FgLightWhite.Render(w.To),
	)
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
	return fmt.Sprintf("'%s' may overflowed from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.",
		color.FgLightWhite.Render(w.Expr),
		color.FgLightWhite.Render(w.From),
		color.FgLightWhite.Render(w.To),
	)
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
	return fmt.Sprintf("'%s' may lost data from '%s' to '%s', please mind the warning of compiler in generated code. Use cast expression if necessary.",
		color.FgLightWhite.Render(w.Expr),
		color.FgLightWhite.Render(w.From),
		color.FgLightWhite.Render(w.To),
	)
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
	return fmt.Sprintf("warnings occurred while importing '%s': %s",
		color.FgLightWhite.Render(w.File.Name),
		str,
	)
}

func (w ImportingWarning) Error() string {
	return w.String()
}

// --------------------

type OptionUnknownWarning struct {
	OptionName string
}

func (w OptionUnknownWarning) String() string {
	return fmt.Sprintf("unknown option '%s'",
		color.FgLightWhite.Render(w.OptionName),
	)
}

func (w OptionUnknownWarning) Error() string {
	return w.String()
}

// --------------------

type NameStyleWarning struct {
	Name string
	Msg  string
}

func (w NameStyleWarning) String() string {
	return fmt.Sprintf("unrecommended name style '%s': %s",
		color.FgLightWhite.Render(w.Name),
		w.Msg,
	)
}

func (w NameStyleWarning) Error() string {
	return w.String()
}
