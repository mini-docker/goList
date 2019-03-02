package main

import (
	"fmt"
	"sync"
	"time"
)
//Mutex 是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。
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

var (
	m    = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	fmt.Println(t.n, sum)
	lock.Lock()
	m[t.n] = sum
	//把上一步执行完成后解锁
	lock.Unlock()
}

func main() {
	for i := 0; i < 16; i++ {
		t := &task{n: i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	lock.Lock()
	fmt.Println("unlock before")
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	//把上一步执行完成后解锁
	lock.Unlock()
	fmt.Println("unlocked")

}
