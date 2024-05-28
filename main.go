package main

import "flag"

func main() {
	flag.Parse()

	// Require at least one args
	if len(flag.Args()) < 1 {
		panic("at least one argument: <binary path>")
	}
	binPath := flag.Args()[len(flag.Args())-1]

	handle(binPath)
}

func handle(binPath string) {
	loadElfFile(binPath)
}
