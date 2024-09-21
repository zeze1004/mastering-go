package main

import "fmt"

func main() {
	aString := "Hello, World! ✋"
	fmt.Println("첫번째 글자:", aString[0])

	// 룬타입
	// 하나의 룬 문자
	r := '✋'
	fmt.Println("룬타입:", r)

	// 룬 문자를 텍스트로 변환
	fmt.Printf("%s를 char로 변환: %c\n", r, r)

	// 룬 슬라이스로 텍스트를 출력하는 법
	// 문자열을 rune 슬라이스로 변환
	runeSlice := []rune(aString)
	lastChar := runeSlice[len(runeSlice)-1]
	fmt.Printf("타입: %T\n", lastChar)
	fmt.Printf("마지막 글자: %c\n", lastChar)

	// 문자를 룬 현태로 출력
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
}
