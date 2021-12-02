package composite_pattern

import (
	"fmt"
	"testing"
)

// Test 测试装饰模式
func Test(t *testing.T) {
	orange := OrangeIceCream{IceCream: BaseIceCream{}}
	orangeAndStrawberry := StrawberryIceCream{IceCream: orange}
	fmt.Println(orangeAndStrawberry.GetPrice())
}
