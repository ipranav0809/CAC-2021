package main

import (
	"fmt"
	"time"
)

func main(){
	timer 1 := time.NewTimer(25*time.Minute)

//timer 1(25 minutes of work)
	<-timer1.C
	fmt.Println("25 minutes starts now")

	stop1 := timer1.Stop()
	if stop1 {
		fmt.Println("Timer stopped")
	}

//timer 2(5 minutes of break)
	timer2 := time.NewTimer(5*time.Minute)

	<-timer2.C
	fmt.Println("5 minute break starts now")

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer stopped")
	}
}
