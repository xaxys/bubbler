package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/xaxys/bubbler/compiler"
	"github.com/xaxys/bubbler/generator"
)

var (
	BuildTags = "unknown"
	BuildTime = "unknown"
	GitCommit = "unknown"
	GoVersion = "unknown"
)

// Font: slant
// http://www.network-science.de/ascii/
func printBanner() {
	fmt.Println()
	fmt.Println(`    __          __    __    __       `)
	fmt.Println(`   / /_  __  __/ /_  / /_  / /__  ___`)
	fmt.Println(`  / __ \/ / / / __ \/ __ \/ / _ \/ _/`)
	fmt.Println(` / /_/ / /_/ / /_/ / /_/ / /  __/ /  `)
	fmt.Println(`/_.___/\__,_/_.___/_.___/_/\___/_/   `)
	fmt.Println()
	fmt.Println("Welcome to use bubbler!")
	fmt.Println("Version:   " + BuildTags)
	fmt.Println("Built:     " + BuildTime)
	fmt.Println("GitCommit: " + GitCommit)
	fmt.Println("GoVersion: " + GoVersion)
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  bubbler [options] <input file>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -t <target>  Target language")
	fmt.Println("  -o <output>  Output Path")
	fmt.Println()
	fmt.Println("Targets:")
	for _, gen := range generator.ListGenerators() {
		fmt.Println("  " + gen)
	}
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  bubble -t c -o output/ example.bb")
	fmt.Println("  bubbler -t c-single -o gen.hpp example.bb")
	fmt.Println("  bubbler -t dump example.bb")
	fmt.Println()
	fmt.Println("For more information, please visit: https://github.com/xaxys/bubbler")
	fmt.Println()
}

func main() {
	if len(os.Args) == 1 {
		printBanner()
		os.Exit(0)
	}

	target := ""
	output := ""
	flag.StringVar(&target, "t", "", "Target Language")
	flag.StringVar(&output, "o", "", "Output Path")
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintln(os.Stderr, "no input files")
		os.Exit(1)
	}
	if len(files) > 1 {
		fmt.Fprintln(os.Stderr, "only single input file is supported")
		for _, f := range files {
			if strings.HasPrefix(f, "-") {
				fmt.Fprintln(os.Stderr, "(please notice that all '-xxx' options should be placed before input file, or they will be treated as input file)")
				break
			}
		}
		os.Exit(1)
	}

	if target == "" {
		fmt.Fprintln(os.Stderr, "no target specified")
		os.Exit(1)
	}
	if !generator.TargetMap.Has(target) {
		fmt.Fprintf(os.Stderr, "target %s is not supported\n", target)
		os.Exit(1)
	}

	input := files[0]
	units, err, warning := compiler.Compile(input)
	if warning != nil {
		fmt.Fprintln(os.Stderr, warning)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = generator.Generate(target, output, units...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
