# gosymdump

A lightweight Go utility to extract and display function symbols from Go binaries by reading the `.gopclntab` section.

## Overview

`gosymdump` parses ELF-format Go binaries to extract function information from the Go program counter line table (`.gopclntab`). This is useful for:

- Reverse engineering Go binaries
- Understanding function layouts and boundaries
- Analyzing stripped Go binaries
- Security research and binary analysis
- Debugging and profiling

## Features

- Extracts function names and memory addresses from Go binaries
- Works with Go 1.3+ binaries (where `.gosymtab` may be empty)
- Displays function entry and exit points
- No external dependencies beyond Go standard library
- Simple, focused tool for quick symbol extraction

## Installation

```bash
go install github.com/maxgio92/gosymdump@latest
```

### From Source

```bash
git clone https://github.com/maxgio92/gosymdump.git
cd gosymdump
go build
```

## Usage

```bash
gosymdump <path-to-go-binary>
```

### Example

```bash
$ gosymdump /usr/bin/kubectl
Functions in binary:
0x401000 - 0x401020: runtime.text
0x401020 - 0x401040: runtime._rt0_amd64_linux
0x401040 - 0x401060: main.main
...
```

## How It Works

1. **Opens the ELF binary** using Go's `debug/elf` package
2. **Reads the `.gopclntab` section** which contains the program counter line table
3. **Locates the `.text` section** to get the base address for code
4. **Creates a line table** from the PC data
5. **Parses the symbol table** using `debug/gosym` package
6. **Displays all functions** with their memory address ranges and names

## Technical Details

### Go Symbol Tables

Go binaries contain two relevant sections:
- **`.gopclntab`**: Program counter to line number mapping (required)
- **`.gosymtab`**: Traditional symbol table (often empty in Go 1.3+)

This tool uses the `.gopclntab` section, which contains function boundaries and names even in stripped binaries compiled with modern Go versions.

### Supported Binary Formats

- ELF binaries (Linux, BSD)
- Go 1.3 and later

## Requirements

- Go 1.16 or later
- Target binaries must be Go binaries in ELF format

## Limitations

- Only works with ELF format binaries (Linux/BSD)
- Does not support PE (Windows) or Mach-O (macOS) formats
- Requires the `.gopclntab` section to be present (standard in Go binaries)

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## Related Projects

- [GoReSym](https://github.com/mandiant/GoReSym) - Comprehensive Go symbol recovery tool

## Related Tools

- `go tool objdump` - Disassemble Go binaries (part of standard Go toolchain)
- `go tool nm` - List symbols in Go object files (part of standard Go toolchain)

## Related Literature

- [Golang Internals: Symbol Recovery](https://cloud.google.com/blog/topics/threat-intelligence/golang-internals-symbol-recovery/) - Google Cloud's comprehensive guide on Go symbol recovery techniques

## References

- [Go debug/gosym package](https://pkg.go.dev/debug/gosym)
- [Go debug/elf package](https://pkg.go.dev/debug/elf)
- [Go Binary Format](https://golang.org/src/debug/gosym/pclntab.go)
