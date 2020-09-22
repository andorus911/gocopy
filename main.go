package main

import (
	"flag"
	"fmt"
)

var fromFilePath string
var toFilePath string
var sourceOffset int64
var byteLimit int64

func init() {
	flag.StringVar(&fromFilePath, "from", "", "source file, from where copy information")
	flag.StringVar(&toFilePath, "to", "", "file where you want to copy information")
	flag.Int64Var(&sourceOffset, "offset", 0, "source offset")
	flag.Int64Var(&byteLimit, "limit", -1, "number of bytes to copy")
	flag.Parse()
}

func main() {
	if fromFilePath == "" || toFilePath == "" {
		fmt.Println("WARNING: The util needs both file paths. Use -from and -to flags.")
		return
	}
	err := Copy(fromFilePath, toFilePath, sourceOffset, byteLimit)
	if err != nil {
		fmt.Println(err.Error())
	}
}
