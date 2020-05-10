package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		var size int64
		if info.Size() == 0 {
			fmt.Println(path, "Empty")
		} else {
			size = info.Size()
			fmt.Println(path, size, "Kb")
		}
		return nil
	})
}
