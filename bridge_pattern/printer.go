package bridge_pattern

import "fmt"

// Printer 打印机-实施层
type Printer interface {
	// Print 打印
	Print()
}

// HPPrinter 惠普打印机
type HPPrinter struct{}

// Print 打印
func (h HPPrinter) Print() {
	fmt.Println("惠普打印机打印")
}

// XMPrinter 小米打印机
type XMPrinter struct{}

// Print 打印
func (x XMPrinter) Print() {
	fmt.Println("小米打印机打印")
}
