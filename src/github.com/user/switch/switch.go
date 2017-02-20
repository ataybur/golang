package main

import (
	"fmt"
	"runtime"
       )

func main() {
     fmt.Print("Go runs on ")

     switch os:= runtime.GOOS; os {
     case "darvin" :
         fmt.Println("OS X")
     case "linux"  :
         fmt.Println("linux")
     default       :
         fmt.Printf("%s",os)
     }
}

