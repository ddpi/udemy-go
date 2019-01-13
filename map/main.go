package main

import "fmt"

func main() {
	// var colors map[string]string // 初期化法(2) 空のMap

	// colors := make(map[string]string) // 初期化法(3) 空のMap

	colors := map[string]string{ // 初期化法(1)
		"red":   "#ff0000",
		"green": "#4bf745",
		"black": "#000000",
	}

	colors["white"] = "#ffffff"     // 要素追加
	colors["red"] = "value changed" // 要素変更
	delete(colors, "green")         // 要素削除

	// fmt.Println(colors)
	printMap(colors)

}

// colors(main内のmap[string]string)を順次プリント
func printMap(c map[string]string) {
	for color, hex := range c { // colorにキーが、hexにvalueが
		fmt.Println(color, " is ", hex)
	}
}
