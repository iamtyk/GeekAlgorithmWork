package byteDanceFaceTry

import (
	"fmt"
	"time"
)

var n = 1

func Producer(channel chan int) {
	for n <= 10 {
		channel <- n
		n++
		time.Sleep(time.Second / 10)
	}
	close(channel)
}

func Customer(channel <-chan int, threadName string, flagChan chan bool) {
	fmt.Println("start")
	for {
		val, ok := <-channel
		if ok {
			fmt.Print(threadName)
			fmt.Println(val)
		} else {
			close(flagChan)
			break
		}
		flagChan <- true
	}

}
