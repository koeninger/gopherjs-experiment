package main

import (
	"syscall/js"
)

func main() {
	var eventCb js.Func
	eventCb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		var resp = js.Global().Get("Response").New("Hello from Go")
		args[0].Call("respondWith", resp)
		return nil
	})
	js.Global().Call("addEventListener", "fetch", eventCb)
}
