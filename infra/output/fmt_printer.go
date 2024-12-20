package output

import (
	"fmt"
)

type FmtPrinter struct{}

func (p *FmtPrinter) Print(a ...interface{}) {
	fmt.Print(a...)
}

func (p *FmtPrinter) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (p *FmtPrinter) Println(a ...interface{}) {
	fmt.Println(a...)
}
