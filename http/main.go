package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// fmt.Println(resp.Status)

	/**
		bs := make([]byte, 999999) // 999999の要素をもつbyte sliceを作成
		resp.Body.Read(bs)
		fmt.Println(string(bs)) // byte sliceをstringにして出力
	**/

	// io.Copy(os.Stdout, resp.Body) // resp.Body.Readして出力と同等 (os.StdoutはWriter interfaceを実装)
	lw := logWriter{}
	io.Copy(lw, resp.Body) // resp.Body.Readして出力と同等 (os.StdoutはWriter interfaceを実装)

}

// カスタムライター io.Copyの第一引数となるWriter(Write)の実装
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
