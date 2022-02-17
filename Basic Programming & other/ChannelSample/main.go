package main

import "fmt"

func factorial(n int, ch chan int) {

	f1 := 1

	for i := 2; i <= n; i++ {
		f1 *= i
	}

	ch <- f1

}

func main() {

	ch := make(chan int)

	defer close(ch)

	// go factorial(4, ch)

	// n := <-ch

	// fmt.Println(n)

	for i := 1; i <= 10; i++ {

		go factorial(i, ch)
		n := <-ch

		fmt.Println(n)

	}

}
