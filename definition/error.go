package definition

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/xaxys/bubbler/util"
)

type TopLevelError interface {
	error
	IsTopLevelError()
}

func TopLevelErrorsJoin(errs ...TopLevelError) TopLevelError {
	var topErrs []TopLevelError
	for _, err := range errs {
		if err == nil {
			continue
		}
		if topErr, ok := err.(*JoinTopLevelErrors); ok {
			topErrs = append(topErrs, topErr.Errs...)
		} else {
			topErrs = append(topErrs, err)
		}
	}
	if len(topErrs) == 0 {
		return nil
	}
	return &JoinTopLevelErrors{
		Errs: topErrs,
	}
}

// ==================== Top Level Error ====================

type JoinTopLevelErrors struct {
	TopLevelError
	Errs []TopLevelError
}

func (e JoinTopLevelErrors) String() string {
	strs := make([]string, len(e.Errs))
	for i, err := range e.Errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}

func (e JoinTopLevelErrors) Error() string {
	return e.String()
}

type GeneralError struct {
	TopLevelError
	Position
	Err error
}

func (e GeneralError) String() string {
	if e.Position == nil {
		return fmt.Sprintf(
			"%s %s",
			color.FgLightRed.Render("error:"),
			e.Err,
		)
	}
	return fmt.Sprintf(
		"%s %s %s",
		color.FgLightWhite.Sprintf("%s:", e.GetPositionString()),
		color.FgLightRed.Render("error:"),
		e.Err,
	)
}

func (e GeneralError) Error() string {
	return e.String()
}

// --------------------

type SyntaxError struct {
	TopLevelError
	Position
	Err error
}

func (e *SyntaxError) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		color.FgLightWhite.Sprintf("%s:", e.GetPositionString()),
		color.FgLightRed.Render("syntax error:"),
		e.Err,
	)
}

func (e SyntaxError) Error() string {
	return e.String()
}

// --------------------

type CompileError struct {
	TopLevelError
	Position
	Err error
}

func (e CompileError) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		color.FgLightWhite.Sprintf("%s:", e.GetPositionString()),
		color.FgLightRed.Render("compile error:"),
		e.Err,
	)
}

func (e CompileError) Error() string {
	return e.String()
}

// --------------------

type GenerateError struct {
	TopLevelError
	Err error
}

func (e GenerateError) String() string {
	return fmt.Sprintf(
		"%s %s",
		color.FgLightRed.Render("generate error:"),
		e.Err,
	)
}

func (e GenerateError) Error() string {
	return e.String()
}

// ==================== Specific Error ====================

type FileNotFoundError struct {
	File *FileIdentifer
	Err  error
}

func (e FileNotFoundError) String() string {
	return fmt.Sprintf("cannot find file '%s' (%s): %s",
		color.FgLightWhite.Render(e.File.Name),
		color.FgLightWhite.Render(e.File.Path),
		e.Err,
	)
}

func (e FileNotFoundError) Error() string {
	return e.String()
}

// --------------------

type FileReadError struct {
	File *FileIdentifer
	Err  error
}

func (e FileReadError) String() string {
	return fmt.Sprintf("cannot read file '%s' (%s): %s",
		color.FgLightWhite.Render(e.File.Name),
		color.FgLightWhite.Render(e.File.Path),
		e.Err,
	)
}

func (e FileReadError) Error() string {
	return e.String()
}

// --------------------

type FileWriteError struct {
	File *FileIdentifer
	Err  error
}

func (e FileWriteError) String() string {
	return fmt.Sprintf("cannot write file '%s' (%s): %s",
		color.FgLightWhite.Render(e.File.Name),
		color.FgLightWhite.Render(e.File.Path),
		e.Err,
	)
}

func (e FileWriteError) Error() string {
	return e.String()
}

// --------------------

type ImportingError struct {
	File *FileIdentifer
	Err  error
}

func (e ImportingError) String() string {
	str := "\n" + util.IndentSpace8(e.Err)
	return fmt.Sprintf("errors occurred while importing '%s':%s",
		color.FgLightWhite.Render(e.File.Name),
		str,
	)
}

func (e ImportingError) Error() string {
	return e.String()
}

// --------------------

type ImportSelfError struct{}

func (e ImportSelfError) String() string {
	return "self import detected"
}

func (e ImportSelfError) Error() string {
	return e.String()
}

// --------------------

type ImportCycleError struct {
	File *FileIdentifer
}

func (e ImportCycleError) String() string {
	return fmt.Sprintf("import cycle detected while importing '%s'",
		color.FgLightWhite.Render(e.File.Name),
	)
}

func (e ImportCycleError) Error() string {
	return e.String()
}

// --------------------

type DefinitionNotFoundError struct {
	DefName string // type, field or method name
}

func (e DefinitionNotFoundError) String() string {
	return fmt.Sprintf("definition of '%s' is not found",
		color.FgLightWhite.Render(e.DefName),
	)
}

func (e DefinitionNotFoundError) Error() string {
	return e.String()
}

// --------------------

type DefinitionTypeConflictError struct {
	DefName string
	Expect  string
	Got     string
}

func (e DefinitionTypeConflictError) String() string {
	return fmt.Sprintf("type conflict of definition '%s', found previous defined as '%s', but got '%s' here",
		color.FgLightWhite.Render(e.DefName),
		color.FgLightWhite.Render(e.Expect),
		color.FgLightWhite.Render(e.Got),
	)
}

func (e DefinitionTypeConflictError) Error() string {
	return e.String()
}

// --------------------

type DefinitionDuplicateError struct {
	DefName string // type, field or method name
	PrevDef Position
}

func (e DefinitionDuplicateError) String() string {
	return fmt.Sprintf(
		"duplicate definition of '%s', found previous definition at %s",
		color.FgLightWhite.Render(e.DefName),
		color.FgLightWhite.Render(e.PrevDef.GetPositionString()),
	)
}

func (e DefinitionDuplicateError) Error() string {
	return e.String()
}

// --------------------

type TypeUnopError struct {
	Expr string
	Type string
}

func (e TypeUnopError) String() string {
	return fmt.Sprintf("operator cannot be applied to '%s' of type '%s'",
		color.FgLightWhite.Render(e.Expr),
		color.FgLightWhite.Render(e.Type),
	)
}

func (e TypeUnopError) Error() string {
	return e.String()
}

// --------------------

type TypeBinopError struct {
	Expr1 string
	Expr2 string
	Type1 string
	Type2 string
}

func (e TypeBinopError) String() string {
	return fmt.Sprintf("operator cannot be applied to '%s' of type '%s', and '%s' of type '%s'",
		color.FgLightWhite.Render(e.Expr1),
		color.FgLightWhite.Render(e.Type1),
		color.FgLightWhite.Render(e.Expr2),
		color.FgLightWhite.Render(e.Type2),
	)
}

func (e TypeBinopError) Error() string {
	return e.String()
}

// --------------------

type TypeNotMatchError struct {
	Type1 string
	Type2 string
}

func (e TypeNotMatchError) String() string {
	return fmt.Sprintf("type not match between '%s' and '%s'",
		color.FgLightWhite.Render(e.Type1),
		color.FgLightWhite.Render(e.Type2),
	)
}

func (e TypeNotMatchError) Error() string {
	return e.String()
}

// --------------------

type EmbeddedFieldError struct {
	DefName string
	Err     error
}

func (e EmbeddedFieldError) String() string {
	str := "\n" + util.IndentSpace8(e.Err)
	return fmt.Sprintf("errors occurred while parsing embedded field '%s': %s",
		color.FgLightWhite.Render(e.DefName),
		str,
	)
}

func (e EmbeddedFieldError) Error() string {
	return e.String()
}

// --------------------

type PackageDefinitionDuplicateError struct {
	PrevDef     Position
	PackageName string
}

func (e PackageDefinitionDuplicateError) String() string {
	return fmt.Sprintf("package name has already been set to '%s' at %s",
		color.FgLightWhite.Render(e.PackageName),
		color.FgLightWhite.Render(e.PrevDef.GetPositionString()),
	)
}

func (e PackageDefinitionDuplicateError) Error() string {
	return e.String()
}

// --------------------

type PackageDuplicateError struct {
	PrevDef Position
	Package *Package
}

func (e PackageDuplicateError) String() string {
	return fmt.Sprintf("duplicate package name '%s', found previous one at %s",
		color.FgLightWhite.Render(e.Package.String()),
		color.FgLightWhite.Render(e.PrevDef),
	)
}

func (e PackageDuplicateError) Error() string {
	return e.String()
}

// --------------------

type PackageNameNotSetError struct{}

func (e PackageNameNotSetError) String() string {
	return fmt.Sprintf("package name not set")
}

func (e PackageNameNotSetError) Error() string {
	return e.String()
}

// --------------------

type OptionDuplicateError struct {
	PrevDef    Position
	OptionName string
}

func (e OptionDuplicateError) String() string {
	return fmt.Sprintf("option '%s' has already been set at %s",
		color.FgLightWhite.Render(e.OptionName),
		color.FgLightWhite.Render(e.PrevDef.GetPositionString()),
	)
}

func (e OptionDuplicateError) Error() string {
	return e.String()
}

// --------------------

type OptionTypeError struct {
	OptionName string
	Expect     string
	Got        string
}

func (e OptionTypeError) String() string {
	return fmt.Sprintf("option '%s' type '%s' is not valid, expect '%s'",
		color.FgLightWhite.Render(e.OptionName),
		color.FgLightWhite.Render(e.Got),
		color.FgLightWhite.Render(e.Expect),
	)
}

func (e OptionTypeError) Error() string {
	return e.String()
}

// --------------------

type OptionValueError struct {
	OptionName string
	Expect     []any
	Got        any
}

func (e OptionValueError) String() string {
	expects := make([]string, len(e.Expect))
	for i, expect := range e.Expect {
		expects[i] = fmt.Sprint(expect)
	}
	expect := strings.Join(expects, ", ")
	got := fmt.Sprint(e.Got)
	return fmt.Sprintf("option '%s' value '%s' is not valid, expect one of [%s]",
		color.FgLightWhite.Render(e.OptionName),
		color.FgLightWhite.Render(got),
		color.FgLightWhite.Render(expect),
	)
}

func (e OptionValueError) Error() string {
	return e.String()
}

// --------------------

type FieldNotAlignedError struct {
	FieldName string
	DynStart  bool
	Start     int64
	Msg       string
}

func (e FieldNotAlignedError) String() string {
	if e.DynStart {
		return fmt.Sprintf("field '%s' start at dynamic+%s is not aligned: %s",
			color.FgLightWhite.Render(e.FieldName),
			color.FgLightWhite.Render(util.ToSizeString(e.Start)),
			e.Msg,
		)
	}
	return fmt.Sprintf("field '%s' start at %s is not aligned: %s",
		color.FgLightWhite.Render(e.FieldName),
		color.FgLightWhite.Render(util.ToSizeString(e.Start)),
		e.Msg,
	)
}

func (e FieldNotAlignedError) Error() string {
	return e.String()
}

// --------------------

type FieldConstValueTypeError struct {
	Constant string
	Expect   string
	Got      string
}

func (e FieldConstValueTypeError) String() string {
	return fmt.Sprintf("constant value '%s' of type '%s' does not match field type '%s'",
		color.FgLightWhite.Render(e.Constant),
		color.FgLightWhite.Render(e.Got),
		color.FgLightWhite.Render(e.Expect),
	)
}

func (e FieldConstValueTypeError) Error() string {
	return e.String()
}

// --------------------

type MethodTypeError struct {
	MethodName string
	PrevDef    Position
	Expect     string
	Got        string
}

func (e MethodTypeError) String() string {
	return fmt.Sprintf("method '%s' type '%s' does not match '%s', found previous definition at %s",
		color.FgLightWhite.Render(e.MethodName),
		color.FgLightWhite.Render(e.Got),
		color.FgLightWhite.Render(e.Expect),
		color.FgLightWhite.Render(e.PrevDef.GetPositionString()),
	)
}

func (e MethodTypeError) Error() string {
	return e.String()
}

// --------------------

type EnumConstValueTypeError struct {
	Constant string
	Got      string
}

func (e EnumConstValueTypeError) String() string {
	return fmt.Sprintf("invalid constant value '%s' of type '%s', enum value must be an integer",
		color.FgLightWhite.Render(e.Constant),
		color.FgLightWhite.Render(e.Got),
	)
}

func (e EnumConstValueTypeError) Error() string {
	return e.String()
}

// --------------------

type InvalidConstIdentifierError struct {
	Name    string
	PrevDef Position
}

func (e InvalidConstIdentifierError) String() string {
	return fmt.Sprintf("identifier '%s' is not a valid constant, previous definition at %s",
		color.FgLightWhite.Render(e.Name),
		color.FgLightWhite.Render(e.PrevDef.GetPositionString()),
	)
}

func (e InvalidConstIdentifierError) Error() string {
	return e.String()
}

// --------------------

type InvalidSizeError struct {
	Size int64
	Msg  string
}

func (e InvalidSizeError) String() string {
	return fmt.Sprintf("invalid size [%s] (%s bits): %s",
		color.FgLightWhite.Render(util.ToSizeString(e.Size)),
		color.FgLightWhite.Sprint(e.Size),
		e.Msg,
	)
}

func (e InvalidSizeError) Error() string {
	return e.String()
}

// --------------------

type InvalidLiteralError struct {
	Literal string
	Err     error
}

func (e InvalidLiteralError) String() string {
	return fmt.Sprintf("invalid literal '%s': %s",
		color.FgLightWhite.Render(e.Literal),
		e.Err,
	)
}

func (e InvalidLiteralError) Error() string {
	return e.String()
}

// --------------------

type InvalidEnumDefError struct {
	DefName string
	Err     error
}

func (e InvalidEnumDefError) String() string {
	return fmt.Sprintf("invalid enum '%s': %s",
		color.FgLightWhite.Render(e.DefName),
		e.Err,
	)
}

func (e InvalidEnumDefError) Error() string {
	return e.String()
}

// --------------------

type InvalidStructDefError struct {
	DefName string
	Err     error
}

func (e InvalidStructDefError) String() string {
	return fmt.Sprintf("invalid struct '%s': %s",
		color.FgLightWhite.Render(e.DefName),
		e.Err,
	)
}

func (e InvalidStructDefError) Error() string {
	return e.String()
}

// --------------------

type InvalidEmbeddedFieldError struct{}

func (e InvalidEmbeddedFieldError) String() string {
	return "invalid embedded field, type must be a struct"
}

func (e InvalidEmbeddedFieldError) Error() string {
	return e.String()
}

// --------------------

type InvalidFieldError struct {
	FieldName string
	Msg       string
}

func (e InvalidFieldError) String() string {
	return fmt.Sprintf("invalid field '%s': %s",
		color.FgLightWhite.Render(e.FieldName),
		e.Msg,
	)
}

func (e InvalidFieldError) Error() string {
	return e.String()
}

// --------------------

type InvalidExprError struct {
	Expr string
	Msg  string
}

func (e InvalidExprError) String() string {
	if e.Expr == "" {
		return fmt.Sprintf("invalid expression: %s", e.Msg)
	}
	return fmt.Sprintf("invalid expression '%s': %s",
		color.FgLightWhite.Render(e.Expr),
		e.Msg,
	)
}

func (e InvalidExprError) Error() string {
	return e.String()
}

// --------------------

type NameStyleError struct {
	Name string
	Msg  string
}

func (e NameStyleError) String() string {
	return fmt.Sprintf("invalid name style '%s': %s",
		color.FgLightWhite.Render(e.Name),
		e.Msg,
	)
}

func (e NameStyleError) Error() string {
	return e.String()
}

// --------------------

type TargetNotSupportedError struct {
	Target string
}

func (e TargetNotSupportedError) String() string {
	return fmt.Sprintf("target '%s' not supported", e.Target)
}

func (e TargetNotSupportedError) Error() string {
	return e.String()
}

// --------------------

type TargetNotSpecifiedError struct{}

func (e TargetNotSpecifiedError) String() string {
	return fmt.Sprintf("target not specified")
}

func (e TargetNotSpecifiedError) Error() string {
	return e.String()
}

// --------------------

type NoInputFileError struct{}

func (e NoInputFileError) String() string {
	return fmt.Sprintf("no input file specified")
}

func (e NoInputFileError) Error() string {
	return e.String()
}

// --------------------

type MultipleInputFileError struct {
	Files []string
}

func (e MultipleInputFileError) String() string {
	hint := false
	for i, file := range e.Files {
		e.Files[i] = color.FgLightWhite.Render(file)
		if strings.HasPrefix(file, "-") {
			hint = true
		}
	}
	str := fmt.Sprintf("only one input file is allowed, but got: %s", strings.Join(e.Files, ", "))
	if hint {
		str += "\n" + "(please notice that all '-xxx' options should be placed before input file, or they will be treated as input file)"
	}
	return str
}

func (e MultipleInputFileError) Error() string {
	return e.String()
}
