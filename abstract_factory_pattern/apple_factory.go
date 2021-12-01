package abstract_factory_pattern

// AppleFactory 苹果工厂
type AppleFactory struct{}

// CreateFruit 创建水果
func (f *AppleFactory) CreateFruit() Fruit {
	return Apple{}
}

// Apple 苹果
type Apple struct{}

// GetName 获取名称
func (a Apple) GetName() string {
	return "apple"
}
