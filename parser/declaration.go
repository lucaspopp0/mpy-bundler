package parser

import "regexp"

type DeclarationKind int

const (
	DeclarationUnknown DeclarationKind = iota
	DeclarationVariable
	DeclarationFunction
	DeclarationClass
)

func (k DeclarationKind) String() string {
	switch k {
	case DeclarationVariable:
		return "variable"
	case DeclarationFunction:
		return "function"
	case DeclarationClass:
		return "class"
	default:
		return "unknown"
	}
}

type Declaration struct {
	Kind DeclarationKind
	Name string
}

var globalVariableExpr = regexp.MustCompile(`^([_a-zA-Z][_a-zA-Z0-9]*) *=[^=]`)

func isGlobalVariable(line string) (variable Declaration, ok bool) {
	captureGroups := globalVariableExpr.FindStringSubmatch(line)
	if len(captureGroups) < 2 {
		ok = false
		return
	}

	return Declaration{
		Kind: DeclarationVariable,
		Name: captureGroups[1],
	}, true
}

var globalFunctionExpr = regexp.MustCompile(`^def ([_a-zA-Z][_a-zA-Z0-9]*)\(`)

func isGlobalFunction(line string) (function Declaration, ok bool) {
	captureGroups := globalFunctionExpr.FindStringSubmatch(line)
	if len(captureGroups) < 2 {
		ok = false
		return
	}

	return Declaration{
		Kind: DeclarationFunction,
		Name: captureGroups[1],
	}, true
}

var globalClassExpr = regexp.MustCompile(`^class ([_a-zA-Z][_a-zA-Z0-9]*)`)

func isGlobalClass(line string) (class Declaration, ok bool) {
	captureGroups := globalClassExpr.FindStringSubmatch(line)
	if len(captureGroups) < 2 {
		ok = false
		return
	}

	return Declaration{
		Kind: DeclarationClass,
		Name: captureGroups[1],
	}, true
}
