package main

import (
	"fmt"
	"github.com/inocencio/ffile/config"
	. "github.com/inocencio/ffile/systemapp"
	"strconv"
)

func main() {
	config.SetupConfigFiles()

	var fullpath = "/home/nihil/teste Arquivo Interno Benjamin Linx.csv"
	fmt.Println("Original: ", fullpath)

	var c = FReadFile(fullpath, false)
	fmt.Println(*c)

	var sc, f = FReadFileScanner(fullpath, true)
	defer CloseFile(f, true)
	var counter = 0

	for sc.Scan() {
		counter += 1
		fmt.Println("line[" + strconv.Itoa(counter) + "]: " + sc.Text())
	}
}
