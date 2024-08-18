package main

import "fmt"

// 포인터를 리턴하는 함수
func returnPtr(x int) *int {
	y := x * x
	return &y // y 변수에 대한 메모리 주소를 리턴
}

func main() {
	sq := returnPtr(10)
	fmt.Println("sq:", *sq) // 100 출력
}
