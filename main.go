package main

import (
	"flag"
	"log"
	"fmt"
	"os"
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

	flag.Parse()
	fileName := fmt.Sprintf("%s.py", *projName)
	dirName := fmt.Sprintf("./%s", *projName)

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