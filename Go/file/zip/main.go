package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// 目录路径
	sourceDir := "/home/lj7671/Developer/practice/go/containers"
	// ZIP文件名
	zipFileName := "output.zip"

	// 创建ZIP文件
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	// 创建ZIP写入器
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历目录并将文件和文件夹添加到ZIP包中
	err = filepath.Walk(sourceDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// 修改文件头中的文件名，以保留相对路径
		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}
		header.Name = relPath

		// 如果是目录，需要额外添加一个文件夹的Header
		if info.IsDir() {
			header.Name += string(filepath.Separator)
		}

		// 创建ZIP文件的文件头
		file, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// 如果是文件，将文件内容复制到ZIP文件
		if !info.IsDir() {
			sourceFile, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			_, err = io.Copy(file, sourceFile)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Directory '%s' has been zipped to '%s'\n", sourceDir, zipFileName)
}
