package main

import (
	"debug/elf"
	"debug/gosym"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run script.go <binary>")
	}

	// Open the ELF binary
	exe, err := elf.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer exe.Close()

	// Read .gopclntab section
	pclnSec := exe.Section(".gopclntab")
	if pclnSec == nil {
		log.Fatal("No .gopclntab section found")
	}
	pclnData, err := pclnSec.Data()
	if err != nil {
		log.Fatalf("Cannot read .gopclntab: %v", err)
	}

	// Create LineTable using .text section address
	textSec := exe.Section(".text")
	if textSec == nil {
		log.Fatal("No .text section found")
	}
	lineTable := gosym.NewLineTable(pclnData, textSec.Addr)

	// Parse symbol table (even if .gosymtab is empty, lineTable has function info)
	// In Go 1.3+, .gosymtab may be empty; LineTable contains function boundaries
	table, err := gosym.NewTable(nil, lineTable)
	if err != nil {
		log.Fatalf("Cannot create symbol table: %v", err)
	}

	// Print all functions
	fmt.Println("Functions in binary:")
	for _, fn := range table.Funcs {
		fmt.Printf("0x%x - 0x%x: %s\n", fn.Entry, fn.End, fn.Name)
	}
}   
