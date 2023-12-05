package compiler

import (
	"fmt"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/parser"
	"github.com/xaxys/bubbler/util"
)

// +-------------------------------------------------+
// |                  Proto Visitor                  |
// +-------------------------------------------------+

type ProtoVisitor struct {
	parser.BasebubblerVisitor
	Unit    *definition.CompilationUnit
	Warning definition.TopLevelWarning
}

func NewParseVisitor(unit *definition.CompilationUnit) *ProtoVisitor {
	return &ProtoVisitor{
		Unit: unit,
	}
}

// ==================== Unused ====================

func (ProtoVisitor) VisitFullIdent(ctx *parser.FullIdentContext) any {
	panic("unreachable")
}

func (ProtoVisitor) VisitPackageStatement(ctx *parser.PackageStatementContext) any {
	panic("unreachable")
}

func (ProtoVisitor) VisitOptionStatement(ctx *parser.OptionStatementContext) any {
	panic("unreachable")
}

// ==================== Top Level ====================

// VisitProto returns nil
func (v *ProtoVisitor) VisitProto(ctx *parser.ProtoContext) any {
	var errs definition.TopLevelError
	for _, child := range ctx.AllTopLevelDef() {
		ret := child.Accept(v)
		switch val := ret.(type) {
		case definition.TopLevelError:
			errs = definition.TopLevelErrorsJoin(errs, val) // try parse multiple errors in one time
		case error:
			return val
		case definition.CustomType:
			v.Unit.AddLocalType(val)
		default:
			panic("unreachable")
		}
	}
	return errs
}

// VisitTopLevelDef returns CustomType or error
func (v *ProtoVisitor) VisitTopLevelDef(ctx *parser.TopLevelDefContext) any {
	if ctx.StructDef() != nil {
		return ctx.StructDef().Accept(v)
	}
	if ctx.EnumDef() != nil {
		return ctx.EnumDef().Accept(v)
	}
	panic("unreachable")
}

// ==================== Struct ====================

// VisitStructDef returns Struct or error
func (v *ProtoVisitor) VisitStructDef(ctx *parser.StructDefContext) any {
	name := ""
	nameRet := ctx.StructName().Accept(v)
	switch val := nameRet.(type) {
	case error:
		return val
	case string:
		name = val
	default:
		panic("unreachable")
	}

	namePos := &definition.BasePosition{
		File:   v.Unit.UnitName.Path,
		Line:   ctx.StructName().GetStart().GetLine(),
		Column: ctx.StructName().GetStart().GetColumn(),
	}
	dupErr := v.Unit.AddLocalName(name, namePos)
	if dupErr != nil {
		return &definition.CompileError{
			Position: namePos,
			Err:      dupErr,
		}
	}

	size := int64(0)
	if ctx.Size_() != nil {
		ret := ctx.Size_().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			size = val
		default:
			panic("unreachable")
		}
	}
	if size%8 != 0 {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Size_().GetStart().GetLine(),
				Column: ctx.Size_().GetStart().GetColumn(),
			},
			Err: &definition.InvalidSizeError{
				Size: size,
				Msg:  "struct size cannot be non-byte-aligned (bit size must be multiple of 8)",
			},
		}
	}

	body := ctx.StructBody().Accept(v)
	fields := make([]definition.Field, 0)
	switch val := body.(type) {
	case error:
		return val
	case []definition.Field:
		if len(val) == 0 {
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.StructBody().GetStart().GetLine(),
					Column: ctx.StructBody().GetStart().GetColumn(),
				},
				Err: &definition.InvalidStructDefError{
					DefName: name,
					Err:     fmt.Errorf("struct must have at least one field"),
				},
			}
		}
		fields = val
	default:
		panic("unreachable")
	}

	structDef := &definition.Struct{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		StructName:    name,
		StructBitSize: size,
		StructFields:  fields,
	}

	for _, field := range fields {
		field.SetFieldBelongs(structDef)
	}

	// TODO: check if struct normal field is byte-aligned
	if size == 0 {
		fixedSize, dynamic := structDef.SumFieldBitSize()
		if fixedSize%8 != 0 {
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.StructBody().GetStart().GetLine(),
					Column: ctx.StructBody().GetStart().GetColumn(),
				},
				Err: &definition.InvalidSizeError{
					Size: fixedSize,
					Msg:  "struct fixed-size field must be byte-aligned (bit size must be multiple of 8)",
				},
			}
		}
		if dynamic {
			structDef.StructBitSize = -1
		} else {
			structDef.StructBitSize = fixedSize
		}
	} else {
		// check struct size
		fixedSize, dynamic := structDef.SumFieldBitSize()
		pos := definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.Size_().GetStart().GetLine(),
			Column: ctx.Size_().GetStart().GetColumn(),
		}
		if dynamic {
			return &definition.CompileError{
				Position: pos,
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  "size declaration is not allowed on struct having variable-size field",
				},
			}
		}
		if fixedSize != size {
			return &definition.CompileError{
				Position: pos,
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  fmt.Sprintf("declared size does not match actual size [%s] (%d bits)", util.ToSizeString(fixedSize), fixedSize),
				},
			}
		}
	}

	return structDef
}

func (v *ProtoVisitor) VisitStructName(ctx *parser.StructNameContext) any {
	name := ctx.Ident().Accept(v).(string)
	if !util.IsCapitalized(name) {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Err: &definition.NameStyleError{
				Name: name,
				Msg:  "struct name must be Capitalized, recommended to use PascalCase",
			},
		}
	}
	standard := util.ToPascalCase(util.ToPascalCase(name))
	if name != standard {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Warning: &definition.NameStyleWarning{
				Name: name,
				Msg:  fmt.Sprintf("non-standard PascalCase detected. use '%s' instead.", standard),
			},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	}
	return name
}

func (v *ProtoVisitor) VisitStructBody(ctx *parser.StructBodyContext) any {
	var errs definition.TopLevelError
	fields := util.NewOrderedMap[string, definition.Field]()
	var addField func(f definition.Field) error
	addField = func(f definition.Field) error {
		switch val := f.(type) {
		case *definition.NormalField:
			prev, ok := fields.Get(val.FieldName)
			if ok {
				return &definition.CompileError{
					Position: f,
					Err: &definition.DefinitionDuplicateError{
						PrevDef: prev,
						DefName: val.FieldName,
					},
				}
			}
			fields.Put(val.FieldName, val)
			return nil

		case *definition.EmbeddedField:
			prev, ok := fields.Get(val.FieldType.StructName)
			if ok {
				return &definition.CompileError{
					Position: f,
					Err: &definition.DefinitionDuplicateError{
						PrevDef: prev,
						DefName: val.FieldType.StructName,
					},
				}
			}
			fields.Put(val.FieldType.StructName, val)

			// extract only when first time
			if val.FieldBelongs == nil {
				// flatten embedded struct embeddedFields to toplevel
				// add virtual 0-bit size anonymous struct field as record
				embeddedFields := val.FieldType.StructFields
				for _, field := range embeddedFields {
					// do copy
					newField := field.Copy()
					// newField.SetFieldFromEmbedded(val)
					err := addField(newField)
					if err != nil {
						return err
					}
				}
			}
			return nil
		case *definition.VoidField:
			fields.PutAnonymous(val)
			return nil
		case *definition.ConstantField:
			if val.FieldName == "" {
				fields.PutAnonymous(val)
				return nil
			}
			prev, ok := fields.Get(val.FieldName)
			if ok {
				return &definition.CompileError{
					Position: f,
					Err: &definition.DefinitionDuplicateError{
						PrevDef: prev,
						DefName: val.FieldName,
					},
				}
			}
			fields.Put(val.FieldName, val)
			return nil
		default:
			panic("unreachable")
		}
	}

	elems := ctx.AllStructElement()
	for _, elem := range elems {
		ret := elem.Accept(v)

		switch val := ret.(type) {
		case nil: // skip empty statement
			break

		case definition.TopLevelError:
			errs = definition.TopLevelErrorsJoin(errs, val)

		case error:
			return val

		case *definition.EmbeddedField:
			err := addField(val)
			if err != nil {
				return &definition.CompileError{
					Position: &definition.BasePosition{
						File:   v.Unit.UnitName.Path,
						Line:   elem.GetStart().GetLine(),
						Column: elem.GetStart().GetColumn(),
					},
					Err: &definition.EmbeddedFieldError{
						DefName: val.FieldType.StructName,
						Err:     err,
					},
				}
			}

		case definition.Field:
			err := addField(val)
			if err != nil {
				return err
			}

		default:
			panic("unreachable")
		}
	}

	if errs != nil {
		return errs
	}

	return fields.Values()
}

func (v *ProtoVisitor) VisitStructElement(ctx *parser.StructElementContext) any {
	if ctx.EmptyStatement_() != nil {
		return nil
	}
	if ctx.Field() != nil {
		return ctx.Field().Accept(v)
	}
	panic("unreachable")
}

// ==================== Field ====================

func (v *ProtoVisitor) VisitField(ctx *parser.FieldContext) any {
	if ctx.FieldVoid() != nil {
		return ctx.FieldVoid().Accept(v)
	}
	if ctx.FieldConstant() != nil {
		return ctx.FieldConstant().Accept(v)
	}
	if ctx.FieldEmbedded() != nil {
		return ctx.FieldEmbedded().Accept(v)
	}
	if ctx.FieldNormal() != nil {
		return ctx.FieldNormal().Accept(v)
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitFieldNormal(ctx *parser.FieldNormalContext) any {
	var type_ definition.Type
	type_Ret := ctx.Type_().Accept(v)
	switch val := type_Ret.(type) {
	case error:
		return val
	case definition.Type:
		type_ = val
	default:
		panic("unreachable")
	}

	var name string
	if ctx.FieldName() != nil {
		nameRet := ctx.FieldName().Accept(v)
		switch val := nameRet.(type) {
		case error:
			return val
		case string:
			name = val
		default:
			panic("unreachable")
		}
	}

	size := int64(0)
	if ctx.Size_() != nil {
		ret := ctx.Size_().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			size = val
		default:
			panic("unreachable")
		}
	}

	if size == 0 {
		size = type_.GetTypeBitSize()
	} else {
		// check field size
		pos := definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.Size_().GetStart().GetLine(),
			Column: ctx.Size_().GetStart().GetColumn(),
		}
		type_Size := type_.GetTypeBitSize()
		if type_Size == -1 {
			return &definition.CompileError{
				Position: pos,
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  "size declaration is not allowed on variable-size type",
				},
			}
		}
		if size > type_Size {
			return &definition.CompileError{
				Position: pos,
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  fmt.Sprintf("declared size cannot be greater than type size [%s] (%d bits)", util.ToSizeString(type_Size), type_Size),
				},
			}
		}
		if arrTy, ok := type_.(*definition.Array); ok {
			if size%arrTy.Length != 0 {
				return &definition.CompileError{
					Position: pos,
					Err: &definition.InvalidSizeError{
						Size: size,
						Msg:  fmt.Sprintf("declared size must be multiple of array length [%d]", arrTy.Length),
					},
				}
			}
		}
		if _, ok := type_.(*definition.Struct); ok && size != type_Size {
			return &definition.CompileError{
				Position: pos,
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  fmt.Sprintf("declared size does not match actual struct size [%s] (%d bits)", util.ToSizeString(type_Size), type_Size),
				},
			}
		}
	}

	options := util.NewOrderedMap[string, *definition.Option]()
	if ctx.FieldOptions() != nil {
		ret := ctx.FieldOptions().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case *util.OrderedMap[string, *definition.Option]:
			options = val
		default:
			panic("unreachable")
		}
	}

	fieldDef := &definition.NormalField{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		FieldName:    name,
		FieldType:    type_,
		FieldBitSize: size,
		FieldBelongs: nil,
		FieldMethods: nil,
		FieldOptions: options,
		FromEmbedded: nil,
	}

	if ctx.FieldMethods() != nil {
		basicType, ok := type_.(*definition.BasicType)
		if !ok {
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.FieldMethods().GetStart().GetLine(),
					Column: ctx.FieldMethods().GetStart().GetColumn(),
				},
				Err: &definition.InvalidFieldError{
					FieldName: name,
					Msg:       "field methods are only allowed on basic type",
				},
			}
		}
		methodVisitor := NewMethodVisitor(v, basicType)
		methods := ctx.FieldMethods().Accept(methodVisitor)
		switch val := methods.(type) {
		case error:
			return val
		case []*definition.Method:
			fieldDef.FieldMethods = val
		default:
			panic("unreachable")
		}
	}

	return fieldDef
}

func (v *ProtoVisitor) VisitFieldEmbedded(ctx *parser.FieldEmbeddedContext) any {
	var type_ *definition.Struct
	type_Ret := ctx.Type_().Accept(v)
	switch val := type_Ret.(type) {
	case error:
		return val
	case *definition.Struct:
		type_ = val
	default:
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Type_().GetStart().GetLine(),
				Column: ctx.Type_().GetStart().GetColumn(),
			},
			Err: &definition.InvalidEmbeddedFieldError{},
		}
	}

	// TODO: allow unmatched size in embedded field
	// size := int64(0)
	// if ctx.Size_() != nil {
	// 	ret := ctx.Size_().Accept(v)
	// 	switch val := ret.(type) {
	// 	case error:
	// 		return val
	// 	case int64:
	// 		size = val
	// 	default:
	// 		panic("unreachable")
	// 	}
	// }

	// if size == 0 {
	// 	size = type_.GetTypeBitSize()
	// } else {
	// 	type_Size := type_.GetTypeBitSize()
	// 	if size > type_Size {
	// 		return &definition.CompileError{
	// 			Position: definition.BasePosition{
	// 				File:   v.Unit.UnitName.Path,
	// 				Line:   ctx.Size_().GetStart().GetLine(),
	// 				Column: ctx.Size_().GetStart().GetColumn(),
	// 			},
	// 			Err: &definition.InvalidSizeError{
	// 				Size: size,
	// 				Msg:  fmt.Sprintf("declared size cannot be greater than type size [%s] (%d bits)", util.ToSizeString(type_Size), type_Size),
	// 			},
	// 		}
	// 	}
	// }

	options := util.NewOrderedMap[string, *definition.Option]()
	if ctx.FieldOptions() != nil {
		ret := ctx.FieldOptions().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case *util.OrderedMap[string, *definition.Option]:
			options = val
		default:
			panic("unreachable")
		}
	}

	fieldDef := &definition.EmbeddedField{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		FieldType:    type_,
		FieldBitSize: type_.StructBitSize,
		FieldBelongs: nil,
		FieldOptions: options,
		FromEmbedded: nil,
	}

	return fieldDef
}

func (v *ProtoVisitor) VisitFieldVoid(ctx *parser.FieldVoidContext) any {
	size := int64(0)
	sizeRet := ctx.Size_().Accept(v)
	switch val := sizeRet.(type) {
	case error:
		return val
	case int64:
		size = val
	default:
		panic("unreachable")
	}

	options := util.NewOrderedMap[string, *definition.Option]()
	if ctx.FieldOptions() != nil {
		ret := ctx.FieldOptions().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case *util.OrderedMap[string, *definition.Option]:
			options = val
		default:
			panic("unreachable")
		}
	}

	fieldDef := &definition.VoidField{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		FieldBitSize: size,
		FieldOptions: options,
		FieldBelongs: nil,
		FromEmbedded: nil,
	}

	return fieldDef
}

func (v *ProtoVisitor) VisitFieldConstant(ctx *parser.FieldConstantContext) any {
	var type_ *definition.BasicType
	type_Ret := ctx.BasicType().Accept(v)
	switch val := type_Ret.(type) {
	case error:
		return val
	case *definition.BasicType:
		type_ = val
	default:
		panic("unreachable")
	}

	var name string
	if ctx.FieldName() != nil {
		nameRet := ctx.FieldName().Accept(v)
		switch val := nameRet.(type) {
		case error:
			return val
		case string:
			name = val
		default:
			panic("unreachable")
		}
	}

	size := int64(0)
	if ctx.Size_() != nil {
		ret := ctx.Size_().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			size = val
		default:
			panic("unreachable")
		}
	}

	if size == 0 {
		size = type_.GetTypeBitSize()
	} else {
		// check field size
		type_Size := type_.GetTypeBitSize()
		if size > type_Size {
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.Size_().GetStart().GetLine(),
					Column: ctx.Size_().GetStart().GetColumn(),
				},
				Err: &definition.InvalidSizeError{
					Size: size,
					Msg:  fmt.Sprintf("declared size cannot be greater than type size [%s] (%d bits)", util.ToSizeString(type_Size), type_Size),
				},
			}
		}
	}

	var constant definition.Literal
	constantRet := ctx.Constant().Accept(v)
	switch val := constantRet.(type) {
	case error:
		return val
	case definition.Literal:
		constant = val
	default:
		panic("unreachable")
	}
	match := true
	switch constant.GetLiteralKind() {
	case definition.LiteralKindID_Bool:
		if !type_.TypeTypeID.IsBool() {
			match = false
		}
	case definition.LiteralKindID_Int:
		if !type_.TypeTypeID.IsIntOrUint() {
			match = false
		}
		if constant.GetLiteralValue().(int64) < 0 && !type_.TypeTypeID.IsInt() {
			match = false
		}
	case definition.LiteralKindID_Float:
		if !type_.TypeTypeID.IsFloat() {
			match = false
		}
	case definition.LiteralKindID_String:
		if !type_.TypeTypeID.IsString() {
			match = false
		}
	default:
		panic("unreachable")
	}
	if !match {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Constant().GetStart().GetLine(),
				Column: ctx.Constant().GetStart().GetColumn(),
			},
			Err: &definition.TypeNotMatchError{
				Type1: type_.String(),
				Type2: constant.GetLiteralKind().String(),
			},
		}
	}

	options := util.NewOrderedMap[string, *definition.Option]()
	if ctx.FieldOptions() != nil {
		ret := ctx.FieldOptions().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case *util.OrderedMap[string, *definition.Option]:
			options = val
		default:
			panic("unreachable")
		}
	}

	fieldDef := &definition.ConstantField{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		FieldName:     name,
		FieldType:     type_,
		FieldBitSize:  size,
		FieldConstant: constant,
		FieldBelongs:  nil,
		FieldOptions:  options,
		FromEmbedded:  nil,
	}

	return fieldDef
}

func (v *ProtoVisitor) VisitSize_(ctx *parser.Size_Context) any {
	bytes := int64(0)
	bits := int64(0)
	if ctx.ByteSize() != nil {
		ret := ctx.ByteSize().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			bytes = val
		default:
			panic("unreachable")
		}
	}
	if ctx.BitSize() != nil {
		ret := ctx.BitSize().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case int64:
			bits = val
		default:
			panic("unreachable")
		}
	}
	total := bytes*8 + bits
	if total == 0 {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
			},
			Err: &definition.InvalidSizeError{
				Size: total,
				Msg:  "size must be greater than 0",
			},
		}
	}
	return total
}

func (v *ProtoVisitor) VisitBitSize(ctx *parser.BitSizeContext) any {
	return ctx.IntLit().Accept(v)
}

func (v *ProtoVisitor) VisitByteSize(ctx *parser.ByteSizeContext) any {
	return ctx.IntLit().Accept(v)
}

func (v *ProtoVisitor) VisitFieldName(ctx *parser.FieldNameContext) any {
	name := ctx.Ident().Accept(v).(string)
	if !util.IsUncapitalized(name) {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Err: &definition.NameStyleError{
				Name: name,
				Msg:  "struct name must be uncapitalized, recommended to use snake_case",
			},
		}
	}
	standard := util.Tosnake_case(util.Tosnake_case(name))
	if name != standard {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Warning: &definition.NameStyleWarning{
				Name: name,
				Msg:  fmt.Sprintf("non-standard snake_case detected. use '%s' instead.", standard),
			},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	}
	return name
}

func (v *ProtoVisitor) VisitFieldOptions(ctx *parser.FieldOptionsContext) any {
	options := util.NewOrderedMap[string, *definition.Option]()
	addOption := func(o *definition.Option) error {
		prev, ok := options.Get(o.OptionName)
		if ok {
			return &definition.CompileError{
				Position: o,
				Err: &definition.OptionDuplicateError{
					PrevDef:    prev,
					OptionName: o.OptionName,
				},
			}
		}
		options.Put(o.OptionName, o)
		return nil
	}

	elems := ctx.AllFieldOption()
	for _, elem := range elems {
		ret := elem.Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case *definition.Option:
			err := FieldOptionValidator.ValidateOption(val)
			if err != nil {
				if warn, ok := err.(definition.TopLevelWarning); ok {
					v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
				} else {
					return err
				}
			}
			err = addOption(val)
			if err != nil {
				return err
			}

		default:
			panic("unreachable")
		}
	}

	return options
}

func (v *ProtoVisitor) VisitFieldOption(ctx *parser.FieldOptionContext) any {
	name := ctx.OptionName().Accept(v).(string)

	var value definition.Literal
	valueRet := ctx.Constant().Accept(v)
	switch val := valueRet.(type) {
	case error:
		return val
	case definition.Literal:
		value = val
	default:
		panic("unreachable")
	}

	option := &definition.Option{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		OptionName:  name,
		OptionValue: value,
	}

	return option
}

func (v *ProtoVisitor) VisitOptionName(ctx *parser.OptionNameContext) any {
	return ctx.Ident().Accept(v)
}

// ==================== Method ====================

func (ProtoVisitor) VisitFieldMethods(ctx *parser.FieldMethodsContext) any { panic("unreachable") }
func (ProtoVisitor) VisitFieldMethod(ctx *parser.FieldMethodContext) any   { panic("unreachable") }
func (ProtoVisitor) VisitMethodName(ctx *parser.MethodNameContext) any     { panic("unreachable") }

// +------------------------------------------------+
// |                 Method Visitor                 |
// +------------------------------------------------+

type MethodVisitor struct {
	*ProtoVisitor
	FieldType *definition.BasicType
}

func NewMethodVisitor(v *ProtoVisitor, fieldType *definition.BasicType) *MethodVisitor {
	return &MethodVisitor{
		ProtoVisitor: v,
		FieldType:    fieldType,
	}
}

func (v *MethodVisitor) VisitFieldMethods(ctx *parser.FieldMethodsContext) any {
	getters := util.NewOrderedMap[string, *definition.Method]()
	setters := util.NewOrderedMap[string, *definition.Method]()
	addGetter := func(m *definition.Method) error {
		prev, ok := getters.Get(m.MethodName)
		if ok {
			return &definition.CompileError{
				Position: m,
				Err: &definition.DefinitionDuplicateError{
					PrevDef: prev,
					DefName: m.MethodName,
				},
			}
		}
		getters.Put(m.MethodName, m)
		return nil
	}
	addSetter := func(m *definition.Method) error {
		prev, ok := setters.Get(m.MethodName)
		if ok {
			return &definition.CompileError{
				Position: m,
				Err: &definition.DefinitionDuplicateError{
					PrevDef: prev,
					DefName: m.MethodName,
				},
			}
		}
		setters.Put(m.MethodName, m)
		return nil
	}

	var errs definition.TopLevelError
	elems := ctx.AllFieldMethod()
	for _, elem := range elems {
		ret := elem.Accept(v)

		switch val := ret.(type) {
		case definition.TopLevelError:
			errs = definition.TopLevelErrorsJoin(errs, val)

		case error:
			return val

		case *definition.Method:
			switch val.MethodKind {
			case definition.MethodKindID_Get:
				err := addGetter(val)
				if err != nil {
					return err
				}
			case definition.MethodKindID_Set:
				err := addSetter(val)
				if err != nil {
					return err
				}
			default:
				panic("unreachable")
			}

		default:
			panic("unreachable")
		}
	}

	if errs != nil {
		return errs
	}

	gettersList := getters.Values()
	settersList := setters.Values()
	methods := append(gettersList, settersList...)
	return methods
}

func (v *MethodVisitor) VisitFieldMethod(ctx *parser.FieldMethodContext) any {
	var kind definition.MethodKindID
	if ctx.GET() != nil {
		kind = definition.MethodKindID_Get
	}
	if ctx.SET() != nil {
		kind = definition.MethodKindID_Set
	}

	name := ""
	if ctx.MethodName() != nil {
		nameRet := ctx.MethodName().Accept(v)
		switch val := nameRet.(type) {
		case error:
			return val
		case string:
			name = val
		default:
			panic("unreachable")
		}
	}

	var type_ *definition.BasicType
	type_Ret := ctx.BasicType().Accept(v)
	switch val := type_Ret.(type) {
	case error:
		return val
	case *definition.BasicType:
		type_ = val
	default:
		panic("unreachable")
	}

	exprVisitor := NewExprVisitor(v, type_)
	var expr definition.Expr
	exprRet := ctx.Expr().Accept(exprVisitor)
	switch val := exprRet.(type) {
	case error:
		return val
	case definition.Expr:
		expr = val
	default:
		panic("unreachable")
	}

	methodDef := &definition.Method{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		MethodKind:      kind,
		MethodName:      name,
		MethodParamType: type_,
		MethodExpr:      expr,
		MethodBelongs:   nil,
	}

	return methodDef
}

func (v *MethodVisitor) VisitMethodName(ctx *parser.MethodNameContext) any {
	name := ctx.Ident().Accept(v).(string)
	if !util.IsUncapitalized(name) {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Err: &definition.NameStyleError{
				Name: name,
				Msg:  "method name must be uncapitalized, recommended to use camelCase",
			},
		}
	}
	standard := util.TocamelCase(util.TocamelCase(name))
	if name != standard {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Warning: &definition.NameStyleWarning{
				Name: name,
				Msg:  fmt.Sprintf("non-standard camelCase detected. use '%s' instead.", standard),
			},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	}
	return name
}

// ==================== Expr ====================

func (ProtoVisitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any         { panic("unreachable") }
func (ProtoVisitor) VisitExprBitAnd(ctx *parser.ExprBitAndContext) any         { panic("unreachable") }
func (ProtoVisitor) VisitExprBitOr(ctx *parser.ExprBitOrContext) any           { panic("unreachable") }
func (ProtoVisitor) VisitExprBitXor(ctx *parser.ExprBitXorContext) any         { panic("unreachable") }
func (ProtoVisitor) VisitExprCast(ctx *parser.ExprCastContext) any             { panic("unreachable") }
func (ProtoVisitor) VisitExprConstant(ctx *parser.ExprConstantContext) any     { panic("unreachable") }
func (ProtoVisitor) VisitExprEquality(ctx *parser.ExprEqualityContext) any     { panic("unreachable") }
func (ProtoVisitor) VisitExprLogicalAnd(ctx *parser.ExprLogicalAndContext) any { panic("unreachable") }
func (ProtoVisitor) VisitExprLogicalOr(ctx *parser.ExprLogicalOrContext) any   { panic("unreachable") }
func (ProtoVisitor) VisitExprMulDivMod(ctx *parser.ExprMulDivModContext) any   { panic("unreachable") }
func (ProtoVisitor) VisitExprParens(ctx *parser.ExprParensContext) any         { panic("unreachable") }
func (ProtoVisitor) VisitExprPower(ctx *parser.ExprPowerContext) any           { panic("unreachable") }
func (ProtoVisitor) VisitExprRelational(ctx *parser.ExprRelationalContext) any { panic("unreachable") }
func (ProtoVisitor) VisitExprShift(ctx *parser.ExprShiftContext) any           { panic("unreachable") }
func (ProtoVisitor) VisitExprTernary(ctx *parser.ExprTernaryContext) any       { panic("unreachable") }
func (ProtoVisitor) VisitExprUnary(ctx *parser.ExprUnaryContext) any           { panic("unreachable") }
func (ProtoVisitor) VisitExprValue(ctx *parser.ExprValueContext) any           { panic("unreachable") }

// +------------------------------------------------+
// |                  Expr Visitor                  |
// +------------------------------------------------+

type ExprVisitor struct {
	*MethodVisitor
	ValueType *definition.BasicType
}

func NewExprVisitor(v *MethodVisitor, valueType *definition.BasicType) *ExprVisitor {
	return &ExprVisitor{
		MethodVisitor: v,
		ValueType:     valueType,
	}
}

func (v *ExprVisitor) VisitExprValue(ctx *parser.ExprValueContext) any {
	return ctx.Value().Accept(v)
}

func (v *ExprVisitor) VisitExprConstant(ctx *parser.ExprConstantContext) any {
	var constant definition.Literal
	constantRet := ctx.Constant().Accept(v)
	switch val := constantRet.(type) {
	case error:
		return val
	case definition.Literal:
		constant = val
	default:
		panic("unreachable")
	}

	expr := &definition.ConstantExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		ConstantValue: constant,
		ConstantType:  nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprParens(ctx *parser.ExprParensContext) any {
	return ctx.Expr().Accept(v)
}

func (v *ExprVisitor) VisitExprPower(ctx *parser.ExprPowerContext) any {
	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       definition.ExprOp_POW,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprUnary(ctx *parser.ExprUnaryContext) any {
	var op definition.ExprOp
	if ctx.ADD() != nil {
		op = definition.ExprOp_ADD
	}
	if ctx.SUB() != nil {
		op = definition.ExprOp_SUB
	}
	if ctx.BNOT() != nil {
		op = definition.ExprOp_BNOT
	}
	if ctx.NOT() != nil {
		op = definition.ExprOp_NOT
	}

	var oprand definition.Expr
	exprRet := ctx.Expr().Accept(v)
	switch val := exprRet.(type) {
	case error:
		return val
	case definition.Expr:
		oprand = val
	default:
		panic("unreachable")
	}

	expr := &definition.UnopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    oprand,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprCast(ctx *parser.ExprCastContext) any {
	var basicType *definition.BasicType
	basicTypeRet := ctx.BasicType().Accept(v)
	switch val := basicTypeRet.(type) {
	case error:
		return val
	case *definition.BasicType:
		basicType = val
	default:
		panic("unreachable")
	}

	var oprand definition.Expr
	exprRet := ctx.Expr().Accept(v)
	switch val := exprRet.(type) {
	case error:
		return val
	case definition.Expr:
		oprand = val
	default:
		panic("unreachable")
	}

	expr := &definition.CastExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		ToType: basicType,
		Expr1:  oprand,
	}

	return expr
}

func (v *ExprVisitor) VisitExprMulDivMod(ctx *parser.ExprMulDivModContext) any {
	var op definition.ExprOp
	if ctx.MUL() != nil {
		op = definition.ExprOp_MUL
	}
	if ctx.DIV() != nil {
		op = definition.ExprOp_DIV
	}
	if ctx.MOD() != nil {
		op = definition.ExprOp_MOD
	}

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprAddSub(ctx *parser.ExprAddSubContext) any {
	var op definition.ExprOp
	if ctx.ADD() != nil {
		op = definition.ExprOp_ADD
	}
	if ctx.SUB() != nil {
		op = definition.ExprOp_SUB
	}

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprShift(ctx *parser.ExprShiftContext) any {
	var op definition.ExprOp
	if ctx.SHL() != nil {
		op = definition.ExprOp_SHL
	}
	if ctx.SHR() != nil {
		op = definition.ExprOp_SHR
	}

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprRelational(ctx *parser.ExprRelationalContext) any {
	var op definition.ExprOp
	if ctx.LT() != nil {
		op = definition.ExprOp_LT
	}
	if ctx.GT() != nil {
		op = definition.ExprOp_GT
	}
	if ctx.LE() != nil {
		op = definition.ExprOp_LE
	}
	if ctx.GE() != nil {
		op = definition.ExprOp_GE
	}

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprEquality(ctx *parser.ExprEqualityContext) any {
	var op definition.ExprOp
	if ctx.EQ() != nil {
		op = definition.ExprOp_EQ
	}
	if ctx.NE() != nil {
		op = definition.ExprOp_NE
	}

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr

}

func (v *ExprVisitor) VisitExprBitAnd(ctx *parser.ExprBitAndContext) any {
	op := definition.ExprOp_BAND

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprBitXor(ctx *parser.ExprBitXorContext) any {
	op := definition.ExprOp_BXOR

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprBitOr(ctx *parser.ExprBitOrContext) any {
	op := definition.ExprOp_BOR

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprLogicalAnd(ctx *parser.ExprLogicalAndContext) any {
	op := definition.ExprOp_AND

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprLogicalOr(ctx *parser.ExprLogicalOrContext) any {
	op := definition.ExprOp_OR

	var left definition.Expr
	leftRet := ctx.Expr(0).Accept(v)
	switch val := leftRet.(type) {
	case error:
		return val
	case definition.Expr:
		left = val
	default:
		panic("unreachable")
	}

	var right definition.Expr
	rightRet := ctx.Expr(1).Accept(v)
	switch val := rightRet.(type) {
	case error:
		return val
	case definition.Expr:
		right = val
	default:
		panic("unreachable")
	}

	expr := &definition.BinopExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		Op:       op,
		Expr1:    left,
		Expr2:    right,
		ExprType: nil, // TODO: set type
	}

	return expr
}

func (v *ExprVisitor) VisitExprTernary(ctx *parser.ExprTernaryContext) any {
	var cond definition.Expr
	condRet := ctx.Expr(0).Accept(v)
	switch val := condRet.(type) {
	case error:
		return val
	case definition.Expr:
		cond = val
	default:
		panic("unreachable")
	}

	var trueExpr definition.Expr
	trueExprRet := ctx.Expr(1).Accept(v)
	switch val := trueExprRet.(type) {
	case error:
		return val
	case definition.Expr:
		trueExpr = val
	default:
		panic("unreachable")
	}

	var falseExpr definition.Expr
	falseExprRet := ctx.Expr(2).Accept(v)
	switch val := falseExprRet.(type) {
	case error:
		return val
	case definition.Expr:
		falseExpr = val
	default:
		panic("unreachable")
	}

	expr := &definition.TenaryExpr{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.Expr(0).GetStart().GetColumn(),
		},
		Cond:  cond,
		Expr1: trueExpr,
		Expr2: falseExpr,
	}

	return expr
}

// ==================== Type ====================

func (v *ProtoVisitor) VisitType_(ctx *parser.Type_Context) any {
	if ctx.BasicType() != nil {
		return ctx.BasicType().Accept(v)
	}
	if ctx.STRING() != nil {
		// TODO: return &definition.String{}
		// return &definition.String{}
		return &definition.GeneralError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.STRING().GetSymbol().GetLine(),
				Column: ctx.STRING().GetSymbol().GetColumn(),
			},
			Err: fmt.Errorf("string type is not supported yet"),
		}
	}
	if ctx.BYTES() != nil {
		// TODO: return &definition.Bytes{}
		// return &definition.Bytes{}
		return &definition.GeneralError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.BYTES().GetSymbol().GetLine(),
				Column: ctx.BYTES().GetSymbol().GetColumn(),
			},
			Err: fmt.Errorf("bytes type is not supported yet"),
		}
	}
	if ctx.ArrayType() != nil {
		return ctx.ArrayType().Accept(v)
	}
	if ctx.StructType() != nil {
		return ctx.StructType().Accept(v)
	}
	if ctx.EnumType() != nil {
		return ctx.EnumType().Accept(v)
	}
	if ctx.Ident() != nil {
		name := ctx.Ident().Accept(v).(string)
		if !v.Unit.GlobalTypes.Has(name) {
			pos := definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			}
			return &definition.CompileError{
				Position: pos,
				Err: &definition.DefinitionNotFoundError{
					DefName: name,
				},
			}
		}
		ty := v.Unit.GlobalTypes.MustGet(name)
		return ty
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitBasicType(ctx *parser.BasicTypeContext) any {
	if ctx.BOOL() != nil {
		return &definition.Bool
	}
	if ctx.INT8() != nil {
		return &definition.Int8
	}
	if ctx.INT16() != nil {
		return &definition.Int16
	}
	if ctx.INT32() != nil {
		return &definition.Int32
	}
	if ctx.INT64() != nil {
		return &definition.Int64
	}
	if ctx.UINT8() != nil {
		return &definition.Uint8
	}
	if ctx.UINT16() != nil {
		return &definition.Uint16
	}
	if ctx.UINT32() != nil {
		return &definition.Uint32
	}
	if ctx.UINT64() != nil {
		return &definition.Uint64
	}
	if ctx.FLOAT32() != nil {
		return &definition.Float32
	}
	if ctx.FLOAT64() != nil {
		return &definition.Float64
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitArrayType(ctx *parser.ArrayTypeContext) any {
	var base definition.Type
	baseRet := ctx.ArrayElementType().Accept(v)
	switch val := baseRet.(type) {
	case error:
		return val
	case definition.Type:
		base = val
	default:
		panic("unreachable")
	}

	length := int64(0)
	lengthRet := ctx.IntLit().Accept(v)
	switch val := lengthRet.(type) {
	case error:
		return val
	case int64:
		length = val
	default:
		panic("unreachable")
	}

	if length == 0 {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.IntLit().GetStart().GetLine(),
				Column: ctx.IntLit().GetStart().GetColumn(),
			},
			Err: &definition.InvalidSizeError{
				Size: length,
				Msg:  "array size must be greater than 0",
			},
		}
	}

	arrayDef := &definition.Array{
		ElementType: base,
		Length:      length,
	}

	return arrayDef
}

func (v *ProtoVisitor) VisitArrayElementType(ctx *parser.ArrayElementTypeContext) any {
	if ctx.BasicType() != nil {
		return ctx.BasicType().Accept(v)
	}
	if ctx.STRING() != nil {
		// TODO: return &definition.String{}
		return &definition.GeneralError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.STRING().GetSymbol().GetLine(),
				Column: ctx.STRING().GetSymbol().GetColumn(),
			},
			Err: fmt.Errorf("string type is not supported yet"),
		}
	}
	if ctx.BYTES() != nil {
		// TODO: return &definition.Bytes{}
		return &definition.GeneralError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.BYTES().GetSymbol().GetLine(),
				Column: ctx.BYTES().GetSymbol().GetColumn(),
			},
			Err: fmt.Errorf("bytes type is not supported yet"),
		}
	}
	if ctx.StructType() != nil {
		// TODO: return ctx.StructType().Accept(v)
		return &definition.GeneralError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.StructType().GetStart().GetLine(),
				Column: ctx.StructType().GetStart().GetColumn(),
			},
			Err: fmt.Errorf("struct type as array element is not supported yet"),
		}
	}
	if ctx.EnumType() != nil {
		return ctx.EnumType().Accept(v)
	}
	if ctx.Ident() != nil {
		name := ctx.Ident().Accept(v).(string)
		if !v.Unit.GlobalTypes.Has(name) {
			pos := definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			}
			return &definition.CompileError{
				Position: pos,
				Err: &definition.DefinitionNotFoundError{
					DefName: name,
				},
			}
		}

		ty := v.Unit.GlobalTypes.MustGet(name)
		switch val := ty.(type) {
		case *definition.Struct:
			// TODO: support struct as array element
			return &definition.GeneralError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.Ident().GetStart().GetLine(),
					Column: ctx.Ident().GetStart().GetColumn(),
				},
				Err: fmt.Errorf("struct type as array element is not supported yet"),
			}

		case *definition.Enum:
			return val

		default:
			panic("unreachable")
		}
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitStructType(ctx *parser.StructTypeContext) any {
	if ctx.StructDef() != nil {
		return ctx.StructDef().Accept(v)
	}
	if ctx.StructName() != nil {
		name := ctx.StructName().Accept(v).(string)
		if !v.Unit.GlobalTypes.Has(name) {
			pos := definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.StructName().GetStart().GetLine(),
				Column: ctx.StructName().GetStart().GetColumn(),
			}
			return &definition.CompileError{
				Position: pos,
				Err: &definition.DefinitionNotFoundError{
					DefName: name,
				},
			}
		}
		ty := v.Unit.GlobalTypes.MustGet(name)
		switch val := ty.(type) {
		case *definition.Struct:
			return val
		default:
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.StructName().GetStart().GetLine(),
					Column: ctx.StructName().GetStart().GetColumn(),
				},
				Err: &definition.DefinitionTypeConflictError{
					DefName: name,
					Expect:  val.GetTypeName(),
					Got:     "struct",
				},
			}
		}
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitEnumType(ctx *parser.EnumTypeContext) any {
	if ctx.EnumName() != nil {
		name := ctx.EnumName().Accept(v).(string)
		if !v.Unit.GlobalTypes.Has(name) {
			pos := definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.EnumName().GetStart().GetLine(),
				Column: ctx.EnumName().GetStart().GetColumn(),
			}
			return &definition.CompileError{
				Position: pos,
				Err: &definition.DefinitionNotFoundError{
					DefName: name,
				},
			}
		}
		ty := v.Unit.GlobalTypes.MustGet(name)
		switch val := ty.(type) {
		case *definition.Enum:
			return val
		default:
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.EnumName().GetStart().GetLine(),
					Column: ctx.EnumName().GetStart().GetColumn(),
				},
				Err: &definition.DefinitionTypeConflictError{
					DefName: name,
					Expect:  val.GetTypeName(),
					Got:     "enum",
				},
			}
		}
	}
	panic("unreachable")
}

// ==================== Enum ====================

func (v *ProtoVisitor) VisitEnumDef(ctx *parser.EnumDefContext) any {
	name := ""
	nameRet := ctx.EnumName().Accept(v)
	switch val := nameRet.(type) {
	case error:
		return val
	case string:
		name = val
	default:
		panic("unreachable")
	}

	namePos := &definition.BasePosition{
		File:   v.Unit.UnitName.Path,
		Line:   ctx.EnumName().GetStart().GetLine(),
		Column: ctx.EnumName().GetStart().GetColumn(),
	}
	dupErr := v.Unit.AddLocalName(name, namePos)
	if dupErr != nil {
		return &definition.CompileError{
			Position: namePos,
			Err:      dupErr,
		}
	}

	size := int64(0)
	sizeRet := ctx.Size_().Accept(v)
	switch val := sizeRet.(type) {
	case error:
		return val
	case int64:
		size = val
	default:
		panic("unreachable")
	}

	body := ctx.EnumBody().Accept(v)
	var values *util.OrderedMap[string, *definition.EnumValue]
	switch val := body.(type) {
	case error:
		return val
	case *util.OrderedMap[string, *definition.EnumValue]:
		if val.Len() == 0 {
			return &definition.CompileError{
				Position: definition.BasePosition{
					File:   v.Unit.UnitName.Path,
					Line:   ctx.EnumBody().GetStart().GetLine(),
					Column: ctx.EnumBody().GetStart().GetColumn(),
				},
				Err: &definition.InvalidEnumDefError{
					DefName: name,
					Err:     fmt.Errorf("enum must have at least one value"),
				},
			}
		}
		values = val
	default:
		panic("unreachable")
	}

	var errs definition.TopLevelError
	for _, val := range values.Values() {
		dupErr := v.Unit.AddLocalName(val.EnumValueName, val)
		if dupErr != nil {
			ex := &definition.CompileError{
				Position: val,
				Err:      dupErr,
			}
			errs = definition.TopLevelErrorsJoin(errs, ex)
		}
	}

	if errs != nil {
		return errs
	}

	enumDef := &definition.Enum{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		EnumName:    name,
		EnumBitSize: size,
		EnumValues:  values,
	}

	for _, val := range values.Values() {
		val.EnumBelongs = enumDef
	}

	return enumDef
}

func (v *ProtoVisitor) VisitEnumBody(ctx *parser.EnumBodyContext) any {
	values := util.NewOrderedMap[string, *definition.EnumValue]()

	var addValue func(val *definition.EnumValue) error
	addValue = func(val *definition.EnumValue) error {
		prev, ok := values.Get(val.EnumValueName)
		if ok {
			return &definition.CompileError{
				Position: val,
				Err: &definition.DefinitionDuplicateError{
					PrevDef: prev,
					DefName: val.EnumValueName,
				},
			}
		}
		if val.EnumValue == -1 {
			newNum := int64(0)
			if values.Len() > 0 {
				newNum = values.Last().Value.EnumValue + 1
			}
			val.EnumValue = newNum
		}
		values.Put(val.EnumValueName, val)
		return nil
	}

	var errs definition.TopLevelError
	for _, elem := range ctx.AllEnumElement() {
		ret := elem.Accept(v)

		switch val := ret.(type) {
		case definition.TopLevelError:
			errs = definition.TopLevelErrorsJoin(errs, val)

		case nil: // skip empty statement
			break

		case error:
			return val

		case *definition.EnumValue:
			err := addValue(val)
			if err != nil {
				return err
			}

		default:
			panic("unreachable")

		}
	}

	if errs != nil {
		return errs
	}

	return values
}

func (v *ProtoVisitor) VisitEnumElement(ctx *parser.EnumElementContext) any {
	if ctx.EmptyStatement_() != nil {
		return nil
	}
	if ctx.EnumValue() != nil {
		return ctx.EnumValue().Accept(v)
	}
	panic("unreachable")
}

func (v *ProtoVisitor) VisitEnumValue(ctx *parser.EnumValueContext) any {
	var name string
	nameRet := ctx.EnumValueName().Accept(v)
	switch val := nameRet.(type) {
	case error:
		return val
	case string:
		name = val
	default:
		panic("unreachable")
	}

	value := int64(-1)
	if ctx.Constant() != nil {
		ret := ctx.Constant().Accept(v)
		switch val := ret.(type) {
		case error:
			return val
		case definition.Literal:
			switch val := val.(type) {
			case *definition.IntLiteral:
				value = val.IntValue
			default:
				return &definition.CompileError{
					Position: definition.BasePosition{
						File:   v.Unit.UnitName.Path,
						Line:   ctx.Constant().GetStart().GetLine(),
						Column: ctx.Constant().GetStart().GetColumn(),
					},
					Err: &definition.EnumConstValueTypeError{
						Constant: fmt.Sprint(val),
						Got:      val.GetLiteralKind().String(),
					},
				}
			}

		default:
			panic("unreachable")
		}
	}

	// TODO: parse enum options
	// options := ctx.EnumValueOptions().Accept(v)

	fieldDef := &definition.EnumValue{
		BasePosition: definition.BasePosition{
			File:   v.Unit.UnitName.Path,
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
		},
		EnumValueName: name,
		EnumValue:     value,
		EnumBelongs:   nil,
	}

	return fieldDef
}

func (v *ProtoVisitor) VisitEnumName(ctx *parser.EnumNameContext) any {
	name := ctx.Ident().Accept(v).(string)
	if !util.IsCapitalized(name) {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Err: &definition.NameStyleError{
				Name: name,
				Msg:  "enum name must be Capitalized, recommended to use PascalCase",
			},
		}
	}
	standard := util.ToPascalCase(util.ToPascalCase(name))
	if name != standard {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Warning: &definition.NameStyleWarning{
				Name: name,
				Msg:  fmt.Sprintf("non-standard PascalCase detected. use '%s' instead.", standard),
			},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	}
	return name
}

func (v *ProtoVisitor) VisitEnumValueName(ctx *parser.EnumValueNameContext) any {
	name := ctx.Ident().Accept(v).(string)
	if !util.IsCapitalized(name) {
		return &definition.CompileError{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Err: &definition.NameStyleError{
				Name: name,
				Msg:  "enum value name must be Capitalized, recommended to use ALLCAP_CASE",
			},
		}
	}
	standard := util.ToALLCAP_CASE(util.ToALLCAP_CASE(name))
	if name != standard {
		warn := &definition.CompileWarning{
			Position: definition.BasePosition{
				File:   v.Unit.UnitName.Path,
				Line:   ctx.Ident().GetStart().GetLine(),
				Column: ctx.Ident().GetStart().GetColumn(),
			},
			Warning: &definition.NameStyleWarning{
				Name: name,
				Msg:  fmt.Sprintf("non-standard ALLCAP_CASE detected. use '%s' instead.", standard),
			},
		}
		v.Warning = definition.TopLevelWarningsJoin(v.Warning, warn)
	}
	return name
}

// TODO: parse enum options
func (v *ProtoVisitor) VisitEnumValueOption(ctx *parser.EnumValueOptionContext) any {
	return nil
}

// TODO: parse enum options
func (v *ProtoVisitor) VisitEnumValueOptions(ctx *parser.EnumValueOptionsContext) any {
	return nil
}

// ==================== Literal ====================

// VisitConstant returns Literal or error
func (v *ProtoVisitor) VisitConstant(ctx *parser.ConstantContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitIdent noexcept returns string
func (v *ProtoVisitor) VisitIdent(ctx *parser.IdentContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitIntLit returns int64 or error
func (v *ProtoVisitor) VisitIntLit(ctx *parser.IntLitContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitFloatLit returns float64 or error
func (v *ProtoVisitor) VisitFloatLit(ctx *parser.FloatLitContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitBoolLit noexcept returns bool
func (v *ProtoVisitor) VisitBoolLit(ctx *parser.BoolLitContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitStrLit returns string or error
func (v *ProtoVisitor) VisitStrLit(ctx *parser.StrLitContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// VisitValue noexcept returns *definition.ValueExpr
func (v *ProtoVisitor) VisitValue(ctx *parser.ValueContext) any {
	lv := NewLiteralVisitor(v.Unit)
	return ctx.Accept(lv)
}

// ==================== Unimplemented ====================

func (v *ProtoVisitor) VisitImportStatement(ctx *parser.ImportStatementContext) any {
	return nil
}

func (v *ProtoVisitor) VisitEmptyStatement_(ctx *parser.EmptyStatement_Context) any {
	return nil
}

// ==================== Util ====================

// TypePromotion input basic type only
// func TypePromotion(ty1, ty2 definition.TypeID) (ret definition.TypeID, warning error, err error) {
// 	if !ty1.IsBasic() || !ty2.IsBasic() {
// 		panic("unexpected usage")
// 	}
// 	if ty1 == ty2 {
// 		return ty1, nil, nil
// 	}
// 	if ty1.IsBool() && ty2.IsBool() {
// 		return definition.TypeID_Bool, nil, nil
// 	}
// 	if ty1.IsBool() || ty2.IsBool() {
// 		return -1, nil, &definition.TypeNotMatchError{
// 			Type1: ty1.String(),
// 			Type2: ty2.String(),
// 		}
// 	}
// 	if ty1.IsFloat() && ty2.IsFloat() {
// 		ty := max(ty1, ty2)
// 		return ty, nil, nil
// 	}
// 	if ty1.IsFloat() || ty2.IsFloat() {
// 		ty := max(ty1, ty2)
// 		return ty, nil, nil
// 	}
// 	if ty1.IsInt() && ty2.IsInt() {
// 		ty := max(ty1, ty2)
// 		return ty, nil, nil
// 	}

// }
