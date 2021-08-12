package main

import (
	"fmt"
	"sync"
)

// var lock = &sync.Mutex{}
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
	if singleInstance == nil {
		// lock.Lock()
		// defer lock.Unlock()

		// if singleInstance == nil {
		// 	fmt.Println("Creating single instnace.")
		// 	singleInstance = &single{}
		// } else {
		// 	fmt.Println("Single instance already created.")
		// }
		once.Do(
			func() {
				fmt.Println("Creating single instnace.")
				singleInstance = &single{}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}
