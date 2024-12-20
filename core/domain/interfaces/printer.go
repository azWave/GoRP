package interfaces

type Printer interface {
	Print(a ...interface{})
	Printf(format string, a ...interface{})
	Println(a ...interface{})
}
