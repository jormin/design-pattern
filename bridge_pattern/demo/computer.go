package demo

import "fmt"

// Computer 计算机-抽象层
type Computer interface {
	// SetPrinter 设置打印机
	SetPrinter(printer Printer)
	// Print 打印
	Print()
}

// Mac 苹果电脑
type Mac struct {
	printer Printer
}

// SetPrinter 设置打印机
func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}

// Print 打印
func (m *Mac) Print() {
	fmt.Println("苹果电脑打印")
	m.printer.Print()
}

// Windows windows电脑
type Windows struct {
	printer Printer
}

// SetPrinter 设置打印机
func (w *Windows) SetPrinter(printer Printer) {
	w.printer = printer
}

// Print 打印
func (w *Windows) Print() {
	fmt.Println("Windows电脑打印")
	w.printer.Print()
}
