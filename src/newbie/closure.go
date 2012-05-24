package main

import "fmt"

func add(init int) (func (int) int) {
  return func (x int) int {
    return init + x
  }
}

func main() {
  add_10 := add(10)
  fmt.Println("add_10(3) =", add_10(3))
  fmt.Println("add_10(5) =", add_10(5))
  fmt.Println("add_10(7) =", add_10(7))
}
