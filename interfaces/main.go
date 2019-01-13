package main

import "fmt"

type bot interface {
	getGreeting() string // "getGreeting() string"を実装するstructはbotとなる
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb) // getGreeting() stringを実装しているので、botを引数とするprintGreetingを呼べる
	printGreeting(sb)

}

/**
// 異なる引数の同名関数は作れない
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
**/

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string { // レシーバーを参照しない場合は省略可能（すべき？）
	return "Hello!"
}

// func (sb spanishBot) getGreeting() string {
func (spanishBot) getGreeting() string {
	return "Hola!"
}
