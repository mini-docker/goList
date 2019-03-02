package main

import (
	"runtime"
	"time"
	"fmt"
)

//一个携程能跑完的，就用一个。不必用其他的，故为同步执行
func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)
	for i := 0; i < 10240; i++ {
		fmt.Println(i,"vvviii")
		go func() {
			for {
				select {
				case <-time.After(time.Second):
					fmt.Println("after")
				}
			}
		}()
	}

	time.Sleep(time.Second * 3)
}
