package composite_pattern

// IceCream 冰激凌
type IceCream interface {
	// GetPrice 获取价格
	GetPrice() int
}

// BaseIceCream 基础冰激凌
type BaseIceCream struct{}

// GetPrice 获取价格
func (b BaseIceCream) GetPrice() int {
	return 7
}

// OrangeIceCream 橘子冰激凌
type OrangeIceCream struct {
	IceCream IceCream
}

// GetPrice 获取价格
func (o OrangeIceCream) GetPrice() int {
	return o.IceCream.GetPrice() + 3
}

// StrawberryIceCream 草莓冰激凌
type StrawberryIceCream struct {
	IceCream IceCream
}

// GetPrice 获取价格
func (s StrawberryIceCream) GetPrice() int {
	return s.IceCream.GetPrice() + 5
}
