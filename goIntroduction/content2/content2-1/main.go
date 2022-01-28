package main

import "fmt"

func main() {
	res1 := basicloop1()
	fmt.Println("res1:", res1)
	res2 := basicloop2()
	fmt.Println("res2:", res2)
	res3 := basicloop3()
	fmt.Println("res3:", res3)

	c := "a"
	fmt.Println(c < "b") // 소문자 영어 아스키코드 > 대문자 영어 아스키코드
	check(c)
	semicolon()

}

// 1. go는 무한루프 while이 없음. for loop로 구현해야 함
func basicloop1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func basicloop2() int {
	sum, i := 0, 0
	// for loop에 조건식만 사용
	for i < 10 {
		sum += i
		i++ // i += 1 도 가능
	}
	return sum
}

func basicloop3() int {
	sum, i := 0, 0
	// 조건식 생략한 후 break 사용한 루트
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++ // 증간 연산은 반환값이 없어야 함!
	}
	return sum
}

// 2. 복잡한 if 대신 switch, case 구문을 사용
func check(c string) {
	switch {
	case "0" < c && c <= "9":
		fmt.Printf("%s는 숫자입니다", c)
	case "a" <= c && c <= "z":
		fmt.Printf("%s는 소문자입니다", c)
	case "A" <= c && c <= "Z":
		fmt.Printf("%s는 대문자입니다", c)
	}
}

// 3. 세미콜론: 원래 Go 컴파일러는 세미콜론 기준으로 코드 문장의 단위를 인식
// 하지만, 중괄호의 여는 괄호는 반드시 제어문, 함수가 시작되는 줄의 끝에 써야 함
func semicolon() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	/* 이런식으로 하면 Go 컴파일러가 i ++ 끝에 자동으로 세미콜론을 붙여 인식해서 컴파일 에러가 발생!
	for i := 0; i <10; i ++
	{
		fmt.Println("i:", i)
	}
	*/
}
