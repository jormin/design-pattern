package demo

// BananaFactory 香蕉工厂
type BananaFactory struct{}

// CreateFruit 创建水果
func (f *BananaFactory) CreateFruit() Fruit {
	return banana{}
}

// banana 香蕉
type banana struct{}

// GetName 获取名称
func (a banana) GetName() string {
	return "banana"
}
