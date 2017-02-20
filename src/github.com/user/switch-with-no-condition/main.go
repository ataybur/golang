package main

import (
	"fmt"
	"time"
       )


func main (){
	t:= time.Now()
 	switch {
	case t.Hour() < 12 :
		fmt.Println("Goodmorning")
	case t.Hour() < 17 :
		fmt.Println("Good Afternoon")
	default :
		fmt.Println("Good evening")
	}

}
