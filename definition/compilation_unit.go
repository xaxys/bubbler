package definition

import (
	"fmt"

	"github.com/xaxys/bubbler/util"
)

type CompilationUnit struct {
	UnitName *FileIdentifer
	Package  *Package
	Options  *util.OrderedMap[string, *Option]

	GlobalImports *util.OrderedMap[string, *CompilationUnit] // GlobalImports contains all the imports, including indirect imports
	LocalImports  *util.OrderedMap[string, *CompilationUnit] // LocalImports contains direct imports of this unit

	// GlobalTypes and GlobalNames include imported types, representing all the defs visible to this unit
	GlobalTypes *util.OrderedMap[string, CustomType] // GlobalTypes contains all the top-level typedefs (inner typedefs excluded)
	GlobalNames *util.OrderedMap[string, Position]   // GlobalNames contains all the names (inner typedefs' name included)

	// LocalTypes and LocalNames are only contained types and names defined in this unit
	LocalTypes *util.OrderedMap[string, CustomType] // LocalTypes contains all the top-level typedefs (inner typedefs excluded)
	LocalNames *util.OrderedMap[string, Position]   // LocalNames contains all the local names (inner typedefs' name included)
}

func NewCompilationUnit(name *FileIdentifer) *CompilationUnit {
	return &CompilationUnit{
		UnitName:      name,
		Package:       nil,
		Options:       util.NewOrderedMap[string, *Option](),
		GlobalImports: util.NewOrderedMap[string, *CompilationUnit](),
		LocalImports:  util.NewOrderedMap[string, *CompilationUnit](),
		GlobalTypes:   util.NewOrderedMap[string, CustomType](),
		LocalTypes:    util.NewOrderedMap[string, CustomType](),
		GlobalNames:   util.NewOrderedMap[string, Position](),
		LocalNames:    util.NewOrderedMap[string, Position](),
	}
}

// AddLocalType adds a typedef to this unit
//
// returns DefinitionDuplicateError if the type name is already defined in global scope
//
// There's no need to call AddLocalName again for the same type
func (c *CompilationUnit) AddLocalType(ty CustomType) error {
	if c.GlobalNames.Has(ty.GetTypeName()) {
		prev := c.GlobalNames.MustGet(ty.GetTypeName())
		if !prev.Eqauals(ty) {
			return &DefinitionDuplicateError{
				PrevDef: prev,
				DefName: ty.GetTypeName(),
			}
		}
	}
	c.GlobalTypes.Put(ty.GetTypeName(), ty)
	c.GlobalNames.Put(ty.GetTypeName(), ty)
	c.LocalTypes.Put(ty.GetTypeName(), ty)
	c.LocalNames.Put(ty.GetTypeName(), ty)
	return nil
}

// AddLocalName adds a name to this unit (for embedded structs or enum values)
//
// returns DefinitionDuplicateError if the name is already defined in global scope
func (c *CompilationUnit) AddLocalName(name string, pos Position) error {
	if c.GlobalNames.Has(name) {
		prev := c.GlobalNames.MustGet(name)
		if !prev.Eqauals(pos) {
			return &DefinitionDuplicateError{
				PrevDef: prev,
				DefName: name,
			}
		}
	}
	c.GlobalNames.Put(name, pos)
	c.LocalNames.Put(name, pos)
	return nil
}

// AddGlobalType adds a typedef to this unit
//
// returns DefinitionDuplicateError if the type name is already defined in global scope
//
// returns nil if found the definition of same position
//
// There's no need to call AddGlobalName again for the same type
func (c *CompilationUnit) AddGlobalType(ty CustomType) error {
	if c.GlobalNames.Has(ty.GetTypeName()) {
		prev := c.GlobalNames.MustGet(ty.GetTypeName())
		if !prev.Eqauals(ty) {
			return &DefinitionDuplicateError{
				PrevDef: prev,
				DefName: ty.GetTypeName(),
			}
		}
	}
	c.GlobalTypes.Put(ty.GetTypeName(), ty)
	c.GlobalNames.Put(ty.GetTypeName(), ty)
	return nil
}

// AddGlobalName adds a name to this unit
//
// returns DefinitionDuplicateError if the name is already defined in global scope
//
// returns nil if found the name of same position
func (c *CompilationUnit) AddGlobalName(name string, pos Position) error {
	if c.GlobalNames.Has(name) {
		prev := c.GlobalNames.MustGet(name)
		if !prev.Eqauals(pos) {
			return &DefinitionDuplicateError{
				PrevDef: prev,
				DefName: name,
			}
		}
	}
	c.GlobalNames.Put(name, pos)
	return nil
}

// AddImport adds a compilation unit to this unit
func (c *CompilationUnit) AddImport(other *CompilationUnit) error {
	var err TopLevelError
	// already checked in precheckFilePath
	// if other.UnitName == c.UnitName {
	// 	return &ImportSelfError{}
	// }
	// if other.Imports.Has(c.UnitName.Path) {
	// 	return &ImportCycleError{File: c.UnitName}
	// }
	if c.GlobalImports.Has(other.UnitName.Path) {
		// skip diamond import for those already imported
		c.LocalImports.Put(other.UnitName.Path, other)
		return nil
	}

	for _, entry := range other.GlobalNames.Entries() {
		name := entry.Key
		pos := entry.Value
		dupErr := c.AddGlobalName(name, pos)
		if dupErr != nil {
			ex := &CompileError{
				Position: pos,
				Err:      dupErr,
			}
			err = TopLevelErrorsJoin(err, ex)
			continue
		}
	}
	for _, ty := range other.GlobalTypes.Values() {
		dupErr := c.AddGlobalType(ty)
		if dupErr != nil {
			ex := &CompileError{
				Position: ty,
				Err:      dupErr,
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

	// package name conflict already checked in precheckPacakgeName

	c.LocalImports.Put(other.UnitName.Path, other)
	c.GlobalImports.Put(other.UnitName.Path, other)
	for _, impo := range other.GlobalImports.Values() {
		c.GlobalImports.Put(impo.UnitName.Path, impo)
	}
	return nil
}

func (c CompilationUnit) String() string {
	options := "[\n"
	for _, option := range c.Options.Values() {
		options += util.IndentSpace8(option) + "\n"
	}
	options += "    ]"

	imports := "[\n"
	for _, unit := range c.GlobalImports.Values() {
		imports += util.IndentSpace8(unit.UnitName.Path) + "\n"
	}
	imports += "    ]"

	localImports := "[\n"
	for _, unit := range c.LocalImports.Values() {
		localImports += util.IndentSpace8(unit.UnitName.Path) + "\n"
	}
	localImports += "    ]"

	types := "[\n"
	for _, ty := range c.GlobalTypes.Values() {
		types += util.IndentSpace8(ty.GetTypeName()) + "\n"
	}
	types += "    ]"

	names := "[\n"
	for _, name := range c.GlobalNames.Keys() {
		names += util.IndentSpace8(name) + "\n"
	}
	names += "    ]"

	localTypes := "[\n"
	for _, ty := range c.LocalTypes.Values() {
		localTypes += util.IndentSpace8(ty.GetTypeName()) + "\n"
	}
	localTypes += "    ]"

	localNames := "[\n"
	for _, name := range c.LocalNames.Keys() {
		localNames += util.IndentSpace8(name) + "\n"
	}
	localNames += "    ]"

	return fmt.Sprintf("CompilationUnit {\n    UnitName: %s\n    Package: %s\n    Options: %s\n    Imports: %s\n    LocalImports: %s\n    Types: %s\n    Names: %s\n    LocalTypes: %s\n    LocalNames: %s\n}",
		c.UnitName, c.Package, options, imports, localImports, types, names, localTypes, localNames)
}
