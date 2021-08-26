package main

import "fmt"

func main() {
    str := "hello go"
    for _,v := range  str {
        fmt.Printf("%c\n",v)
    }
}
