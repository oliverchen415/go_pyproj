package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"path/filepath"
)

var (
	newFile *os.File
	err error
)

func createFile(name string) {
	newFile, err = os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	newFile.Close()
}

func main() {
	projName := flag.String("name", "dummyProj", "your project name")
	fileType := flag.String("type", "py", "file type")

	flag.Parse()
	fileName := fmt.Sprintf("%s.%s", *projName, *fileType)
	dirName := filepath.Join(".", *projName)

	_, err := os.Stat(dirName)
	if !os.IsNotExist(err) {
		log.Fatal("Directory with the same name exists")
	}
	os.Mkdir(dirName, 0700)
	os.Chdir(dirName)

	createFile(fileName)
	createFile("README.md")
	createFile(".gitignore")

	fmt.Println("Project Name:", *projName)
	fmt.Println("Files: README.md, .gitignore, and", fileName, "created.")
}