package main

import "fmt"

// 함수를 리턴하는 함수
func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	// 같은 함수를 호출했지만, 서로 독립적으로 실행됨
	i := funReturnFun()
	j := funReturnFun()

	fmt.Println("i-1: ", i()) // 1 출력
	fmt.Println("i-2: ", i()) // 4 출력
	fmt.Println("j-3: ", j()) // 1 출력
	fmt.Println("j-1: ", j()) // 4 출력
	fmt.Println("i-3: ", i()) // 9 출력
}
