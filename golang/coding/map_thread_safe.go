package main

import (
	"fmt"
	"sync"
)

type UserAges struct {
	ages map[string]int
	//读写锁RWMutex 读写锁是一种更加细粒度的锁，它可以分别控制读和写的操作，这样就可以大大提高程序的并发性
	//互斥锁是一种独占锁，同一时间只有一个协程可以获取到锁，其他协程只能阻塞等待锁的释放
	sync.RWMutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return 0
}

/**
 * map属于引用类型，并发读写时多个协程间是通过指针访问同一个地址，即访问共享变量，此时同时读写资源存在竞争关系；
 * 为了保证并发安全，需要加锁，Go语言提供了sync包中的互斥锁sync.Mutex和读写锁sync.RWMutex，前者是独占锁，后者是共享锁；
 * 如果没有ua.RLock()和ua.RUnlock()，则会报错：fatal error: concurrent map read and map write
 */
func main() {
	count := 10
	// 开启协程等待
	wg := sync.WaitGroup{}
	wg.Add(count * 3)
	// 开启10个协程，每个协程对map进行写操作
	// 必须对ages进行初始化，否则会panic
	u := UserAges{ages: make(map[string]int)}
	add := func(i int) {
		defer wg.Done()
		u.Add(fmt.Sprintf("user_%d", i), i)
	}
	// 使用goroutine添加数据
	for i := 0; i < count; i++ {
		go add(i)
		go add(i + 1)
	}

	get := func(i int) {
		defer wg.Done()
		fmt.Println(u.Get(fmt.Sprintf("user_%d", i)))
	}
	// 开启10个协程，每个协程对map进行读操作
	for i := 0; i < count; i++ {
		go get(i)
		//go func(i int) {
		//	defer wg.Done()
		//	fmt.Println(u.Get(fmt.Sprintf("user_%d", i)))
		//}(i)
	}
	// 阻塞代码，等待所有协程执行完毕
	wg.Wait()
	fmt.Println("Done")
}
