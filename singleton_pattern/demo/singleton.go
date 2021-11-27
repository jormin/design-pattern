package demo

import "sync"

var instance *singletonCounter

// singletonCounter 单例计数器
type singletonCounter struct {
	num int
}

// Count 计数
func (s *singletonCounter) Count() {
	s.num++
}

// GetNum 获取计数器当前数量
func (s *singletonCounter) GetNum() int {
	return s.num
}

// GetSingletonCounter 获取单例计数器
func GetSingletonCounter() *singletonCounter {
	var once sync.Once
	once.Do(
		func() {
			instance = &singletonCounter{
				num: 0,
			}
		},
	)
	return instance
}
