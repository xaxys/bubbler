package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xaxys/bubbler/compiler"
	"github.com/xaxys/bubbler/definition"
	"github.com/xaxys/bubbler/generator"
	"github.com/xaxys/bubbler/generator/gen"
	"github.com/xaxys/bubbler/util"
)

var (
	BuildTags = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
	GoVersion = "unknown"
)

// Font: slant
// http://www.network-science.de/ascii/
const bannerTemplate = `
{{- define "banner" }}\
    __          __    __    __       
   / /_  __  __/ /_  / /_  / /__  ___
  / __ \/ / / / __ \/ __ \/ / _ \/ _/
 / /_/ / /_/ / /_/ / /_/ / /  __/ /  
/_.___/\__,_/_.___/_.___/_/\___/_/   

Welcome to use bubbler!
Version:   {{ .BuildTags }}
Built:     {{ .BuildTime }}
GitCommit: {{ .GitCommit }}
GoVersion: {{ .GoVersion }}

Usage:
  bubbler [options] <input file>

Options:
  -t <target>        Target Language
  -o <output>        Output Path
  -inner             Generate Inner Class
  -single            Generate Single File
  -minimal           Generate Minimal Code
  -decnum            Force Generate Decimal Format for Constant Value
  -memcpy            Allocate Memory and Copy Data for Variable-Size Type
  -signext <method>  Sign Extension Method (shift, arith)

Targets:
{{ range .Generators }}  {{ . }}
{{ end }}
Examples:
  bubbler -t c -minimal -o output/ example.bb
  bubbler -t c -single -o gen.hpp example.bb
  bubbler -t py -decnum -signext arith -o output example.bb

For more information, please visit: https://github.com/xaxys/bubbler
{{ end }}
`

func printBanner() {
	data := map[string]interface{}{
		"BuildTags":  BuildTags,
		"BuildTime":  BuildTime,
		"GitCommit":  GitCommit,
		"GoVersion":  GoVersion,
		"Generators": generator.ListGenerators(),
	}
	tmpl := util.ExecuteTemplate(bannerTemplate, "banner", nil, data)
	fmt.Println(tmpl)
}

func main() {
	if len(os.Args) == 1 {
		printBanner()
		os.Exit(0)
	}

	target := ""
	output := ""
	inner := false
	single := false
	minimal := false
	decnum := false
	memcpy := false
	signext := ""
	flag.StringVar(&target, "t", "", "Target Language")
	flag.StringVar(&output, "o", "", "Output Path")
	flag.BoolVar(&inner, "inner", false, "Generate Inner Class")
	flag.BoolVar(&single, "single", false, "Generate Single File")
	flag.BoolVar(&minimal, "minimal", false, "Generate Minimal Code")
	flag.BoolVar(&decnum, "decnum", false, "Force Generate Decimal Format for Constant Value")
	flag.BoolVar(&memcpy, "memcpy", false, "Allocate Memory and Copy Data for Variable-Size Type")
	flag.StringVar(&signext, "signext", "", "Sign Extension Method (shift, arith)")
	flag.Parse()

	// check input file
	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, &definition.GeneralError{
			Err: &definition.NoInputFileError{},
		})
		os.Exit(1)
	}
	if len(files) > 1 {
		fmt.Fprintln(os.Stderr, &definition.GeneralError{
			Err: &definition.MultipleInputFileError{
				Files: files,
			},
		})
		os.Exit(1)
	}

	// check target before compiling
	gentor, err := generator.GetGenerator(target)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// check option before compiling
	signextOpt, err := gen.SignExtMethod(signext)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// compile
	input := files[0]
	units, err, warning := compiler.Compile(input)
	if warning != nil {
		fmt.Fprintln(os.Stderr, warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// generate
	options := gen.NewGenOptions(
		gen.SingleFile(single),
		gen.InnerClass(inner),
		gen.MinimalCode(minimal),
		gen.DecimalNumber(decnum),
		gen.MemoryCopy(memcpy),
		signextOpt,
	)
	ctx := &gen.GenCtx{
		Units:      units,
		GenOptions: options,
		OutputPath: output,
	}
	err, warning = gentor.Generate(ctx)
	if warning != nil {
		fmt.Fprintln(os.Stderr, warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
