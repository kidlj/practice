package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type fileInfo struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Type    string `json:"type"`
	Size    int64  `json:"size"`
	ModTime int64  `json:"mtime"`
}

func main() {
	root := "file"
	entries, err := os.ReadDir(root)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("not found:", err)
		}

		if os.IsPermission(err) {
			log.Fatal("permission denied:", err)
		}

		log.Fatal(err)
	}

	files := make([]fileInfo, 0, len(entries))

	for _, e := range entries {
		fi, err := e.Info()
		if err != nil {
			continue
		}
		path := filepath.Join(root, fi.Name())
		if !fi.IsDir() {
			files = append(files, fileInfo{
				Name:    fi.Name(),
				Path:    path,
				Type:    "file",
				Size:    fi.Size(),
				ModTime: fi.ModTime().UnixNano() / 1e6,
			})
		} else {
			size := estimateDirSize(path)
			files = append(files, fileInfo{
				Name:    fi.Name(),
				Path:    path,
				Type:    "dir",
				Size:    size,
				ModTime: fi.ModTime().UnixNano() / 1e6,
			})
		}
	}

	for _, fi := range files {
		fmt.Printf("%s: %d\n", fi.Name, fi.Size)
	}
}

const maxDirSizeEstimated = 1 << 30 // 1GB

// Estimate directory size up to `maxDirSizeEstimated`, if exceeded, return early.
func estimateDirSize(dir string) (size int64) {
	// Safe to ignore error
	_ = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return filepath.SkipDir
		}
		fmt.Println(path, info.Size())
		size += info.Size()
		if size > maxDirSizeEstimated {
			return filepath.SkipAll
		}

		return nil
	})

	return
}
