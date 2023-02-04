package interpreter

import "github.com/src2crs/srcrunner/sourcecode"

// Example_importandfunction assembles a source
// string with a function and runs that function.
func Example_importandfunction() {
	// Create the source code to be evaluated.
	src := sourcecode.New()
	src.AddImports("fmt")
	src.AddCode("func Greet() {\n  fmt.Println(\"Hello Interpreter\")\n}\n")

	// Create an Interpreter.
	// This object will hold its state between succesive calls to its Evaluate method.
	interpreter := New()

	// Evaluate the source code.
	// This produces no output, it is similar to compiling the code.
	interpreter.Evaluate(src.Source())

	// Evaluate a call of the function Greet.
	// This will run the function.
	// As Greet calls fmt.Println, we will see its output on the console.
	interpreter.Evaluate("Greet()")

	// Extract the function Greet as a callable object.
	interpreter.Evaluate("Greet")
	function := interpreter.LastResult().(func())

	// Call that function
	function()

	// Output:
	// Hello Interpreter
	// Hello Interpreter
}
