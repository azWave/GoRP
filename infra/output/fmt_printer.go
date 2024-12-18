package output

import (
	"fmt"
)

type FmtPrinter struct{}

func (p *FmtPrinter) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (p *FmtPrinter) Println(a ...interface{}) {
	fmt.Println(a...)
}
