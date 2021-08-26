 package main

 import "fmt"

 func main(){
	var input int
	fmt.Scanln(&input)
	for i:=1;i<=input;i++ {
           if i%15==0 {
            fmt.Println("FizzBuzz")
           }
          if i%3==0 {
            fmt.Println("Fizz")  
          }
          if i%5==0 {
           fmt.Println("Buzz")
          }
        }
 }
