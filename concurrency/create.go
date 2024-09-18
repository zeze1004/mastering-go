package main

import (
	"fmt"
	"time"
)

func main() {
	go func(x int) {
		fmt.Printf("x: %d\n", x)
	}(10)

	go printMe("ODEE")

	time.Sleep(time.Second)
	fmt.Println("프로그램이 종료됐습니다.")
}

func printMe(x string) {
	fmt.Println(x)
	return
}
