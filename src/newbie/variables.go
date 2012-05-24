package main

import "fmt"

func main() {
  var i int
  var x float32
  var complex_pointer *complex128
  int_pointer := &i
  another_int_pointer := new(int)
  generic_channel := make(chan interface{})
  
  fmt.Println("i =", i)
  fmt.Println("x =", x)
  fmt.Println("complex_pointer =", complex_pointer)
  fmt.Println("int_pointer =", int_pointer)
  fmt.Println("another_int_pointer =", another_int_pointer)
  fmt.Println("generic_channel =", generic_channel)
  fmt.Println("*int_pointer =", *int_pointer)
}
