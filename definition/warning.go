package definition

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/xaxys/bubbler/util"
)

type TopLevelWarning interface {
	error
	IsTopLevelWarning()
}

func TopLevelWarningsJoin(warnings ...TopLevelWarning) TopLevelWarning {
	var topWarnings []TopLevelWarning
	for _, warning := range warnings {
		if warning == nil {
			continue
		}
		if topWarning, ok := warning.(*JoinTopLevelWarnings); ok {
			topWarnings = append(topWarnings, topWarning.Warnings...)
		} else {
			topWarnings = append(topWarnings, warning)
		}
	}
	if len(topWarnings) == 0 {
		return nil
	}
	return &JoinTopLevelWarnings{
		Warnings: topWarnings,
	}
}

// ==================== Top Level Warning ====================

type JoinTopLevelWarnings struct {
	TopLevelWarning
	Warnings []TopLevelWarning
}

func (w JoinTopLevelWarnings) String() string {
	strs := make([]string, len(w.Warnings))
	for i, warning := range w.Warnings {
		strs[i] = warning.Error()
	}
	return strings.Join(strs, "\n")
}

func (w JoinTopLevelWarnings) Error() string {
	return w.String()
}

type GeneralWarning struct {
	TopLevelWarning
	Position
	Warning error
}

func (w GeneralWarning) String() string {
	if w.Position == nil {
		return fmt.Sprintf(
			"%s %s",
			color.FgYellow.Render("warning:"),
			w.Warning,
		)
	}
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
		color.FgLightYellow.Render("compile warning:"),
		w.Warning,
	)
}

func (w CompileWarning) Error() string {
	return w.String()
}

// --------------------

type GenerateWarning struct {
	TopLevelWarning
	Warning error
}

func (w GenerateWarning) String() string {
	return fmt.Sprintf(
		"%s %s",
		color.FgLightYellow.Render("generate warning:"),
		w.Warning,
	)
}

func (w GenerateWarning) Error() string {
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

type PackageNameNotSetWarning struct{}

func (w PackageNameNotSetWarning) String() string {
	return fmt.Sprintf("package name not set")
}

func (w PackageNameNotSetWarning) Error() string {
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

type OptionNotSetWarning struct {
	OptionName string
	Reason     string
}

func (w OptionNotSetWarning) String() string {
	if w.Reason == "" {
		return fmt.Sprintf("option '%s' not set in current settings",
			color.FgLightWhite.Render(w.OptionName),
		)
	}
	return fmt.Sprintf("option '%s' not set: %s",
		color.FgLightWhite.Render(w.OptionName),
		w.Reason,
	)
}

func (w OptionNotSetWarning) Error() string {
	return w.String()
}

// --------------------

type OptionNotAvailableWarning struct {
	OptionName string
	Reason     string
}

func (w OptionNotAvailableWarning) String() string {
	if w.Reason == "" {
		return fmt.Sprintf("option '%s' not available in current settings",
			color.FgLightWhite.Render(w.OptionName),
		)
	}
	return fmt.Sprintf("option '%s' not available: %s",
		color.FgLightWhite.Render(w.OptionName),
		w.Reason,
	)
}

func (w OptionNotAvailableWarning) Error() string {
	return w.String()
}

// --------------------

type NameStyleWarning struct {
	Name string
	Msg  string
}

func (w NameStyleWarning) String() string {
	return fmt.Sprintf("non-recommended name style '%s': %s",
		color.FgLightWhite.Render(w.Name),
		w.Msg,
	)
}

func (w NameStyleWarning) Error() string {
	return w.String()
}

// --------------------

type NameStyleNotStandardWarning struct {
	OriginName string
	RecommName string
	Standard   string
}

func (w NameStyleNotStandardWarning) String() string {
	return fmt.Sprintf("non-standard %s '%s', use '%s' instead",
		w.Standard,
		color.FgLightWhite.Render(w.OriginName),
		color.FgLightWhite.Render(w.RecommName),
	)
}

func (w NameStyleNotStandardWarning) Error() string {
	return w.String()
}
