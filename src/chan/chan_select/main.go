package main

import "fmt"
import "time"

func main() {
	var ch chan int
	ch = make(chan int, 10)
	ch2 := make(chan int, 10)
	go func() {
		var i int
		for {
			ch <- i
			time.Sleep(time.Second)
			ch2 <- i * i
			fmt.Println("ch2执行")
			time.Sleep(time.Second)
			i++
			fmt.Println("zuihou")
		}
	}()
	for {
		select {
		case v := <-ch:
			fmt.Println(v,"ch")
		case v := <-ch2:
			fmt.Println(v,"ch2")
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
//0 ch
//0 ch2
//ch2执行
//zuihou
//1 ch
//get data timeout
//ch2执行
//zuihou
//2 ch
//1 ch2
//get data timeout
//ch2执行
//4 ch2
//zuihou
//3 ch
//get data timeout
//ch2执行
//9 ch2
//zuihou

// 在timeout过程中ch的传输被忽略