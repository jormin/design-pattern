package adapter_pattern

import (
	"testing"
)

// Test 测试适配器模式
func Test(t *testing.T) {
	var newInterface NewInterface
	newInterface = &Adapter{oldInterface: Adaptee{}}
	_, _ = newInterface.SaveData()
}
