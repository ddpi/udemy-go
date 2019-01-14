package main

import (
	"fmt"
	"io"
	"os"
)

// 第一引数で指定されたファイルを標準出力に表示
func main() {
	fname := os.Args[1]
	// fmt.Println(fname)
	f, err := os.Open(fname)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)

}
