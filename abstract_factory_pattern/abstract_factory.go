package abstract_factory_pattern

// FruitFactory 水果工厂
type FruitFactory interface {
	// CreateFruit 创建水果
	CreateFruit() Fruit
}

// Fruit 水果接口
type Fruit interface {
	// GetName 获取水果名称
	GetName() string
}
