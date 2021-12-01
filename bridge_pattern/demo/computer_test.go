package demo

import (
	"fmt"
	"testing"
)

// 测试桥接模式
func Test(t *testing.T) {
	hp := HPPrinter{}
	xm := XMPrinter{}

	mac := Mac{}

	mac.SetPrinter(hp)
	mac.Print()
	fmt.Println()

	mac.SetPrinter(xm)
	mac.Print()
	fmt.Println()

	windows := Windows{}

	windows.SetPrinter(hp)
	windows.Print()
	fmt.Println()

	windows.SetPrinter(xm)
	windows.Print()
	fmt.Println()
}
