package main

import (
	"fmt"
	"sync"
)

var (
	allData = make(map[string]string)
	rwm     sync.RWMutex
)

func get(key string) string {
	rwm.RLock()
	defer rwm.RUnlock()
	return allData[key]

}

func set(key string, value string) {
	rwm.Lock()
	defer rwm.Unlock()
	allData[key] = value

}

func main() {
	// map도 슬라이스의 한 유형이다
	// 키가 존재하는 경우 ok == true
	//Declare
	// employeeSalary := make(map[string]int)

	// //Adding a key value
	// employeeSalary["Tom"] = 2000
	// fmt.Println("Key exists case")
	// val, ok := employeeSalary["Tom"]
	// fmt.Printf("Val: %d, ok: %t\n", val, ok)
	// fmt.Println("Key doesn't exists case")

	// val, ok = employeeSalary["Sam"]
	// fmt.Printf("Val: %d, ok: %t\n", val, ok)

	// for k, v := range employeeSalary {
	// 	fmt.Printf("key :%s value: %d\n", k, v)
	// }

	// 스레드 세이프 하지 않다.
	// 싱크 활용
	set("a", "Some Data")
	result := get("a")
	fmt.Println(result)
}
