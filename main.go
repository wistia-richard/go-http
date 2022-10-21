package main

import (
	"fmt"
	"os"

	"github.com/wistia-richard/test-package/environ"
)

func main() {
	// create a file
	test_file, err := os.Create("test")
	defer test_file.Close()

	if err != nil {
		fmt.Println("Couldnt create the file")
		os.Exit(1)
	}

	// write to the file
	test_file.WriteString("Hello Richard")

	//read data
	data, _ := os.ReadFile("test")

	// write to stdout
	fmt.Println(string(data))

	// open file
	test1, err := os.OpenFile("test1", os.O_RDWR, 0755)
	defer test1.Close()

	if err != nil {
		fmt.Println("The file named 'test1' doesn't exist")
		os.Exit(1)
	}

	// write to file
	data1 := make([]byte, 100)
	data1 = []byte("Hello Richard again")

	fmt.Println(data1)
	_, err1 := test1.Write(data1)

	if err1 != nil {
		fmt.Println(err1)
	}
	//read data
	data_read, _ := os.ReadFile("test1")

	// write to stdout
	fmt.Println(string(data_read))

	// get environment variables
	for _, i := range environ.Env() {
		fmt.Println(i)
	}

}
