package main

import (
	"fmt"
	"strings"
)

func main() {
	// 슬라이스는 배열의 포인터를 가지고 있는 자료타입이다
	// 기존 슬라이스를 복사해도 슬라이스가 참조하는 배열도
	// 포인터 형태로 복사되므로 마치 참조하는 것처럼 사용할 수 있다.
	// numbers := [5]int{1, 2, 3, 4, 5}

	// //Both start and end
	// num1 := numbers[2:4]
	// fmt.Println("Both start and end")
	// fmt.Printf("num1=%v\n", num1)
	// fmt.Printf("length=%d\n", len(num1))
	// fmt.Printf("capacity=%d\n", cap(num1))

	// //Only start
	// num2 := numbers[2:]
	// fmt.Println("\nOnly start")
	// fmt.Printf("num1=%v\n", num2)
	// fmt.Printf("length=%d\n", len(num2))
	// fmt.Printf("capacity=%d\n", cap(num2))

	// //Only end
	// num3 := numbers[:3]
	// fmt.Println("\nOnly end")
	// fmt.Printf("num1=%v\n", num3)
	// fmt.Printf("length=%d\n", len(num3))
	// fmt.Printf("capacity=%d\n", cap(num3))

	// //None
	// num4 := numbers[:]
	// fmt.Println("\nOnly end")
	// fmt.Printf("num1=%v\n", num4)
	// fmt.Printf("length=%d\n", len(num4))
	// fmt.Printf("capacity=%d\n", cap(num4))

	// 기존배열을 참고하는 것이 아닌, 새로운 배열을 만들고 그 배열을 참조하는 슬라이스를 반환
	// numbers := make([]int, 3, 5)
	// fmt.Printf("numbers=%v\n", numbers)
	// fmt.Printf("length=%d\n", len(numbers))
	// fmt.Printf("capacity=%d\n", cap(numbers))

	// //With capacity ommited
	// numbers = make([]int, 3)
	// fmt.Println("\nCapacity Ommited")
	// fmt.Printf("numbers=%v\n", numbers)
	// fmt.Printf("length=%d\n", len(numbers))
	// fmt.Printf("capacity=%d\n", cap(numbers))

	// 슬라이스는 원본 배열의 참조이므로, 값을 변경하면 원본배열도 값이 변경된다
	// array := [5]int{1, 2, 3, 4, 5}
	// slice := array[:]

	// //Modifying the slice
	// slice[1] = 7
	// fmt.Println("Modifying Slice")
	// fmt.Printf("Array=%v\n", array)
	// fmt.Printf("Slice=%v\n", slice)

	// //Modifying the array. Would reflect back in slice too
	// array[1] = 2
	// fmt.Println("\nModifying Underlying Array")
	// fmt.Printf("Array=%v\n", array)
	// fmt.Printf("Slice=%v\n", slice)

	// numbers := make([]int, 3, 5)
	// numbers[0] = 1
	// numbers[1] = 2
	// numbers[2] = 3
	// fmt.Printf("numbers=%v\n", numbers)
	// fmt.Printf("length=%d\n", len(numbers))
	// fmt.Printf("capacity=%d\n", cap(numbers))

	// //Append number 4
	// numbers = append(numbers, 4)
	// fmt.Println("\nAppend Number 4")
	// fmt.Printf("numbers=%v\n", numbers)
	// fmt.Printf("length=%d\n", len(numbers))
	// fmt.Printf("capacity=%d\n", cap(numbers))

	// //Append number 5
	// numbers = append(numbers, 5)
	// fmt.Println("\nAppend Number 5")
	// fmt.Printf("numbers=%v\n", numbers)
	// fmt.Printf("length=%d\n", len(numbers))
	// fmt.Printf("capacity=%d\n", cap(numbers))

	// 콜렉션 펑션
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))

}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
