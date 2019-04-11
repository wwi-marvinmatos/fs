package fs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getHome() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

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

// ReadDirWithIoutilReturn function
func ReadDirWithIoutilReturn(dir string) int {
	files, _ := ioutil.ReadDir(dir)
	numFiles := len(files)
	// Surpressing the output, i will comment out, later if i add a verbose flag, ill have this printed
	// fmt.Printf("Number of files to upload=>   %d\n", numFiles)
	return numFiles
}

// ReadCurrentDir function
func ReadCurrentDir(dir string) {
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

// CountFilesInDirWithIoutil function is used to read and print the number of files in the given directory.
func CountFilesInDirWithIoutil(dir string) {
	files, _ := ioutil.ReadDir(dir)
	fmt.Printf("Number of files in directory=>   %d\n", len(files))
}

func AppendFile(path, text string) {
	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text + "\n"); err != nil {
		log.Println(err)
	}
}

// CreateAndWriteFile funtion is a helper for creating a file and writing to it
func CreateAndWriteFile(str, path string) {
	CreateFile(path)
	WriteFile(str, path)
}
