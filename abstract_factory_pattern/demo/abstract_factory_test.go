package demo

import (
	"reflect"
	"testing"
)

// Test 测试抽象工厂模式
func Test(t *testing.T) {
	type args struct {
		factory FruitFactory
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{factory: &AppleFactory{}},
			want: "apple",
		},
		{
			name: "02",
			args: args{factory: &OrangeFactory{}},
			want: "orange",
		},
		{
			name: "03",
			args: args{factory: &BananaFactory{}},
			want: "banana",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := tt.args.factory.CreateFruit().GetName(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("test failed, got %v, but want %v.", got, tt.want)
				}
			},
		)
	}
}
