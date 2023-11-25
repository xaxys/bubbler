package generator

import (
	"fmt"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/fileio"
	"github.com/xaxys/bubbler/generator/c"
	"github.com/xaxys/bubbler/generator/c_minimal"
	"github.com/xaxys/bubbler/generator/dump"
	"github.com/xaxys/bubbler/util"
)

var TargetMap *util.OrderedMap[string, Generator]

func init() {
	dumpGen := dump.NewDumpGenerator()
	cGen := c.NewCGenerator()
	c_minimalGen := c_minimal.NewCMinimalGenerator()

	TargetMap = util.NewOrderedMap[string, Generator]()
	TargetMap.Put("dump", dumpGen)
	TargetMap.Put("c", cGen)
	TargetMap.Put("c_minimal", c_minimalGen)
	TargetMap.Put("c-minimal", c_minimalGen)
	TargetMap.Put("c_min", c_minimalGen)
	TargetMap.Put("c-min", c_minimalGen)
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
	Generate(unit *definition.CompilationUnit) (string, error)
}

func Generate(target string, file string, unit *definition.CompilationUnit) error {
	gen, ok := TargetMap.Get(target)
	if !ok {
		return fmt.Errorf("target %s is not supported", target)
	}

	text, err := gen.Generate(unit)
	if err != nil {
		return err
	}

	if file == "" {
		fmt.Println(text)
		return nil
	}

	id, err := fileio.GetFileIdentifer(file)
	if err != nil {
		return err
	}

	err = fileio.WriteFileContent(id, text)
	if err != nil {
		return err
	}

	return nil
}
