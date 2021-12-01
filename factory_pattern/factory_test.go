package factory_pattern

import (
	"reflect"
	"testing"
)

// Test 测试工厂模式
func Test(t *testing.T) {
	type args struct {
		factoryType uint8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "01",
			args: args{factoryType: FruitTypeApple},
			want: "apple",
		},
		{
			name: "02",
			args: args{factoryType: FruitTypeOrange},
			want: "orange",
		},
		{
			name: "03",
			args: args{factoryType: FruitTypeBanana},
			want: "banana",
		},
		{
			name: "04",
			args: args{factoryType: 4},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				fruit := ProduceFruit(tt.args.factoryType)
				if got := fruit.GetName(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("test failed, got %v, but want %v.", got, tt.want)
				}
			},
		)
	}
}
