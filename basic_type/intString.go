package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 44032 // '가'의 코드포인트는 U+AC00(16진수), (10진수로 44032)
	// strconv 패키지로 정수를 문자열로 변환
	input := strconv.Itoa(n) // 정수 값을 해당 정수 값고 동일한 문자열로 변환
	fmt.Printf("strconv.Itoa(): %s 의 type %T\n", input, input)

	input = strconv.FormatInt(int64(n), 10) // 정수 값을 해당 정수 값고 동일한 문자열로 변환
	fmt.Printf("strconv.FormatInt(): %s 의 type %T\n", input, input)

	// string()의로 정수를 변환
	input = string(n)
	fmt.Printf("string(): %s 의 type %T\n", input, input)
}
