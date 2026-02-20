package main

import (
    "math/rand"
    "fmt"
    "time"
)

func main() {
    var array [][]byte
    for true {
        if len(array) < 100 {
            random := make([]byte, 5 * 1024 * 1024)
            rand.Read(random)
            array = append(array, random)
        }

	    fmt.Println("Hello world!")
	    time.Sleep(1 * time.Second)
	}
}