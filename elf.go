package main

import (
	"debug/elf"
	"fmt"
)

type ElfData struct {
	File    *elf.File
	Symbols []elf.Symbol
}

func (elfData *ElfData) Close() {
	elfData.File.Close()
}

func (elfData *ElfData) ListSymbols() {
	for index, symbol := range elfData.Symbols {
		fmt.Printf("%d %s\n", index, symbol.Name)
	}
}

func loadElfFile(binPath string) ElfData {
	elfFile, err := elf.Open(binPath)
	if err != nil {
		panic(err)
	}

	symbols, err := elfFile.Symbols()
	if err != nil {
		panic(err)
	}

	data := ElfData{
		File:    elfFile,
		Symbols: symbols,
	}
	return data
}
