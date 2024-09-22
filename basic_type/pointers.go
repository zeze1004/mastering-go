package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

// float64 포인터를 입력으로 받음
func processPointer(x *float64) {
	*x = *x * *x // x의 역참조
}

// &를 사용해 변수의 메모리 주소 반환
func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

// x가 저장한 메모리 주소에 있는 값을 가져와서 2배를 한 후 그 값을 저장한 메모리 주소를 반환
func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	f := 12.123
	fmt.Println("Memory address of f:", &f)

	fP := &f
	fmt.Println("Memory address of f:", fP) // FP의 값은 f의 메모리 주소
	fmt.Println("Value of f:", *fP)         // fp가 가르키는 주소의 값을 역참조함

	// f의 값은 변경됨
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f) // 제
	// f의 값이 바뀌지 않음
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)
	// f의 값을 두 배로 바꾼 후에 그 값을 저장한 메모리 주소를 반환하므로 f의 값은 바뀌지 않음
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	// Check for empty structure
	var k *aStructure
	// This is nil because currently k points to nowhere
	fmt.Println(k)
	// Therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
