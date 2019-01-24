package fs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func scanFolder(dirname string) {
	var files []string

	root := "upload/iam"
	root = dirname
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

// ReadFilesInInputDir function is used to list the files in a given directory passed in by the user
func ReadFilesInInputDir(dir string) {
	finDir := dir + "/."
	file, err := os.Open(finDir)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println(name)
	}
}

// ReadFilesInCurrentDir function will read the current directory and print the files inside the current dir
func ReadFilesInCurrentDir() {
	file, err := os.Open("/.")
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	for _, name := range list {
		fmt.Println(name)
	}
}

// CountFilesInDirWithIoutil function is used to read and print the number of files in the given directory.
func CountFilesInDirWithIoutil(dir string) {
	files, _ := ioutil.ReadDir(dir)
	fmt.Printf("Number of files in directory=>   %d\n", len(files))
}
