package main

import (
	"fmt"
	"time"
)

func main() {
	// 송수신 전용채널 생성 가능
	// func process(ch chan<- int){ //doSomething }
	// func process(ch <-chan int){ //doSomething }
	// ch := make(chan int)
	// fmt.Println("Sending value to channel start")
	// go send(ch)
	// val := <-ch
	// fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)

	// ch := make(chan int, 3)
	// ch <- 2
	// ch <- 2
	// ch <- 2
	// close(ch)
	// go sum(ch)
	// time.Sleep(time.Second * 1)

	// select에서 각 case 문은 채널에서 보내기 또는 받기 작업을 대기
	// 여러 case 문이 준비되면 무작위로 하나를 선택하여 진행
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go goOne(ch1)
	// go goTwo(ch2)

	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println(msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println(msg2)
	// }

	// 반복문으로 해당 케이스 모두 점검가능
	// for i := 0; i < 2; i++ {
	//     select {
	//     case msg1 := <-ch1:
	//         fmt.Println(msg1)
	//     case msg2 := <-ch2:
	//         fmt.Println(msg2)
	//     }
	// }

	// time 패키지에서 제공하는 채널
	// ch1 := make(chan string)
	// go goOne(ch1)

	// select {
	// case msg := <-ch1:
	// 	fmt.Println(msg)
	// case <-time.After(time.Second * 1):
	// 	fmt.Println("Timeout")
	// }

	// 조건이 완료될때까지 무한히 동작
	news := make(chan string)
	go newsFeed(news)
	printAllNews(news)
}

func printAllNews(news chan string) {
	for {
		select {
		case n := <-news:
			fmt.Println(n)
			// 한번 실행한 이후에는 해당 채널은 무시된다
			// news = nil
		case <-time.After(time.Second * 1):
			fmt.Println("Timeout: News feed finished")
			return
		}
	}
}

func newsFeed(ch chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Millisecond * 400)
		ch <- fmt.Sprintf("News: %d", i+1)
	}
}

func goOne(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goOne goroutine"
}

// func send(ch chan int) {
// 	ch <- 1
// }

// func sum(ch chan int) {
// 	sum := 0
// 	// 채널이 닫힐때까지 데이터를 수신함 <- 암시적
// 	for val := range ch {
// 		sum += val
// 	}
// 	fmt.Printf("Sum: %d\n", sum)
// }

// func goOne(ch chan string) {
// 	ch <- "From goOne goroutine"
// }

// func goTwo(ch chan string) {
// 	ch <- "From goTwo goroutine"
// }
