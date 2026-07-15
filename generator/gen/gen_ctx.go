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
	// Keep generated files deterministic and editor-friendly without relying
	// on a language formatter after generation.
	text = strings.ReplaceAll(text, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\t", "    ")

	lines := strings.Split(text, "\n")
	normalized := make([]string, 0, len(lines))
	blankLines := 0
	for _, line := range lines {
		line = strings.TrimRight(line, " \t")
		if line == "" {
			blankLines++
			if blankLines > 2 {
				continue
			}
		} else {
			blankLines = 0
		}
		normalized = append(normalized, line)
	}

	return strings.TrimRight(strings.Join(normalized, "\n"), "\n") + "\n"
}

func (ctx *GenCtx) WritePackage(pkg *definition.Package, ext string, text string) error {
	text = normalizeGeneratedIndent(text)
	if ctx.OutputPath == "" {
		fmt.Print(text)
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
		fmt.Print(text)
		return nil
	}
	path := filepath.Join(ctx.OutputPath, file)
	return ctx.WriteFileAbs(path, text)
}

func (ctx *GenCtx) WriteFileAbs(path string, text string) error {
	text = normalizeGeneratedIndent(text)
	if path == "" {
		fmt.Print(text)
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
