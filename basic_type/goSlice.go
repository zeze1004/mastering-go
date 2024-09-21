package main

import "fmt"

func main() {
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice)) // [] 0 0

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice)) // [1234.56 -34] with length 2

	// 슬라이스 길이: 4
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	t = append(t, -5)
	fmt.Println(t) // [-1 -2 -3 -4 -5]

	// 2차원 슬라이스
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	for _, i := range twoD { // 행 수 반복
		for _, k := range i { // 행열값 출력
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	// 2차원 슬라이스를 만들고 초기화하는
	make2D := make([][]int, 2)
	fmt.Println(make2D) // [[] []]
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D) // [[1 2 3 4] [-1 -2 -3 -4]]
}
