package parser

import (
	"strings"
)

type ParsedFile struct {
	Dir  string
	Name string

	Contents string

	Imports []Import

	Declarations map[string]Declaration
}

// Parse imports
// Parse import references
// Parse global declarations
// Parse references to global declarations
func ParseFile(
	contents string,
) (*ParsedFile, error) {
	parsed := &ParsedFile{
		Contents: contents,
	}

	// Gather global declarations and imports
	parsed.Declarations = map[string]Declaration{}
	parsed.Imports = []Import{}

	for line := range strings.SplitSeq(contents, "\n") {

		if variable, ok := isGlobalVariable(line); ok {
			parsed.Declarations[variable.Name] = variable

		} else if function, ok := isGlobalFunction(line); ok {
			parsed.Declarations[function.Name] = function

		} else if class, ok := isGlobalClass(line); ok {
			parsed.Declarations[class.Name] = class

		} else if fullModuleImports, ok := isFullModuleImports(line); ok {
			for _, fullModuleImport := range fullModuleImports {
				parsed.Imports = append(parsed.Imports, fullModuleImport)
			}

		} else if declarationImports, ok := isDeclarationImports(line); ok {
			for _, declarationImport := range declarationImports {
				parsed.Imports = append(parsed.Imports, declarationImport)
			}

		}
	}

	return parsed, nil
}

// import asyncio

// from .core import (
//     ensure_active,
//     ble,
//     log_info as y, log_error,
//     log_warn,
//     register_irq_handler,
// )
// from .device import Device, DeviceConnection, DeviceTimeout as x
