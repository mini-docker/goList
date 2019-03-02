package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 10)

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("ch",i,len(ch))
		if(len(ch)>9){
			close(ch)
		}
		fmt.Println("ch",i,len(ch))
	}
	fmt.Println("ch",ch)

	//for {
	//	var b int
	//	b = <-ch
	//	if ok == false {
	//		fmt.Println("chan is close")
	//		break
	//	}
	//	fmt.Println(b)
	//}
}
