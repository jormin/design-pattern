package demo

// Phone 手机
type Phone struct {
	Model    string   `json:"model" desc:"型号"`
	CPU      string   `json:"cpu" desc:"CPU"`
	Memories []string `json:"memory" desc:"内存"`
	Camera   Camera   `json:"camera" desc:"摄像头"`
	Display  string   `json:"display" desc:"屏幕"`
	Storages []string `json:"storage" desc:"容量"`
	Colors   []string `json:"color" desc:"颜色"`
}

// Camera 摄像头
type Camera struct {
	Frontend string   `json:"frontend" desc:"前置摄像头"`
	Backend  []string `json:"backend" desc:"后置摄像头"`
}

// PhoneBuilder 手机制造
type PhoneBuilder interface {
	NewPhone()
	BuildModel()
	BuildCPU()
	BuildMemories()
	BuildCamera()
	BuildDisplay()
	BuildStorages()
	BuildColors()
	Build() Phone
}

// Director 指挥者
type Director struct {
	builder PhoneBuilder
}

// SetBuilder 设置构造器
func (d *Director) SetBuilder(builder PhoneBuilder) {
	d.builder = builder
}

// Build 制造
func (d *Director) Build() Phone {
	d.builder.NewPhone()
	d.builder.BuildModel()
	d.builder.BuildCPU()
	d.builder.BuildMemories()
	d.builder.BuildCamera()
	d.builder.BuildDisplay()
	d.builder.BuildStorages()
	d.builder.BuildColors()
	return d.builder.Build()
}
