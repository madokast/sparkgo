package exports

func (fun) PrintHello() {
	println("Hwllo, world!")
}

func (fun) AddInt64(a, b int64) int64 {
	return a + b
}

// --------------- reflect helpers -----------------------

type fun struct{}

var F = fun{}
