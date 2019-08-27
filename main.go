package main

import (
	"fmt"
	"ginkgo-talk/fibonacci"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Failed to read: " + err.Error())
	}
	fibo, err := fibonacci.F(n)
	if err != nil {
		panic("Failed to get Fibonacci sequence: " + err.Error())
	}
	fmt.Println(fibo)
}
