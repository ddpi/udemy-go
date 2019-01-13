package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo // structの入れ子
	contactInfo // typeと同じ名前のフィールド名
}

func (pointerToPerson *person) updateName(newFirstName string) { // *はtype description(a pointer to a person)
	(*pointerToPerson).firstName = newFirstName // *はオペレータ、*によってポインタが指す値を得る。2つの*の意味が異なることに注意
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func main() {
	// alex := person{"Alex","Anderson"} // structの初期化法1: 定義順に初期化されるが、定義が変わった時に危険
	// alex := person{firstName: "Alex", lastName: "Anderson"} // structの初期化法2
	var alex person // structの初期化法3: 全てのフィールドがゼロ値(Zero Value, stringなら"")で初期化される
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo.email = "alex@example.com"
	alex.contactInfo.zipCode = 99999
	alex.print()
	// fmt.Println(alex)
	// fmt.Printf("%+v\n", alex) // %+v フィールド名も出力される

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		// contact: contactInfo{
		contactInfo: contactInfo{
			email:   "jim@example.com",
			zipCode: 94000, // 初期化で改行する時、structの最終行も,が必要
		},
	}
	jim.print()

	// jimPointer := &jim // jimへの参照（アドレス）を渡す
	// jimPointer.updateName("jimmy")
	jim.updateName("jimmy") // レシーバーがポインタの場合、goが参照に変換してくれる
	jim.print()
	// fmt.Println(jim)
	// fmt.Printf("%+v\n", jim) // %+v フィールド名も出力される

}
