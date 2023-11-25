package definition

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/util"
)

type TopLevelError interface {
	error
	IsTopLevelError()
}

// ==================== Top Level Error ====================

type GeneralError struct {
	TopLevelError
	Position
	Err error
}

func (e GeneralError) String() string {
	return fmt.Sprintf("%s: error: %s", e.GetPositionString(), e.Err)
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
	return fmt.Sprintf("%s: syntax error: %s", e.GetPositionString(), e.Err)
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
	return fmt.Sprintf("%s: compile error: %s", e.GetPositionString(), e.Err)
}

func (e CompileError) Error() string {
	return e.String()
}

// ==================== Specific Error ====================

type FileNotFoundError struct {
	File *FileIdentifer
	Err  error
}

func (e FileNotFoundError) String() string {
	return fmt.Sprintf("cannot find file '%s' (%s): %s", e.File.Name, e.File.Path, e.Err)
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
	return fmt.Sprintf("cannot read file '%s' (%s): %s", e.File.Name, e.File.Path, e.Err)
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
	return fmt.Sprintf("cannot write file '%s' (%s): %s", e.File.Name, e.File.Path, e.Err)
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
	return fmt.Sprintf("errors occurred while importing '%s':%s", e.File.Name, str)
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
	return fmt.Sprintf("import cycle detected while importing '%s'", e.File.Name)
}

func (e ImportCycleError) Error() string {
	return e.String()
}

// --------------------

type DefinitionNotFoundError struct {
	DefName string // type, field or method name
}

func (e DefinitionNotFoundError) String() string {
	return fmt.Sprintf("definition of '%s' is not found", e.DefName)
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
	return fmt.Sprintf("type conflict of definition '%s', found previous defined as '%s', but got '%s' here", e.DefName, e.Expect, e.Got)
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
		e.DefName,
		e.PrevDef.GetPositionString(),
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
	return fmt.Sprintf("operator cannot be applied to '%s' of type '%s'", e.Expr, e.Type)
}

func (e TypeUnopError) Error() string {
	return e.String()
}

type TypeBinopError struct {
	Expr1 string
	Expr2 string
	Type1 string
	Type2 string
}

func (e TypeBinopError) String() string {
	return fmt.Sprintf("operator cannot be applied to '%s' of type '%s', and '%s' of type '%s'", e.Expr1, e.Type1, e.Expr2, e.Type2)
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
	return fmt.Sprintf("type not match between '%s' and '%s'", e.Type1, e.Type2)
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
	return fmt.Sprintf("errors occurred while parsing embedded field '%s':%s", e.DefName, str)
}

func (e EmbeddedFieldError) Error() string {
	return e.String()
}

// --------------------

type OptionTypeError struct {
	OptionName string
	Expect     string
	Got        string
}

func (e OptionTypeError) String() string {
	return fmt.Sprintf("option '%s' type '%s' is not valid, expect '%s'", e.OptionName, e.Got, e.Expect)
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
	return fmt.Sprintf("option '%s' value '%s' is not valid, expect one of [%s]", e.OptionName, got, expect)
}

func (e OptionValueError) Error() string {
	return e.String()
}

// --------------------

type InvalidSizeError struct {
	Size int64
	Msg  string
}

func (e InvalidSizeError) String() string {
	return fmt.Sprintf("invalid size [%s] (%d bits): %s", util.ToSizeString(e.Size), e.Size, e.Msg)
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
	return fmt.Sprintf("invalid literal '%s': %s", e.Literal, e.Err)
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
	return fmt.Sprintf("invalid enum '%s': %s", e.DefName, e.Err)
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
	return fmt.Sprintf("invalid struct '%s': %s", e.DefName, e.Err)
}

func (e InvalidStructDefError) Error() string {
	return e.String()
}

// --------------------

type InvalidEmbeddedError struct{}

func (e InvalidEmbeddedError) String() string {
	return "invalid embedded field, type must be a struct"
}

func (e InvalidEmbeddedError) Error() string {
	return e.String()
}

// --------------------

type InvalidFieldError struct {
	FieldName string
	Msg       string
}

func (e InvalidFieldError) String() string {
	return fmt.Sprintf("invalid field '%s': %s", e.FieldName, e.Msg)
}

func (e InvalidFieldError) Error() string {
	return e.String()
}
