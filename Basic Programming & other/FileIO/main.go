package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// file, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE, 0664)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer file.Close()

	// bufferWriter := bufio.NewWriter(file)

	// bs := []byte{1, 2, 3, 4}

	// byteWriter, err := bufferWriter.Write(bs)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("Byte written %d\n", byteWriter)

	// bytesAvai := bufferWriter.Available()
	// log.Printf("Byte Available %d\n", bytesAvai)

	// byteString, err := bufferWriter.WriteString("\nRandom String")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _ = byteString

	// bufferWriter.Flush()

	// file, err := os.Open("a.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bs := make([]byte, 2)

	// numberOfByte, err := io.ReadFull(file, bs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%d", numberOfByte)
	// log.Printf("%s", bs)

	// file, err = os.Open("main.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// log.Printf("%s", data)
	// log.Printf("%d", len(data))

	file, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	status := scanner.Scan()

	if status == false {

		if scanner.Err() == nil {
			log.Println("Scan Complete")
		} else {
			log.Fatal(scanner.Err())
		}

	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
