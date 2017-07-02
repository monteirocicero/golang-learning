package main

import (
	"fmt"
	"os"
)

func createFiles(dirBase string, files ...string)  {
	for _, name := range files {
		dirFile := fmt.Sprintf("%s/%s.%s", dirBase, name, "txt")

		arq, err := os.Create(dirFile)

		defer arq.Close()

		if err != nil {
			fmt.Print("Error the create file %s: %v\n", name, err)
			os.Exit(1)
		}

		fmt.Printf("File %s created.\n", arq.Name())
	}
}

func main() {
	tmp := os.TempDir()

	createFiles(tmp)
	createFiles(tmp, "test1")
	createFiles(tmp, "test2", "test3", "test4")
}
