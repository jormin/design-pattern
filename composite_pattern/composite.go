package composite_pattern

import "fmt"

// Component 组件
type Component interface {
	// Search 搜索
	Search(string)
}

// File 文件
type File struct {
	Name string `json:"name"`
}

// Search 搜索
func (f *File) Search(s string) {
	fmt.Printf("在文件[%s]中搜索[%s]\n", f.Name, s)
}

// Folder 文件夹
type Folder struct {
	Components []Component `json:"components"`
	Name       string      `json:"name"`
}

// Search 搜索
func (f *Folder) Search(s string) {
	fmt.Printf("在文件夹[%s]中搜索[%s]\n", f.Name, s)
	for _, v := range f.Components {
		v.Search(s)
	}
}

// Add 添加组件
func (f *Folder) Add(component Component) {
	f.Components = append(f.Components, component)
}
