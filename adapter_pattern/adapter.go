package adapter_pattern

import "fmt"

// OldInterface 旧接口
type OldInterface interface {
	// SaveToDB 存储入库
	SaveToDB() (bool, error)
}

// NewInterface 新接口
type NewInterface interface {
	// SaveData 存储数据
	SaveData() (bool, error)
}

// Adaptee 适配者
type Adaptee struct{}

// SaveToDB 存储入库
func (a Adaptee) SaveToDB() (bool, error) {
	fmt.Println("save to db success")
	return true, nil
}

// Adapter 适配器
type Adapter struct {
	oldInterface OldInterface
}

// SaveData 存储数据
func (a Adapter) SaveData() (bool, error) {
	return a.oldInterface.SaveToDB()
}
