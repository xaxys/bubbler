package generator

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/generator/target/c"
	"github.com/xaxys/bubbler/generator/target/commonjs"
	"github.com/xaxys/bubbler/generator/target/cpp"
	"github.com/xaxys/bubbler/generator/target/csharp"
	"github.com/xaxys/bubbler/generator/target/golang"
	"github.com/xaxys/bubbler/generator/target/java"
	"github.com/xaxys/bubbler/generator/target/python"
	"github.com/xaxys/bubbler/util"
)

type Generator interface {
	Generate(ctx *gen.GenCtx) (err error, warning error)
}

var TargetMap *util.OrderedMap[string, Generator]

func init() {
	// dumpGen := dump.NewDumpGenerator()
	cGen := c.NewCGenerator()
	cppGen := cpp.NewCppGenerator()
	csharpGen := csharp.NewCSharpGenerator()
	commonjsGen := commonjs.NewCommonJSGenerator()
	goGen := golang.NewGoGenerator()
	javaGen := java.NewJavaGenerator()
	pythonGen := python.NewPythonGenerator()

	TargetMap = util.NewOrderedMap[string, Generator]()
	// TargetMap.Put("dump", dumpGen)
	TargetMap.Put("c", cGen)
	TargetMap.Put("cpp", cppGen)
	TargetMap.Put("csharp", csharpGen)
	TargetMap.Put("cs", csharpGen)
	TargetMap.Put("commonjs", commonjsGen)
	TargetMap.Put("cjs", commonjsGen)
	TargetMap.Put("go", goGen)
	TargetMap.Put("java", javaGen)
	TargetMap.Put("python", pythonGen)
	TargetMap.Put("py", pythonGen)
}

func ListGenerators() []string {
	var ret []string
	var last Generator
	var aliases []string
	for _, entry := range TargetMap.Entries() {
		key := entry.Key
		gen := entry.Value
		if gen == last {
			aliases = append(aliases, key)
		} else {
			str := key
			if aliases != nil {
				ret[len(ret)-1] += fmt.Sprintf(" [%s]", strings.Join(aliases, ", "))
				aliases = nil
			}
			ret = append(ret, str)
		}
		last = gen
	}
	if aliases != nil {
		ret[len(ret)-1] += fmt.Sprintf(" [%s]", strings.Join(aliases, ", "))
	}
	return ret
}

func GetGenerator(target string) (Generator, error) {
	if target == "" {
		return nil, &definition.GeneralError{
			Err: &definition.TargetNotSpecifiedError{},
		}
	}
	generator, ok := TargetMap.Get(target)
	if !ok {
		return nil, &definition.GeneralError{
			Err: &definition.TargetNotSupportedError{
				Target: target,
			},
		}
	}
	return generator, nil
}

func Generate(target string, output string, options *gen.GenOptions, units ...*definition.CompilationUnit) (retErr error, retWarning error) {
	generator, err := GetGenerator(target)
	if err != nil {
		return err, nil
	}

	ctx := &gen.GenCtx{
		Units:      units,
		GenOptions: options,
		OutputPath: output,
	}

	err, warning := generator.Generate(ctx)

	return err, warning
}
