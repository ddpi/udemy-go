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

## Map
  * keyは全て同じタイプ
  * valueも全て同じタイプ（keyと異なって良い）

* Mapとstructの違い
  * Map
    * キーの型は全て同じ
    * バリューの型も全て同じ
    * キーはインデックスされ、イテレート可能
    * リファレンスタイプ（関数に参照で渡される）
    * 関連する属性の集合を表現
    * コンパイル時に全てのキーが決まっていなくて良い（実行時追加削除変更可能）

  * Struct
    * バリューの型は異なって良い
    * キーはインデックスされない
    * バリュータイプ（関数に値渡し）
    * コンパイル時にフィールドは決まっていなければならない
    * 様々な属性を持つ"thing"を表現

## Interfaces

* Go: 型付言語 → 様々な型に適用可能な関数は？
* 引数が異なっていても、同名の関数は作れない
* interfaceに関数を宣言。その関数を実装するstructはinterfaceのメンバーとなる
* interfaceのメンバーになるのは関数のシグネチャ（引数、返り値）まで同じでなければならない
* Concrete TypeとInterface Type
  * Concrete Type: 値を作成できる (map, struct, int, string,,, eg. englishBot)
  * Interface Type: 値を作成できない (eg. bot)
* interfaceはgeneric typeではない（Javaとは違う）
* interfaceは暗黙的である（自動的に判断される）
* interfaceはcontractであり、型管理を助ける（セマンティクスを実装者が管理しなければならない）
  

### https://golang.org/pkg/io/#Reader
* Reader interface 様々な入力を[]byteとして出力し、汎用的に扱えるようにする