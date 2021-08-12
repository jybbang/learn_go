package main

import "fmt"

func main() {
	house1 := getBuilder().setDoorType("door1").setWindowType("windows1").setNumFloor(10).build()
	house2 := getBuilder().setDoorType("door2").setWindowType("windows2").build()
	fmt.Println(house1)
	fmt.Println(house2)
}
