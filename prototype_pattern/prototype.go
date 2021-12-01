package prototype_pattern

// Student 学生
type Student struct {
	Name   string `json:"name" desc:"姓名"`
	Age    int    `json:"age" desc:"年龄"`
	Gender string `json:"gender" desc:"性别"`
	Class  *Class `json:"class" desc:"班级"`
}

// Clone 克隆学生-浅克隆
func (s *Student) Clone() *Student {
	stu := s
	stu.Class = s.Class.Clone()
	return stu
}

// DeepClone 克隆学生-深克隆
func (s *Student) DeepClone() *Student {
	stu := *s
	stu.Class = s.Class.DeepClone()
	return &stu
}

// Class 班级
type Class struct {
	Name string `json:"name" desc:"名称"`
}

// Clone 克隆班级-浅克隆
func (c *Class) Clone() *Class {
	class := c
	return class
}

// DeepClone 克隆班级-深克隆
func (c *Class) DeepClone() *Class {
	class := *c
	return &class
}
