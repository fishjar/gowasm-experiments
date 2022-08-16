// main.go
package main

import "syscall/js"

func fib(i int) int {
	if i == 0 || i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	// js.Value 可以将 Js 的值转换为 Go 的值，比如 args[0].Int()，则是转换为 Go 语言中的整型。
	// js.ValueOf，则用来将 Go 的值，转换为 Js 的值。
	return js.ValueOf(fib(args[0].Int()))
}

func main() {
	done := make(chan int)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
