package main

import "fmt"

type taxSystem interface {
	calculateTax() int
}
type indianTax struct {
	taxPercentage int
	income        int
}

func (i *indianTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type singaporeTax struct {
	taxPercentage int
	income        int
}

func (i *singaporeTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type usaTax struct {
	taxPercentage int
	income        int
}

func (i *usaTax) calculateTax() int {
	tax := i.income * i.taxPercentage / 100
	return tax
}

type animal interface {
	breathe()
	walk()
}

type lion struct {
	age int
}

func (l lion) breathe() {
	fmt.Println("Lion breathes")
}

func (l lion) walk() {
	fmt.Println("Lion walk")
}

type dog struct {
	age int
}

func (d dog) breathe() {
	fmt.Println("Dog breathes")
}

func (d dog) walk() {
	fmt.Println("Dog walk")
}

type cat string

func (c cat) breathe() {
	fmt.Println("Cat breathes")
}

func (c cat) walk() {
	fmt.Println("Cat walk")
}

func main() {
	// indianTax := &indianTax{
	// 	taxPercentage: 30,
	// 	income:        1000,
	// }
	// singaporeTax := &singaporeTax{
	// 	taxPercentage: 10,
	// 	income:        2000,
	// }

	// taxSystems := []taxSystem{indianTax, singaporeTax}
	// totalTax := calculateTotalTax(taxSystems)

	// fmt.Printf("Total Tax is %d\n", totalTax)

	// 인퍼테이스는 덕타이핑이다
	// 심지어 해당 함수만 구현되면 primitive 한 객체도 interface로 볼수 있다
	var a animal

	// a = cat("smokey")
	// a.breathe()
	// a.walk()

	a = lion{age: 10}
	print(a)
}

// func calculateTotalTax(taxSystems []taxSystem) int {
// 	totalTax := 0
// 	for _, t := range taxSystems {
// 		totalTax += t.calculateTax() //This is where runtime polymorphism happens
// 	}
// 	return totalTax
// }

func print(a animal) {
	// 인터페이스 타입 어서션
	// l, ok := a.(lion)
	// if ok {
	// 	fmt.Println(l)
	// } else {
	// 	fmt.Println("a is not of type lion")
	// }

	// d, ok := a.(dog)
	// if ok {
	// 	fmt.Println(d)
	// } else {
	// 	fmt.Println("a is not of type lion")
	// }

	switch v := a.(type) {
	case lion:
		fmt.Println("Type: lion")
	case dog:
		fmt.Println("Type: dog")
	default:
		fmt.Printf("Unknown Type %T", v)
	}
}
