/*
	Go 언어 공부 시작
	25.05.14 . razorbacks
*/

package main

import "fmt"

func main() {

	// Golang에서는 직접적으로 변수를 선언하지 않아도 간결하게 ':=' 를 이용하여 자동으로 결정하여 변수를 선언할 수 있다.
	// 이것이 Golang의 간결성이다.

	name := "홍길동"
	city := "경기도 이천"
	age := 35
	height := 173.6
	var dead bool = true
	const Pi = 3.14

	// deadStatus 변수를 선언하고 반복문을 통해 deadStatus에 값을 전달한다.
	var deadStatus string
	if dead {
		deadStatus = "사망"
	} else {
		deadStatus = "생존"
	}

	// const Pi는 int 혹은 float인 경우에는 자동으로 선택해주는데, char 또는 string인 경우에는 처리하지 못 함.

	// var name string = "윤찬렬"
	// var city string = "경기도 부천"
	// var age int = 35
	// var height float32 = 173.6

	fmt.Println("이름 : ", name)
	fmt.Println("지역 : ", city)
	fmt.Println("사망 여부 : ", deadStatus)
	fmt.Printf("나이 : %d세, 키 : %.1fcm\n", age, height)
	fmt.Println("파이는 몇 인가? : ", Pi)
}
