package main

import (
	"fmt"
	"log"

	"github.com/bytecodealliance/wasmtime-go/v9"
)

func main() {
	// Set the path to the WebAssembly module
	modulePath := "wasm-rust-setup/wasm/wasm_rust_setup_bg.wasm"

	// Create an engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Load the WebAssembly module
	module, err := wasmtime.NewModuleFromFile(store.Engine, modulePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Module")
	fmt.Println(module)
	// Instantiate the module
	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Instance")
	fmt.Println(instance)

	// Get the exported "generate_fibonacci" function
	generateFibonacci := instance.GetExport(store, "generate_fibonacci")
	if generateFibonacci == nil {
		log.Fatal("Failed to find function export `generate_fibonacci`")
	}
	fmt.Println("generate Fibonacci")
	fmt.Println(generateFibonacci)
	// Call the "generate_fibonacci" function
	_, err = generateFibonacci.Func().Call(store, 5) // Pass the desired value for `n`
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fibonacci sequence generated and written to file")
}

// func init() {
// 	// Set the working directory to the directory of the Go executable
// 	executablePath, err := os.Executable()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	exeDir := filepath.Dir(executablePath)
// 	err = os.Chdir(exeDir)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
