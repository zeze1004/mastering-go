package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(x int) {
			fmt.Printf("고루틴 개수: %d\n", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Printf("프로그램이 종료됐습니다.")
}
