// infra/output/mock_printer.go
package output

import (
	"bytes"
	"fmt"
)

type MockPrinter struct {
	Output bytes.Buffer
}

func (p *MockPrinter) Print(a ...interface{}) {
	p.Output.WriteString(fmt.Sprint(a...))
}

func (p *MockPrinter) Printf(format string, a ...interface{}) {
	p.Output.WriteString(fmt.Sprintf(format, a...))
}

func (p *MockPrinter) Println(a ...interface{}) {
	p.Output.WriteString(fmt.Sprintln(a...))
}
