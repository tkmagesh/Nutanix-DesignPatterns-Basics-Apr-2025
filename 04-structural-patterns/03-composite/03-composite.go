package main

import "fmt"

type FileSystemComponent interface {
	Show(indentation string)
}

type File struct {
	Name string
}

func (f *File) Show(indentation string) {
	fmt.Println(indentation + f.Name)
}

type Folder struct {
	Name     string
	Children []FileSystemComponent
}

func (f *Folder) Add(child FileSystemComponent) {
	f.Children = append(f.Children, child)
}

func (f *Folder) Show(indentation string) {
	fmt.Println(indentation + f.Name + "/")
	for _, child := range f.Children {
		child.Show(indentation + "  ")
	}
}

func main() {
	root := &Folder{Name: "root"}
	home := &Folder{Name: "home"}
	user := &Folder{Name: "user"}
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}

	user.Add(file1)
	home.Add(user)
	root.Add(home)
	root.Add(file2)

	root.Show("")
}
