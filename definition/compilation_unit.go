package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

type CompilationUnit struct {
	UnitName *FileIdentifer
	Imports  *util.OrderedMap[string, *CompilationUnit]
	Types    *util.OrderedMap[string, CustomType] // Types contains all the top-level typedefs (inner typedefs excluded)
	Names    *util.OrderedMap[string, Position]   // Names contains all the names (inner typedefs' name included)
}

func NewCompilationUnit(name *FileIdentifer) *CompilationUnit {
	return &CompilationUnit{
		UnitName: name,
		Imports:  util.NewOrderedMap[string, *CompilationUnit](),
		Types:    util.NewOrderedMap[string, CustomType](),
		Names:    util.NewOrderedMap[string, Position](),
	}
}

func (c *CompilationUnit) AddImport(other *CompilationUnit) error {
	var err TopLevelError
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
			ex := &CompileError{
				Position: &BasePosition{
					File:   other.UnitName.Path,
					Line:   0,
					Column: 0,
				},
				Err: fmt.Errorf("unexpected top level type %s without its position", ty.GetTypeName()),
			}
			err = TopLevelErrorsJoin(err, ex)
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
			err = TopLevelErrorsJoin(err, ex)
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

	names := "[\n"
	for _, name := range c.Names.Keys() {
		names += util.IndentSpace8(name) + "\n"
	}
	names += "    ]"

	return fmt.Sprintf("CompilationUnit {\n    UnitName: %s\n    Imports: %s\n    Types: %s\n	Names: %s\n}", c.UnitName, imports, types, names)
}
