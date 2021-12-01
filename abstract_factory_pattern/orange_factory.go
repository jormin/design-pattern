package abstract_factory_pattern

// OrangeFactory 橘子工厂
type OrangeFactory struct{}

// CreateFruit 创建水果
func (f *OrangeFactory) CreateFruit() Fruit {
	return orange{}
}

// orange 橘子
type orange struct{}

// GetName 获取名称
func (a orange) GetName() string {
	return "orange"
}
