package builder_pattern

// Iphone11Builder iPhone11构造器
type Iphone11Builder struct {
	Phone Phone `json:"phone"`
}

// NewPhone 生成新手机
func (i *Iphone11Builder) NewPhone() {
	i.Phone = Phone{}
}

// BuildModel 构建型号
func (i *Iphone11Builder) BuildModel() {
	i.Phone.Model = "iPhone 11"
}

// BuildCPU 构建CPU
func (i *Iphone11Builder) BuildCPU() {
	i.Phone.CPU = "A13"
}

// BuildMemories 构建内存
func (i *Iphone11Builder) BuildMemories() {
	i.Phone.Memories = []string{"4GB"}
}

// BuildCamera 构建摄像头
func (i *Iphone11Builder) BuildCamera() {
	i.Phone.Camera = Camera{
		Frontend: "1200万像素ƒ/2.2光圈",
		Backend:  []string{"双1200万后置ƒ/2.4光圈像素超广角及ƒ/1.8光圈广角"},
	}
}

// BuildDisplay 构建屏幕
func (i *Iphone11Builder) BuildDisplay() {
	i.Phone.Display = "Liquid视网膜高清显示屏；6.1英寸（对角线）LCD全面屏；多点触控显示屏；1792x828像素分辨率；326 ppi；1400:1对比度(典型)；广色域显示(P3)；625尼特最大亮度(典型)"
}

// BuildStorages 构建容量
func (i *Iphone11Builder) BuildStorages() {
	i.Phone.Storages = []string{"64GB", "128GB", "256GB"}
}

// BuildColors 构建颜色
func (i *Iphone11Builder) BuildColors() {
	i.Phone.Colors = []string{"紫色", "白色", "绿色", "黄色", "黑色", "红色"}
}

// Build 构建
func (i *Iphone11Builder) Build() Phone {
	return i.Phone
}
