# Udemy Go: The Complete Developer's Guide (Golang) 学習メモ

https://www.udemy.com/go-the-complete-developers-guide/learn/v4/overview

コース中のコード: https://github.com/StephenGrider/GoCasts

コーススライド: 　https://www.udemy.com/go-the-complete-developers-guide/learn/v4/t/lecture/7859182?start=0

##
* GoはOOではない
* レシーバーとして関数を定義
```
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```
  * d は他の言語のthisやself
* rangeの使わない返り値は_で受ける
* structはCの構造体っぽい
* ポインタが重要
  * Pass by Valueされる
  * レシーバーで受け取るのはコピーされた値
  * & 参照、 * ポインタ　（C++との違いはあるのか？）
  * *addressでaddressをvalueに変換
  * &valueでvalueをaddressに変換
  
```
func (pointerToPerson *person) updateName(newFirstName string) { // *はtype description(a pointer to a person)
	(*pointerToPerson).firstName = newFirstName // *はオペレータ、*によってポインタが指す値を得る。2つの*の意味が異なることに注意
}
```

* https://play.golang.org でテスト可能

* sliceはarrayへのポインタを内部構造で保持（pointer to head, capacity, lengthを管理）
  * sliceを引数でわたすと、sliceのデータ構造がコピーして渡される
  * s[0]などのインデックスはarray要素へのポインタ
```
func main() {
  mySlice := []string{"Hi", "There", "How", "Are", "You"}
  updateSlice(mySliece)
  fmt.Println(mySlice)
}

func updateSlice( s []string) {
  s[0] = "Bye"
}

--->

[Bye There How Are You]  // なぜ？
```

* value typeとreference type
  * Value Types（関数に渡すときにポインタに注意）: int, float, string, bool, structs
  * Reference Types（関数に渡すときにポインタ考慮は不要）: slices, maps, channels, pointers, functions