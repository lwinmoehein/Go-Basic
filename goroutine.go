package main

import (
        "fmt"
        "time")

type Chunk struct {
start int
end int
count int
}


func getChunk(c *Chunk,ch chan Chunk){
  ch<-*c
}

func main(){

   chunk1:=make(chan Chunk)
   chunk2:=make(chan Chunk)
   c1:=Chunk{1,1,3}
   c2:=Chunk{2,4,5}
   
   go getChunk(&c1,chunk1)
   go getChunk(&c2,chunk2)
   
   for {   
   select {
     case ch1:=<-chunk1:
        fmt.Println(ch1)
        return
     case ch2:=<-chunk2:
        fmt.Println(ch2)
        return
     default:
        fmt.Println("none completed")
       time.Sleep(50*time.Millisecond)
   }
  }
}
