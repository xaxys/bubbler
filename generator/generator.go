package generator

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator/genctx"
	"github.com/xaxys/bubbler/generator/target/c"
	"github.com/xaxys/bubbler/generator/target/c_minimal"
	"github.com/xaxys/bubbler/generator/target/c_minimal_single"
	"github.com/xaxys/bubbler/generator/target/c_single"
	"github.com/xaxys/bubbler/generator/target/dump"
	"github.com/xaxys/bubbler/generator/target/python"
	"github.com/xaxys/bubbler/generator/target/python_single"
	"github.com/xaxys/bubbler/util"
)

var TargetMap *util.OrderedMap[string, Generator]

func init() {
	dumpGen := dump.NewDumpGenerator()
	cGen := c.NewCGenerator()
	c_singleGen := c_single.NewCSingleGenerator()
	c_minimalGen := c_minimal.NewCMinimalGenerator()
	c_minimal_singleGen := c_minimal_single.NewCMinimalSingleGenerator()
	pythonGen := python.NewPythonGenerator()
	python_singleGen := python_single.NewPythonSingleGenerator()

	TargetMap = util.NewOrderedMap[string, Generator]()
	TargetMap.Put("dump", dumpGen)
	TargetMap.Put("c", cGen)
	TargetMap.Put("c-single", c_singleGen)
	TargetMap.Put("c_single", c_singleGen)
	TargetMap.Put("c_minimal", c_minimalGen)
	TargetMap.Put("c-minimal", c_minimalGen)
	TargetMap.Put("c_min", c_minimalGen)
	TargetMap.Put("c-min", c_minimalGen)
	TargetMap.Put("c_minimal_single", c_minimal_singleGen)
	TargetMap.Put("c-minimal-single", c_minimal_singleGen)
	TargetMap.Put("c_min_single", c_minimal_singleGen)
	TargetMap.Put("c-min-single", c_minimal_singleGen)
	TargetMap.Put("python", pythonGen)
	TargetMap.Put("py", pythonGen)
	TargetMap.Put("python-single", python_singleGen)
	TargetMap.Put("python_single", python_singleGen)
	TargetMap.Put("py-single", python_singleGen)
	TargetMap.Put("py_single", python_singleGen)
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

type Generator interface {
	Generate(ctx *genctx.GenCtx) error
}

func Generate(target string, output string, units ...*definition.CompilationUnit) error {
	gen, ok := TargetMap.Get(target)
	if !ok {
		return fmt.Errorf("target %s is not supported", target)
	}

	ctx := &genctx.GenCtx{
		Units:      units,
		OutputPath: output,
	}

	err := gen.Generate(ctx)

	if err != nil {
		return err
	}

	return nil
}
