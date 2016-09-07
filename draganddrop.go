package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please drag a file into the exe")
		os.Exit(1)
	} else {
		fmt.Println(os.Args[1])
		time.Sleep(time.Second * 10)
	}
}
