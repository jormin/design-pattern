package demo

const (
	// FruitTypeApple 水果类型：苹果
	FruitTypeApple uint8 = iota
	// FruitTypeOrange 水果类型：橘子
	FruitTypeOrange
	// FruitTypeBanana 水果类型：香蕉
	FruitTypeBanana
)

// produceFruitFuncs 生产水果方法集合
var produceFruitFuncs = map[uint8]func() Fruit{
	FruitTypeApple: func() Fruit {
		return Apple{}
	},
	FruitTypeOrange: func() Fruit {
		return orange{}
	},
	FruitTypeBanana: func() Fruit {
		return banana{}
	},
}

// ProduceFruit 生产水果
func ProduceFruit(fruitType uint8) Fruit {
	f, ok := produceFruitFuncs[fruitType]
	if !ok {
		return EmptyFruit{}
	}
	return f()
}

// Fruit 水果接口
type Fruit interface {
	// GetName 获取水果名称
	GetName() string
}

// EmptyFruit 空水果
type EmptyFruit struct{}

// GetName 获取名称
func (e EmptyFruit) GetName() string {
	return ""
}

// Apple 苹果
type Apple struct{}

// GetName 获取名称
func (a Apple) GetName() string {
	return "apple"
}

// orange 橘子
type orange struct{}

// GetName 获取名称
func (a orange) GetName() string {
	return "orange"
}

// banana 香蕉
type banana struct{}

// GetName 获取名称
func (a banana) GetName() string {
	return "banana"
}
