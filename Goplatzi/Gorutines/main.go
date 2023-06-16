package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("world")
	wg.Add(1)
	go say("hello", &wg)
	wg.Wait()

	//time.Sleep(time.Second * 1)

}
