package Mchannel

import "time"

func TestChan(){

	ch := make(chan int)
	// channel 写入需要在另一个携程中
	go doPlus(2,3,ch)
	println(<-ch)
}

func doPlus(a,b int, ch chan int){
	time.Sleep(10000)
	ch <- a + b
	close(ch)
}


