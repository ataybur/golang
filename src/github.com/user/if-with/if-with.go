package main

import (
	"fmt"
	"math"
	)

func pow(x,n,lim float64) float64 {
	if x := math.Pow(x,n) ; x < lim {
		return x
	}
	return lim
}

func main(){
	fmt.Println(pow(3,2,10))
	fmt.Println(pow(3,3,20))
}
