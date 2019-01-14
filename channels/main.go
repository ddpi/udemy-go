package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverlow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// fmt.Println(link)
		go checkLink(link, c) // 新しいGo Routineを生成（並行処理）
	}

	for l := range c {
		// time.Sleep(5 * time.Second) // main routine内でsleepすると他の処理もブロックされる
		// go checkLink(l, c)
		go func(link string) { // 匿名の関数リテラルによって子routine内でスリープさせる。値はコピー渡し
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link+"might be down!", err)
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
