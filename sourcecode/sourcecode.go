package sourcecode

import (
	"fmt"
	"strings"
)

// TODO: Add documentation.
// TODO: Use ast internally? Or is it better to stay with this very simple version?
// TODO: Add function to read a file.

// SourceCode is a container for source code elements like imports or functions.
// It exposes an interface for adding such elements and constructing a string
// containing the actual code.
type SourceCode struct {
	imports      []string
	codesnippets []string
}

// New creates an empty source code object.
func New() *SourceCode {
	return &SourceCode{[]string{}, []string{}}
}

// AddImports expects a list of strings and adds import statements for each of them to the code.
// These strings should be package names as they would appear in import statements.
func (sourcecode *SourceCode) AddImports(importLines ...string) {
	sourcecode.imports = append(sourcecode.imports, importLines...)
}

// AddCodes expects an arbitrary string of code and adds it to the sourcecode.
func (sourcecode *SourceCode) AddCode(functionCode string) {
	sourcecode.codesnippets = append(sourcecode.codesnippets, functionCode)
}

// Source assembles all code elements that have been added previously and returns the result.
func (sourcecode *SourceCode) Source() string {
	importLines := strings.Join(sourcecode.imports, "\n")
	imports := fmt.Sprintf("import (\n\"%s\"\n)", importLines)
	functions := strings.Join(sourcecode.codesnippets, "\n\n")

	result := fmt.Sprintf("%s\n\n%s\n", imports, functions)
	return result
}
