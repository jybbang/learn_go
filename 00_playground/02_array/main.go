package main

import "fmt"

func main() {
	// 배열은 길이까지 하나의 타입으로 판단
	// 따라서 아래는 다른 배열
	// sample1 := [1]int{1}
	// sample2 := [2]int{1,2}

	// //Both number of elements and actual elements
	// sample1 := [2]int{1, 2}
	// fmt.Printf("Sample1: Len: %d, %v\n", len(sample1), sample1)

	// //Only actual elements
	// sample2 := [...]int{2, 3}
	// fmt.Printf("Sample2: Len: %d, %v\n", len(sample2), sample2)

	// //Only number of elements
	// sample3 := [2]int{}
	// fmt.Printf("Sample3: Len: %d, %v\n", len(sample3), sample3)

	// //Without both number of elements and actual elements
	// sample4 := [...]int{}
	// fmt.Printf("Sample4: Len: %d, %v\n", len(sample4), sample4)

	// 배열은 값타입이다. 배열을 전달해도 값복사가 기본이다
	sample1 := [2]string{"a", "b"}
	fmt.Printf("Sample1 Before: %v\n", sample1)
	sample2 := sample1
	sample2[0] = "c"
	fmt.Printf("Sample1 After assignment: %v\n", sample1)
	fmt.Printf("Sample2: %v\n", sample2)
	test(sample1)
	fmt.Printf("Sample1 After Test Function Call: %v\n", sample1)
}

func test(sample [2]string) {
	sample[0] = "d"
	fmt.Printf("Sample in Test function: %v\n", sample)
}
