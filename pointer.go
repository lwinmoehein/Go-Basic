package main

import "fmt"

func main(){
//declare a pointet variable
var pointer *string
name:="lwinmoehein is god"
pointer=&name
fmt.Println(pointer)
fmt.Println("value is"+*pointer)
}
