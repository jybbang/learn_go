package main

import (
	"fmt"
)

// type employee struct {
// 	name   string
// 	age    int
// 	salary int
// }

// go 는 소문자로 시작하는 변수는 unexported 로 무시하므로
// json 패키지에서 참조할 수 있도록 대문자 사용
// type employee struct {
// 	Name   string
// 	Age    int
// 	salary int
// }

type employee struct {
	name   string
	age    int
	salary int
	address
}

type address struct {
	city    string
	country string
}

func main() {
	// 00 구조체포인터  생성 with & 키워드
	// emp := employee{name: "Sam", age: 31, salary: 2000}
	// empP := &emp
	// fmt.Printf("Emp: %+v\n", empP)
	// empP = &employee{name: "John", age: 30, salary: 3000}
	// fmt.Printf("Emp: %+v\n", empP)

	// 01 구조체포인터  생성 with new 키워드
	// empP := new(employee)
	// fmt.Printf("Emp Pointer Address: %p\n", empP)
	// fmt.Printf("Emp Pointer: %+v\n", empP)
	// fmt.Printf("Emp Value: %+v\n", *empP)
	//Marshal

	// 02 구조체 JSON 변혼
	// emp := employee{Name: "Sam", Age: 31, salary: 2000}
	// empJSON, err := json.Marshal(emp)
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("Marshal funnction output %s\n", string(empJSON))

	// //MarshalIndent
	// empJSON, err = json.MarshalIndent(emp, "", "  ")
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }
	// fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))

	// 03 익명구조체
	// 정확히 상속은 아니지만, 상속과 비슷한 효과
	address := address{city: "London", country: "UK"}

	emp := employee{name: "Sam", age: 31, salary: 2000, address: address}

	fmt.Printf("City: %s\n", emp.address.city)
	fmt.Printf("Country: %s\n", emp.address.country)

	fmt.Printf("City: %s\n", emp.city)
	fmt.Printf("Country: %s\n", emp.country)
}
