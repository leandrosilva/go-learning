package main

import "fmt"

type Car struct {
  Color string
  x int
}

func main() {
  var c1 Car
  var c2 *Car
  c3 := Car{Color: "black", x: 10}
  c4 := &Car{Color: "red", x: 2}
  
  fmt.Println("c1 =", c1)
  fmt.Println("c2 =", c2)
  fmt.Println("c3 =", c3)
  fmt.Println("c4 =", c4)
}