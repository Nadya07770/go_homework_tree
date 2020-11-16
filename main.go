package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func getFileSize(info os.FileInfo) string {
	var sizeResult string

	size := info.Size()
	if size == 0 {
		sizeResult = "empty"
	} else {
		sizeResult = strconv.FormatInt(size, 10) + "b"
	}

	return " (" + sizeResult + ")"
}

func printTree(out io.Writer, path string, printFiles bool, prefix string) error{
	fork_string := "├───"
	corner_string := "└───"
	wall_string := "│\t"
	space_string := "\t"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}

	var directories []os.FileInfo
	// separate files and directories
	for _, file := range files {
		if !printFiles && !file.IsDir() {
			continue
		}
		directories = append(directories, file)
	}

	for i := 0; i < len(directories); i++ {

		if !printFiles && !directories[i].IsDir() {
			continue
		}

		fileName := directories[i].Name()

		if !directories[i].IsDir() {
			fileName = fileName + getFileSize(directories[i])
		}

		if i != len(directories) - 1 {
			fmt.Fprintln(out, prefix + fork_string + fileName)
		} else {
			fmt.Fprintln(out, prefix + corner_string + fileName)
		}

		currentPrefix := prefix

		if i != len(directories) - 1 {
			currentPrefix = currentPrefix + wall_string
		} else {
			currentPrefix = currentPrefix + space_string
		}

		if directories[i].IsDir() {
			childPath := path + string(os.PathSeparator) + fileName
			printTree(out, childPath, printFiles, currentPrefix)
		}
	}

	return nil
}


func dirTree(out io.Writer, path string, printFiles bool) error {
	prefix := ""

	printTree(out, path, printFiles, prefix)
	return nil
}

func main() {
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(os.Stdout, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
