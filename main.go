package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 0; i < num; i++ {
		fmt.Printf("%d: ", i)
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		fmt.Print("\n")
	}
}
