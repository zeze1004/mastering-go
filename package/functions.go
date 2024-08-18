package main

import (
	"fmt"
	"os"
	"strconv"
)

// 익명함수(클로저)
// Go언어의 함수는 여러 개의 값을 동시에 리턴 가능
// 익명함수를 사용하면 함수로부터 리턴되는 여러 값에 접근하기 위해 별도로 구조체를 정의하지 않아도 돼서 편리

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}

func main() {
	arguments := os.Args // 명령행 인수를 읽어옴
	if len(arguments) != 2 {
		fmt.Println("The program needs 1 argument!")
		return
	}

	num, err := strconv.Atoi(arguments[1]) // 문자열을 정수로 변환
	if err != nil {
		fmt.Println(err)
		return
	}

	// square와 double 함수를 익명함수로 구현
	// 각 변수는 각 익명함수를 할당받을 수 있음, 변수를 서로 바꿔쓰면 다른 익명함수가 호출돼서 익명함수에는 기존 변수를 수정하면 안됨
	square := func(s int) int { // 익명함수
		return s * s
	}
	fmt.Println("The square of", num, "is", square(num))

	double := func(s int) int { // 익명함수
		return s + s
	}
	fmt.Println("The double of", num, "is", double(num))

	fmt.Println(doubleSquare(num))
	d, s := doubleSquare(num)
	fmt.Println(d, s)
}
