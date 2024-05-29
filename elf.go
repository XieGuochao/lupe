package main

import (
	"debug/elf"
	"fmt"
)

func loadElfFile(binPath string) {
	elfFile, err := elf.Open(binPath)
	if err != nil {
		panic(err)
	}
	defer elfFile.Close()

	symbols, err := elfFile.Symbols()
	if err != nil {
		panic(err)
	}
	for index, symbol := range symbols {
		fmt.Printf("%d %s\n", index, symbol.Name)
	}
}
