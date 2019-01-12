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
* 