package main

import (
	"fmt"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1() called")
	wg.Done()
}

func f2() {
	fmt.Println("f2() called")
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	// fmt.Println("No of cpus", runtime.NumCPU())
	// fmt.Println("No of go routines", runtime.NumGoroutine())
	// fmt.Println("OS:", runtime.GOOS)
	// fmt.Println("Arch", runtime.GOARCH)

	// fmt.Println("GOMAXPROS", runtime.GOMAXPROCS(0))

	go f1(&wg)

	//fmt.Println("No of go routines after f1()", runtime.NumGoroutine())
	wg.Wait()
	f2()

}
