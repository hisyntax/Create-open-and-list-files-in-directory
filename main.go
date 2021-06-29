package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	//Create the file
	file, err := os.Create("create.pdf")

	if err != nil {
		// handle the error
		return
	}

	//Close the file
	defer file.Close()

	//Write data into the created file

	file.WriteString("This file is created using Go")

	//Read the file
	readFile, err := ioutil.ReadFile("create.pdf")
	if err != nil {
		// handle the error
		return
	}
	// Convert to strings
	str := string(readFile)

	// Display the file

	fmt.Println(str)

	//List all the files in the current working directory

	dir, err := os.Open(".")

	if err != nil {
		// handle the error
		return
	}

	//Close the file
	defer dir.Close()

	//List all the files in the working directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		//handle the error
		return
	}
	//Run a for loop to get all the files in the directory
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
