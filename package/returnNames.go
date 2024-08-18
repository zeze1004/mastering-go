package main

import (
	"fmt"
	"os"
	"strconv"
)

// Go에서는 함수의 리턴 값에도 이름을 붙일 수 있음(네임드 리턴 값)

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return // 네임드 리턴 값은 return문에 변수를 명시하지 않아도 됨
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return min, max // 명시해도 됨
}

func main() {
	argumets := os.Args
	if len(argumets) < 3 {
		fmt.Println("The program needs at least 2 arguments!")
		return
	}

	a1, _ := strconv.Atoi(argumets[1]) // 문자열을 정수로 변환
	a2, _ := strconv.Atoi(argumets[2])

	fmt.Println(minMax(a1, a2))
	min, max := minMax(a1, a2)
	fmt.Println(min, max)

	// 네임드 리턴 값은 함수 내부에서 선언된 변수와 동일한 이름을 사용하므로 함수 내부에서 변수를 선언하지 않아도 됨
	fmt.Println(namedMinMax(a1, a2))
	min, max = namedMinMax(a1, a2)
	fmt.Println(min, max)
}
