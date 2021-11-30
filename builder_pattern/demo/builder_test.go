package demo

import (
	"reflect"
	"testing"
)

// Test 测试建造者模式
func Test(t *testing.T) {
	type fields struct {
		builder PhoneBuilder
	}
	tests := []struct {
		name   string
		fields fields
		want   Phone
	}{
		{
			name:   "01",
			fields: fields{builder: &Iphone11Builder{}},
			want: Phone{
				Model:    "iPhone 11",
				CPU:      "A13",
				Memories: []string{"4GB"},
				Camera: Camera{
					Frontend: "1200万像素ƒ/2.2光圈",
					Backend:  []string{"双1200万后置ƒ/2.4光圈像素超广角及ƒ/1.8光圈广角"},
				},
				Display:  "Liquid视网膜高清显示屏；6.1英寸（对角线）LCD全面屏；多点触控显示屏；1792x828像素分辨率；326 ppi；1400:1对比度(典型)；广色域显示(P3)；625尼特最大亮度(典型)",
				Storages: []string{"64GB", "128GB", "256GB"},
				Colors:   []string{"紫色", "白色", "绿色", "黄色", "黑色", "红色"},
			},
		},
		{
			name:   "02",
			fields: fields{builder: &OnePlus9Builder{}},
			want: Phone{
				Model:    "OnePlus 9",
				CPU:      "高通骁龙888",
				Memories: []string{"8GB", "12GB"},
				Camera: Camera{
					Frontend: "索尼IMX689，索尼IMX766",
					Backend: []string{
						"双1200万后置ƒ/2.4光圈像素超广角及ƒ/1.8光圈广角",
						"5000万像素,超广角摄像头，7P镜头，索尼IMX766",
						"200万像素",
					},
				},
				Display:  "6.55 英寸；2400 X 1080 像素 402 ppi；20:9；柔性 AMOLED 120Hz 瞳孔屏；支持 sRGB、Display P3；康宁® 大猩猩® 玻璃",
				Storages: []string{"128GB", "256GB"},
				Colors:   []string{"鲸蓝色", "黑色", "紫翼色"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				d := &Director{}
				d.SetBuilder(tt.fields.builder)
				if got := d.Build(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Build() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
