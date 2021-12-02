package composite_pattern

import "testing"

// Test 测试组合模式
func Test(t *testing.T) {
	file1 := &File{Name: "file1"}
	file2 := &File{Name: "file2"}
	file3 := &File{Name: "file3"}

	folder1 := &Folder{
		Name: "folder1",
	}
	folder2 := &Folder{
		Name: "folder2",
	}

	folder1.Add(file1)

	folder2.Add(folder1)
	folder2.Add(file2)
	folder2.Add(file3)

	folder2.Search("test")
}
