package main

import (
	"errors"
	"fmt"
	"os"
)

// check 인자 a, b가 모두 0이면 에러를 리턴
func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("이거슨 커스텀 에러입니다")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("%d, %d 값이 모두 0입니다, UserID: %d", a, b, os.Getuid()) //os.Getuid() 함수는 현재 실행 중인 프로세스의 사용자의 UID를 반환하는데, 로컬 macOS 환경에서 실행하는 경우 501이 관리자 계정의 UID / root 계정은 0
	}

	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("에러 없슴둥")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "이거슨 커스텀 에러입니다" {
		fmt.Println("커스텀 에러 발생")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
