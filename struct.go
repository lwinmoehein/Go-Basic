package main

import "fmt"

type Student struct {
name string
age int
}

func (s *Student) introduce() {
  fmt.Println(s.age)
}

func (s *Student) aged() {
  s.age++
  fmt.Println(s.age)
}
func main(){

   var lwin = Student {"lwin",22}
   fmt.Println(lwin.name)
   lwin.introduce()
   lwin.aged()
   lwin.aged()
}
