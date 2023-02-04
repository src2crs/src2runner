package sourcecode

import (
	"fmt"
)

// Example_importandfunction shows how to assemble a source string
// containing an import statement and a function.
func Example_importandfunction() {
	// Initialize a source code container.
	src := New()

	// The package to be imported.
	packageString := "fmt"
	src.AddImports(packageString)

	// The function code to be added.
	functionString := "func Greet() {\n  fmt.Println(\"Hello SourceCode\")\n}\n"
	src.AddCode(functionString)

	// Assemble the complete source code.
	fmt.Println(src.Source())

	// Output:
	// import (
	// "fmt"
	// )
	//
	// func Greet() {
	//   fmt.Println("Hello SourceCode")
	// }
}
