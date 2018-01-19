package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	//	fmt.Println("MaxProCs:\t", runtime.GOMAXPROCS(n))
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("")
	fmt.Println("After goroutins call")
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println("")
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
