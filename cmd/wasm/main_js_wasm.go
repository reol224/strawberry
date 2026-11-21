package main

import "syscall/js"

var runFunc js.Func

func main() {
	runFunc = js.FuncOf(run)
	js.Global().Set("strawberryRun", runFunc)
	select {}
}

func run(_ js.Value, args []js.Value) interface{} {
	response := js.Global().Get("Object").New()
	response.Set("output", "")
	response.Set("error", "")

	if len(args) != 1 || args[0].Type() != js.TypeString {
		response.Set("error", "Expected one source string.")
		return response
	}

	output, err := runSource(args[0].String())
	if err != nil {
		response.Set("error", err.Error())
		return response
	}

	response.Set("output", output)
	return response
}
