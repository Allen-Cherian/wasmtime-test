package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bytecodealliance/wasmtime-go"
)

type WasmtimeRuntime struct {
	store   *wasmtime.Store
	memory  *wasmtime.Memory
	handler *wasmtime.Func

	input  []byte
	output []byte
}

func (r *WasmtimeRuntime) Init(wasmFile string) {
	engine := wasmtime.NewEngine()
	linker := wasmtime.NewLinker(engine)
	linker.DefineWasi()
	wasiConfig := wasmtime.NewWasiConfig()
	r.store = wasmtime.NewStore(engine)
	r.store.SetWasi(wasiConfig)
	linker.FuncWrap("env", "load_input", r.loadInput)
	linker.FuncWrap("env", "dump_output", r.dumpOutput)
	wasmBytes, _ := os.ReadFile(wasmFile)
	module, _ := wasmtime.NewModule(r.store.Engine, wasmBytes)
	instance, _ := linker.Instantiate(r.store, module)
	r.memory = instance.GetExport(r.store, "memory").Memory()
	r.handler = instance.GetFunc(r.store, "handler")
}

func (r *WasmtimeRuntime) loadInput(pointer int32) {
	copy(r.memory.UnsafeData(r.store)[pointer:pointer+int32(len(r.input))], r.input)
}

func (r *WasmtimeRuntime) dumpOutput(pointer int32, length int32) {
	r.output = make([]byte, length)
	copy(r.output, r.memory.UnsafeData(r.store)[pointer:pointer+length])
}

func (r *WasmtimeRuntime) RunHandler(data []byte) []byte {
	r.input = data
	r.handler.Call(r.store, len(data))
	return r.output
}

func main() {
	size := 16
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = 'a'
	}

	runtime := &WasmtimeRuntime{}
	runtime.Init("wasm-rust-setup/target/wasm32-wasi/release/wasm_rust_setup.wasm")
	for i := 0; i < size; i++ {
		time.Sleep(100 * time.Millisecond)
		output := runtime.RunHandler(buf)
		fmt.Println(string(output))
	}
}
