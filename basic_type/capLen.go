package main

import "fmt"

func main() {
	// 길이만 정의했으므로, 길이 = 용량
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a)) // 4, 4

	// 슬라이스를 초기화하는 경우도 길이 = 용량
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b)) // 5, 5

	// 길이와 용량이 같은 경우, IDE에서 집어줌
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice) // [0 0 0 0]

	// 용량이 4인데 5번째 요소를 추가하면 용량인 두 배인 8개가 됨
	aSlice = append(aSlice, 5) // [0 0 0 0 5]
	fmt.Println(aSlice)
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice)) // 5, 8

	// 길이가 5개인데 4개의 요소가 더 추가되면, 용량은 두 배인 16이
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)                               // [0 0 0 0 5 -1 -2 -3 -4]
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice)) // 9, 16
}
