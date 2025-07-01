package main

import (
	"fmt"
	"os"
)

// ListFolders lists all folders in the given directory path.
func listFolders(path string) ([]string, error) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return nil, err
	}

	var folders []string

	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}

	return folders, nil
}

func main() {
	folders, err := listFolders("C:/Users/delse/Videos/Video Template")

	if err != nil {
		return
	}

	for _, folder := range folders {
		fmt.Println(folder)
	}
}
