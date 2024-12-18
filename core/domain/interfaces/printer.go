package interfaces

type Printer interface {
	Printf(format string, a ...interface{})
	Println(a ...interface{})
}
