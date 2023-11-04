package main

import (
	"fmt"
	"os"
)

func main() {
	tarFilePath := "/Users/ashishkumarsingh/Desktop/tidbits/learning_go" + "/imgpkg-tar-image"
	tmpFile, err := os.Create(tarFilePath)
	defer tmpFile.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println(tmpFile.Name())
	fmt.Println(tmpFile)
}
