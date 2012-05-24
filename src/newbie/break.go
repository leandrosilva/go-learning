package main

import "fmt"

func main() {
  for i :=0; i < 10; i++ {
    L:
    
    for {
      fmt. Println("Infinity #1")
      
      for {
        fmt. Println("Infinity #2")
        break L
      }
    }
    
    fmt.Println("i =", i)
  }
}
