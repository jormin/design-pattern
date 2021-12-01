package demo

import (
	"reflect"
	"testing"
)

// 测试原型模式
func Test(t *testing.T) {
	stu1 := &Student{
		Name:   "A",
		Age:    20,
		Gender: "男",
		Class:  &Class{Name: "一班"},
	}
	stu2 := stu1.Clone()
	stu3 := stu1.DeepClone()

	stu1.Name = "B"
	stu1.Age = 25
	stu1.Gender = "女"
	stu1.Class = &Class{Name: "二班"}

	if !reflect.DeepEqual(stu1, stu2) {
		t.Error("Clone failed")
	}

	if reflect.DeepEqual(stu1, stu3) {
		t.Error("DeepClone failed")
	}
}
