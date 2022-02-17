package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func checkUrl(url string, ch chan string) {

	resp, err := http.Get(url)

	if err != nil {

		s := fmt.Sprintf("%s is down", url)
		s += fmt.Sprintf("Error:%v", err)

		fmt.Println(s)

		ch <- url

	} else {

		//defer resp.Body.Close()

		if resp.StatusCode == 200 {

			// respBodyByte, _ := ioutil.ReadAll(resp.Body)

			// file := strings.Split(url, "//")[1]

			// file += ".txt"

			//s := fmt.Sprintf("Writing data to %s", file)

			// err = ioutil.WriteFile(file, respBodyByte, 0664)

			// if err != nil {
			// 	s += "Error while writing file"
			// 	ch <- s
			// }
			s := fmt.Sprintf("%s is up", url)
			fmt.Println(s)
			ch <- url

		}

	}

	// wg.Done()

}

func main() {

	urls := []string{"https://www.google.com", "https://www.golang.org", "https://www.facebook.com"}

	// var wg sync.WaitGroup

	// wg.Add(len(urls))

	ch := make(chan string)

	for _, url := range urls {
		go checkUrl(url, ch)
		//fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of GOROUTINES", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {

	// 	fmt.Println(<-ch)

	// }

	// for {
	// 	fmt.Println(strings.Repeat("#", 20))
	// 	go checkUrl(<-ch, ch)
	// 	time.Sleep(time.Second * 2)
	// }

	// for url := range ch {
	// 	time.Sleep(time.Second * 2)
	// 	go checkUrl(url, ch)
	// }

	for url := range ch {
		go func(u string) {
			time.Sleep(time.Second * 2)
			checkUrl(u, ch)
		}(url)

	}

	// wg.Wait()

}
