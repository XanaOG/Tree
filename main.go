package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Excluded directories, add the folder then true to exclude
var excludedDirs = map[string]bool{
	"node_modules": true,
	".svelte-kit":  true,
}

func PrintTree(root string, indent string) {
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	for i, file := range files {
		if file.IsDir() && excludedDirs[file.Name()] {
			continue
		}
		isLast := i == len(files)-1
		if isLast {
			fmt.Printf("%s└── %s\n", indent, file.Name())
		} else {
			fmt.Printf("%s├── %s\n", indent, file.Name())
		}
		if file.IsDir() {
			newRoot := filepath.Join(root, file.Name())
			newIndent := indent
			if isLast {
				newIndent += "    "
			} else {
				newIndent += "│   "
			}
			PrintTree(newRoot, newIndent)
		}
	}
}

func main() {
	fmt.Println(".")
	PrintTree(".", "")
}
