package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	n := 0

	for i := 0; i < gr; i++ {

		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()

	}

	wg.Wait()

	fmt.Println("Final Value of n", n)

}
