package parser

import (
	"strings"
)

type Import interface {
	Name() string
}

type FullModuleImport struct {
	ModulePath []string
	ImportedAs string
}

func (i FullModuleImport) Name() string {
	return i.ImportedAs
}

func isFullModuleImports(line string) ([]FullModuleImport, bool) {
	if rawImport, ok := strings.CutPrefix(line, "import "); ok {
		return parseFullModuleImports(rawImport)
	}

	return nil, false
}

func parseFullModuleImports(rawPackages string) ([]FullModuleImport, bool) {
	individualImports := strings.Split(rawPackages, ",")
	out := make([]FullModuleImport, len(individualImports))
	for i, individualImport := range individualImports {
		var ok bool
		out[i], ok = parseFullModuleImport(individualImport)
		if !ok {
			return out, false
		}
	}

	return out, true
}

func parseFullModuleImport(rawImport string) (FullModuleImport, bool) {
	parts := strings.Split(rawImport, " as ")
	modulePath := strings.Split(parts[0], ".")

	switch len(parts) {
	case 2:
		return FullModuleImport{
			ModulePath: modulePath,
			ImportedAs: parts[1],
		}, true
	case 1:
		return FullModuleImport{
			ModulePath: modulePath,
			ImportedAs: modulePath[len(modulePath)-1],
		}, true
	default:
		return FullModuleImport{}, false
	}
}

type DeclarationImport struct {
	ModulePath      []string
	DeclarationName string
	ImportedAs      string
}

func (i DeclarationImport) Name() string {
	return i.ImportedAs
}

func isDeclarationImports(line string) ([]DeclarationImport, bool) {
	if rawImport, ok := strings.CutPrefix(line, "from "); ok {
		return parseDeclarationImports(rawImport)
	}

	return nil, false
}

func parseDeclarationImports(rawImport string) ([]DeclarationImport, bool) {
	parts := strings.Split(rawImport, " import ")
	if len(parts) != 2 {
		return nil, false
	}

	modulePath := strings.Split(parts[0], ".")

	rawItems := strings.Split(parts[1], ",")
	out := make([]DeclarationImport, len(rawItems))
	for i, rawItem := range rawItems {
		var ok bool
		out[i], ok = parseDeclarationImport(modulePath, rawItem)
		if !ok {
			return out, false
		}
	}

	return out, true
}

func parseDeclarationImport(modulePath []string, rawItem string) (DeclarationImport, bool) {
	rawItem = strings.TrimSpace(rawItem)
	parts := strings.Split(rawItem, " as ")

	if len(parts) == 2 {
		return DeclarationImport{
			ModulePath:      modulePath,
			DeclarationName: parts[0],
			ImportedAs:      parts[1],
		}, true
	} else if len(parts) == 1 {
		return DeclarationImport{
			ModulePath:      modulePath,
			DeclarationName: parts[0],
			ImportedAs:      parts[0],
		}, true
	}

	return DeclarationImport{}, false
}
