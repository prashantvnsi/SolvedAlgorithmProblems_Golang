package main

import (
	"fmt"
)

func main() {
	eve:= make(chan int)
	odd:= make(chan int)
	quit:= make(chan int)

	// send to the channel
	go send(eve, odd, quit)

	// receive to the channel
	receive(eve, odd, quit)
}

func send (e, o, q chan<- int){
	for i:=0; i<10; i++{
		if i%2 == 0{
			e, ok = <- i
		}else{
			o <- i
		}
	}
	q <- 0
}

func receive (e, o, q <-chan int){
	for{
		select{
		case v := <- e:
			fmt.Println("from the even channel", v)
		case v := <- o:
			fmt.Println("from the odd channel", v)
		case v := <- q:
			fmt.Println("from the quit channel", v)
			return
		}
	}
}