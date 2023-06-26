package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bytecodealliance/wasmtime-go/v9"
)

func main() {
	// Set the path to the WebAssembly module
	modulePath := "./wasm/fibonacci_bg.wasm"

	// Create a Wasmtime engine and store
	engine := wasmtime.NewEngine()
	store := wasmtime.NewStore(engine)

	// Load the WebAssembly module
	module, err := wasmtime.NewModuleFromFile(engine, modulePath)
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate the module
	instance, err := wasmtime.NewInstance(store, module, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Get the exported "generate_fibonacci" function
	generateFibonacci := instance.GetExport(store, "generate_fibonacci").Func()
	if generateFibonacci == nil {
		log.Fatal("Failed to find function export `generate_fibonacci`")
	}

	// Call the "generate_fibonacci" function
	err = generateFibonacci.Call(store, 10) // Pass the desired value for `n`
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Fibonacci sequence generated and written to file")
}

func init() {
	// Set the working directory to the directory of the Go executable
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exeDir := filepath.Dir(executablePath)
	err = os.Chdir(exeDir)
	if err != nil {
		log.Fatal(err)
	}
}
