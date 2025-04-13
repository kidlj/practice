package main

import (
	"fmt"
	"os"
)

func main() {
	size := sumSize(".")
	fmt.Printf("dir total size: %d\n", size)
}

func sumSize(path string) (size int64) {
	file, err := os.Stat(path)
	if err != nil {
		return
	}
	if !file.IsDir() {
		return file.Size()
	}
	size += file.Size()
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("read dir error:", err)
		return
	}
	for _, f := range files {
		size += sumSize(path + "/" + f.Name())
	}

	return size
}
