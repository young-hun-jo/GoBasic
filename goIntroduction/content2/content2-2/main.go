package main

import "fmt"

// 4. 변수 정의: [var 변수명 타입] 순으로 정의
// int, float64는 초깃값 지정하지 않으면 0(0.)으로 할당됨
// string은 ""(빈 문자열)로 초깃값 할당
var a int
var b string
var c float64

// 변수 여러개를 한 번에 생성
var name, id, address string // 같은 타입일 때
// 특정 구간에 사용할 타입이 서로 다른 변수들을 소괄호로 묶어 한 번에 정의
var (
	nickname string
	age      int
	weight   float32
)

// 특정 초깃값을 할당함으로써 변수 선언시 타입을 생략도 가능
var flag = true
var size = uint16(1024)

// 5. 상수 정의: bool, 숫자, 문자열 타입으로만 선언 가능
const (
	limit  = 64
	keep   = true
	userId = "1345"
)
const limit2 uint8 = 20 // 상수 정의할 때 명시적으로 타입을 지정하고 싶다면!
// 특정한 계산식의 결과로 상수를 지정할 수도 있는데, 함수나 메소드로 연산하는 결과는 할당 불가(왜냐면 컴파일할 때 연산이 가능해야 함)
const value1 = 10 * 10

// const value2 = getNumber() // error 발생

// 상수의 묶음인 열거형: 차례로 1씩 증가하는 상수의 묶음
const (
	zero = iota
	one
	two
	three
)

type Color int

const (
	Red Color = iota
	Yello
	Green
)

type ByteSize int64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota) // 1의 비트를 왼쪽으로 10 * iota 만큼 움직인 후의 비트값!
	MB
	GB
)

func main() {
	fmt.Println("a:", a, "b:", b, "c:", c)
	// 짧은 선언: 변수 선언과 동시에 초깃값을 할당할 때-> 단, global한 scope(전역변수)에서는 짧은 선언 불가
	init := 1
	fmt.Println("init:", init)

	fmt.Println(zero, one, two, three)
	fmt.Println(Red, Yello, Green)
	fmt.Println(KB, MB, GB)
}

func getNumber() int {
	return 100
}
