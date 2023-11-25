package definition

import (
	"errors"
	"fmt"

	"github.com/xaxys/bubbler/util"
)

type CompilationUnit struct {
	UnitName *FileIdentifer
	Imports  *util.OrderedMap[string, *CompilationUnit]
	Types    *util.OrderedMap[string, CustomType]
}

func NewCompilationUnit(name *FileIdentifer) *CompilationUnit {
	return &CompilationUnit{
		UnitName: name,
		Imports:  util.NewOrderedMap[string, *CompilationUnit](),
		Types:    util.NewOrderedMap[string, CustomType](),
	}
}

func (c *CompilationUnit) AddImport(other *CompilationUnit) error {
	var err error
	if other.UnitName == c.UnitName {
		return &ImportSelfError{}
	}
	if other.Imports.Has(c.UnitName.Path) {
		return &ImportCycleError{File: c.UnitName}
	}
	if c.Imports.Has(other.UnitName.Path) {
		// skip diamond import for those already imported
		return nil
	}
	for _, ty := range other.Types.Values() {
		pos, ok := ty.(Position)
		if !ok {
			ex := fmt.Errorf("unexpected top level type %s without its position", ty.GetTypeName())
			err = errors.Join(err, ex)
			continue
		}
		success := c.Types.Put(ty.GetTypeName(), ty)
		if !success {
			prev := c.Types.MustGet(ty.GetTypeName())
			ex := &CompileError{
				Position: pos,
				Err: &DefinitionDuplicateError{
					PrevDef: prev.(Position),
					DefName: ty.GetTypeName(),
				},
			}
			err = errors.Join(err, ex)
			continue
		}
	}
	if err != nil {
		err := &ImportingError{
			File: c.UnitName,
			Err:  err,
		}
		return err
	}
	c.Imports.Put(other.UnitName.Path, other)
	return nil
}

func (c CompilationUnit) String() string {
	imports := "[\n"
	for _, unit := range c.Imports.Values() {
		imports += util.IndentSpace8(unit.UnitName.Path) + "\n"
	}
	imports += "    ]"

	types := "[\n"
	for _, ty := range c.Types.Values() {
		types += util.IndentSpace8(ty) + "\n"
	}
	types += "    ]"

	return fmt.Sprintf("CompilationUnit {\n    UnitName: %s\n    Imports: %s\n    Types: %s\n}", c.UnitName, imports, types)
}
