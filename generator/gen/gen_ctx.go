package gen

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/fileio"
)

type GenCtx struct {
	Units      []*definition.CompilationUnit
	GenOptions *GenOptions
	OutputPath string
}

func normalizeGeneratedIndent(text string) string {
	// Enforce space-based indentation in generated outputs.
	return strings.ReplaceAll(text, "\t", "    ")
}

func (ctx *GenCtx) WritePackage(pkg *definition.Package, ext string, text string) error {
	text = normalizeGeneratedIndent(text)
	if ctx.OutputPath == "" {
		fmt.Println(text)
		return nil
	}

	file := pkg.ToFilePath(ext)
	return ctx.WriteFile(file, text)
}

func (ctx *GenCtx) WriteFile(file string, text string) error {
	text = normalizeGeneratedIndent(text)
	for _, rmpath := range ctx.GenOptions.RemovePath {
		if strings.HasPrefix(file, rmpath) {
			file = file[len(rmpath):]
			for strings.HasPrefix(file, "/") {
				file = file[1:]
			}
			break
		}
	}
	if ctx.OutputPath == "" {
		fmt.Println(text)
		return nil
	}
	path := filepath.Join(ctx.OutputPath, file)
	return ctx.WriteFileAbs(path, text)
}

func (ctx *GenCtx) WriteFileAbs(path string, text string) error {
	text = normalizeGeneratedIndent(text)
	if path == "" {
		fmt.Println(text)
		return nil
	}

	id, err := fileio.GetFileIdentifer(path)
	if err != nil {
		return err
	}

	err = fileio.WriteFileContent(id, text)
	if err != nil {
		return err
	}

	return nil
}
