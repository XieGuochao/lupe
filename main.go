package main

import "flag"

type Options struct {
	BinPath string
	List    bool
	Verbose bool
}

func main() {
	options := Options{}
	flag.StringVar(&options.BinPath, "bin", "", "Binary path")
	flag.BoolVar(&options.List, "list", false, "List symbols")
	flag.BoolVar(&options.Verbose, "verbose", false, "Verbose output")

	flag.Parse()

	// Require at least one args
	if options.BinPath == "" {
		panic("Binary path is required.")
	}

	handle(options)
}

func handle(options Options) {
	elf := loadElfFile(options.BinPath)
	defer elf.Close()

	if options.List {
		elf.ListSymbols()
	}
}
