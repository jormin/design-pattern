package demo

import (
	"fmt"
	"testing"
)

// Test 测试单例模式
func Test(t *testing.T) {
	counter := GetSingletonCounter()
	for want := 1; want <= 10; want++ {
		counter.Count()
		got := counter.GetNum()
		fmt.Println(&counter, got, want)
		if got != want {
			t.Errorf("test failed, got %v, but want %v", got, want)
		}
	}
}
