package demo

// OnePlus9Builder OnePlus9构造器
type OnePlus9Builder struct {
	Phone Phone `json:"phone"`
}

// NewPhone 生成新手机
func (i *OnePlus9Builder) NewPhone() {
	i.Phone = Phone{}
}

// BuildModel 构建型号
func (i *OnePlus9Builder) BuildModel() {
	i.Phone.Model = "OnePlus 9"
}

// BuildCPU 构建CPU
func (i *OnePlus9Builder) BuildCPU() {
	i.Phone.CPU = "高通骁龙888"
}

// BuildMemories 构建内存
func (i *OnePlus9Builder) BuildMemories() {
	i.Phone.Memories = []string{"8GB", "12GB"}
}

// BuildCamera 构建摄像头
func (i *OnePlus9Builder) BuildCamera() {
	i.Phone.Camera = Camera{
		Frontend: "索尼IMX689，索尼IMX766",
		Backend: []string{
			"双1200万后置ƒ/2.4光圈像素超广角及ƒ/1.8光圈广角",
			"5000万像素,超广角摄像头，7P镜头，索尼IMX766",
			"200万像素",
		},
	}
}

// BuildDisplay 构建屏幕
func (i *OnePlus9Builder) BuildDisplay() {
	i.Phone.Display = "6.55 英寸；2400 X 1080 像素 402 ppi；20:9；柔性 AMOLED 120Hz 瞳孔屏；支持 sRGB、Display P3；康宁® 大猩猩® 玻璃"
}

// BuildStorages 构建容量
func (i *OnePlus9Builder) BuildStorages() {
	i.Phone.Storages = []string{"128GB", "256GB"}
}

// BuildColors 构建颜色
func (i *OnePlus9Builder) BuildColors() {
	i.Phone.Colors = []string{"鲸蓝色", "黑色", "紫翼色"}
}

// Build 构建
func (i *OnePlus9Builder) Build() Phone {
	return i.Phone
}
