package main

import (
	"debug/dwarf"
	"debug/elf"
	"fmt"
)

type ElfData struct {
	Dwarf   *dwarf.Data
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

// TODO: Implement this
func (elfData *ElfData) ListDwarf() {
	reader := elfData.Dwarf.Reader()
	for {
		entry, err := reader.Next()
		if err != nil {
			panic(err)
		}
		if entry == nil {
			break
		}
		fmt.Println(entry)
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

	d, err := elfFile.DWARF()
	if err != nil {
		panic(err)
	}

	data := ElfData{
		Dwarf:   d,
		File:    elfFile,
		Symbols: symbols,
	}
	return data
}
