package fs

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

func GetHome() string {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

// isError function is a helper function for error handeling for all crud functions
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func CreateAndWriteFile(str, path string) {
	CreateFile(path)
	WriteFile(str, path)
}

// createFile function is used to create a file given the save path
func CreateFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file ", path)
}

// writeFile funtion is used to write to a pre-existing file, opening the file and writing the passed string into the file
func WriteFile(str, path string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(str)
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> done writing to file " + path)
}

func writeExecutableFile(str, path string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0777)
	if isError(err) {
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(str)
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> done writing to file ", path)
}

// ReadFile function is used read a given file
func ReadFile(path string) {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

	fmt.Println("==> done reading from " + path)
	fmt.Println(string(text))
}

// DeleteFile function is used to delete a given file
func DeleteFile(path string) {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("==> done deleting " + path)
}

func appendFile(path, text string) {
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

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}
