package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 고루틴을 생성하기 바로 전에 Add(1)을 호출해 경쟁 상태를 피했다
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done() // defer 키워드로 익명 함수가 끝나기 전에 Done() 호출
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()             // waitGroup 변수의 카운터가 0이 될 때까지 기다렸다가 반환
	fmt.Println("프로그램이 종료됐습니다.") // 모든 고루틴이 끝나면 출력
}
