package main

import "fmt"

// 포인터를 매개변수로 받는 함수
func getPtr(v *float64) float64 {
	// 포인터를 통해 값을 변경
	*v = 13.0
	return *v * *v
}

func main() {
	x := 12.2
	fmt.Println(getPtr(&x)) // x의 제곱 출력
	fmt.Println(x)
	x = 12
	fmt.Println(getPtr(&x))
}
