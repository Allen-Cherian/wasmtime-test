// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/bytecodealliance/wasmtime-go/v9"
// )

// func main() {
// 	config := wasmtime.NewConfig()
// 	config.SetConsumeFuel(true)
// 	engine := wasmtime.NewEngineWithConfig(config)
// 	store := wasmtime.NewStore(engine)
// 	err := store.AddFuel(10000)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Compile and instantiate a small example with an infinite loop.
// 	wasm, err := wasmtime.Wat2Wasm(`
// 	(module
// 	  (func $fibonacci (param $n i32) (result i32)
// 	    (if
// 	      (i32.lt_s (local.get $n) (i32.const 2))
// 	      (return (local.get $n))
// 	    )
// 	    (i32.add
// 	      (call $fibonacci (i32.sub (local.get $n) (i32.const 1)))
// 	      (call $fibonacci (i32.sub (local.get $n) (i32.const 2)))
// 	    )
// 	  )
// 	  (export "fibonacci" (func $fibonacci))
// 	)
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	module, err := wasmtime.NewModule(store.Engine, wasm)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

//		// Invoke `fibonacci` export with higher and higher numbers until we exhaust our fuel.
//		fibonacci := instance.GetFunc(store, "fibonacci")
//		if fibonacci == nil {
//			log.Fatal("Failed to find function export `fibonacci`")
//		}
//		for n := 0; ; n++ {
//			fuelBefore, _ := store.FuelConsumed()
//			output, err := fibonacci.Call(store, n)
//			if err != nil {
//				break
//			}
//			fuelAfter, _ := store.FuelConsumed()
//			fmt.Printf("fib(%d) = %d [consumed %d fuel]\n", n, output, fuelAfter-fuelBefore)
//			err = store.AddFuel(fuelAfter - fuelBefore)
//			if err != nil {
//				log.Fatal(err)
//			}
//		}
//	}
package main

import (
	"fmt"
	"log"

	"github.com/bytecodealliance/wasmtime-go/v9"
)

func main() {
	// Create the Wasmtime configuration and engine.
	config := wasmtime.NewConfig()
	config.SetConsumeFuel(true)
	engine := wasmtime.NewEngineWithConfig(config)

	// Create the store and add fuel.
	store := wasmtime.NewStore(engine)
	err := store.AddFuel(10000)
	if err != nil {
		log.Fatal(err)
	}

	// Load the WebAssembly module generated by Rust.
	module, err := wasmtime.NewModuleFromFile(store.Engine, "wasm-rust-setup/wasm/wasm_rust_setup_bg.wasm")
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate the WebAssembly module.
	instance, err := wasmtime.NewInstance(store, module, []wasmtime.AsExtern{})
	if err != nil {
		log.Fatal(err)
	}

	// Get the exported `fibonacci` function.
	fibonacci := instance.GetExport(store, "fibonacci").Func()
	if fibonacci == nil {
		log.Fatal("Failed to find function export `fibonacci`")
	}

	write_state := instance.GetExport(store, "write_state").Func()
	fmt.Println(write_state)
	if write_state == nil {
		log.Fatal("Failed at write_state")
	}
	// Invoke the `fibonacci` function with higher and higher numbers until we exhaust our fuel.
	for n := 0; ; n++ {
		fuelBefore, _ := store.FuelConsumed()
		output, err := fibonacci.Call(store, int32(n))
		if err != nil {
			break
		}
		fuelAfter, _ := store.FuelConsumed()
		fmt.Printf("fib(%d) = %d [consumed %d fuel]\n", n, output.(int32), fuelAfter-fuelBefore)
		err = store.AddFuel(fuelAfter - fuelBefore)
		if err != nil {
			log.Fatal(err)
		}
	}
}
