package main
import (
	"fmt"
	"sync"
	"time"
)
// Mutex 是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。
// Mutex 类型的锁和线程无关，可以由不同的线程加锁和解锁。

// Lock 方法锁住 m，如果 m 已经加锁，则阻塞直到 m 解锁。类似 js的promise then方法
// func (m *Mutex) Lock()

// Unlock 方法解锁 m，如果 m 未加锁会导致运行时错误。 类似 js的promise catch方法
// func (m *Mutex) Unlock()

//注意
//在一个 goroutine 获得 Mutex 后，其他 goroutine 只能等到这个 goroutine 释放该 Mutex
//使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
//在 Lock() 之前使用 Unlock() 会导致 panic 异常
//已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
//在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
//适用于读写不确定，并且只有一个读或者写的场景


func main() {
	var mutex sync.Mutex
	wait := sync.WaitGroup{}
	fmt.Println("Locked")
	mutex.Lock()
	fmt.Println("lock after")
	for i := 1; i <= 3; i++ {
		wait.Add(1)
		go func(i int) {
			fmt.Println("Not lock:", i)
			mutex.Lock()
			fmt.Println("Lock:", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock:", i)
			mutex.Unlock()
			defer wait.Done()
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlocked")
	mutex.Unlock()
	fmt.Println("waite before")
	wait.Wait() //等到循环完成，wait才执行 done
	fmt.Println("waited")
}

//Locked
//Not lock: 2
//Not lock: 1
//Not lock: 3
//Unlocked
//Lock: 2
//Unlock: 2
//Lock: 1
//Unlock: 1
//Lock: 3
//Unlock: 3