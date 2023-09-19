package main

/*
go build -buildmode=plugin -o export.so export.go
*/

func PrintHello() {
	println("Hello, world!")
}

func AddInt64(a, b int64) int64 {
	return a + b
}
