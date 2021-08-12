package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	// a := []string{"a", "b"}
	// checkAndPrintWithRecover(a, 2)
	// fmt.Println("Exiting normally")

	// 동일한 고 루틴에 있는 패닉만 복구할 수 있다
	a := []string{"a", "b"}
	checkAndPrintWithRecover(a, 2)
	time.Sleep(time.Second)
	fmt.Println("Exiting normally")
}

// recover로 복구 가능
// panic시에도 defer는 실행되므로 defer 로 선언해야함
// 패닉이 복구 될때 반환값은 자료형 기본값
// func checkAndPrintWithRecover(a []string, index int) {
// 	defer handleOutOfBounds()
// 	checkAndPrint(a, 2)
// }
func checkAndPrintWithRecover(a []string, index int) {
	defer handleOutOfBounds()
	go checkAndPrint(a, 2)
}
func checkAndPrint(a []string, index int) {
	if index > (len(a) - 1) {
		panic("Out of bound access for slice")
	}
	fmt.Println(a[index])
}

func handleOutOfBounds() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
		// 패닉 스택확인 가능
		debug.PrintStack()
	}
}
