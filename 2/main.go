package main  

import (
	"fmt"
)

func main(){
	
	file1 := &File{name: "File 1"}
	file2 := &File{name: "File 2"}
	file3 := &File{name: "File 3"}

	folder1 := &Folder{
		children: []Inode{file1},
		name: "Folder 1",
	}

	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name: "Folder 2",
	}
 
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.print("   ")


	cloneFolder := folder2.clone()

	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.print("   ") 
}