package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!") // 대문자로 변경
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE")) // 소문자로 변경

	f("%s\n", s.Title("tHis wiLL be A title!")) // 우와 앞글자만 대문자

	// 대소문자 무시하고 같은 문자열인지 비교
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	// prefix로 해당 문자열이 시작됐는지 비교
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	// suffix로 해당 문자열이 끝났는지 비교
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))

	// ha가 있는 인덱스(첫글자 h)를 반환, 없으면 -1 반환
	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))

	// 해당 글자의 개수를 반환
	f("Count: %v\n", s.Count("Mihalis", "i"))
	f("Count: %v\n", s.Count("Mihalis", "I"))

	// 해당 문자열을 5번 반복 ababababab
	f("Repeat: %s\n", s.Repeat("ab", 5))

	// 문자열의 앞뒤 공백을 제거
	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	// 문자열의 왼쪽 공백을 제거
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	// 문자열의 오른쪽 공백을 제거
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	// 두 문자열을 비교, 첫번째 문자열이 두번째 문자열보다 작으면 -1, 같으면 0, 크면 1
	f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
	f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
	f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis")) // 소문자의 유니코드가 더 크므로, -1 반환

	// 띄어쓰기 기준으로 문자열 개수
	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	// 문자열을 분리
	f("%s\n", s.Split("abcd efg", "")) // [a b c d   e f g]

	// 글자 사이에 _ 추가
	f("%s\n", s.Replace("abcd efg", "", "_", -1)) // _a_b_c_d_ _e_f_g_
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2)) // _a_bcd efg

	// 슬라이스 요소 사이에 문자열 추
	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++")) // Join: Line 1+++Line 2+++Line 3

	// 문자열 분리
	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++")) // [123++ 432++ ]

	// 문자가 아니면 문자열에서 trim
	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction)) // TrimFunc: abc ABC
}
