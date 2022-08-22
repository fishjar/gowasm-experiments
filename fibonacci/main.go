package main

import "syscall/js"

func fib(i int) int {
	if i == 0 || i == 1 {
		return i
	}
	return fib(i-1) + fib(i-2)
}

func fibFunc(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(fib(args[0].Int()))
}

func main() {
	done := make(chan int)
	js.Global().Set("fibFunc", js.FuncOf(fibFunc))
	<-done
}
