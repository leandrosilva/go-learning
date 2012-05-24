package main

import "fmt"

const (
  Red       = (iota)
  Green     = (iota)
  Blue      = (iota)
  ColorMask = (1 << (iota + 1)) - 1
)

const (
  Low    = (iota)
  Medium = (iota)
  High   = (iota)
)

const (
  i complex128 = complex(0, 1)
  j complex128 = complex(2, 3)
)

const (
  A = "a"
  B = "b"
  C = "c"
)

func main() {
  fmt.Println("Red       =", Red)
  fmt.Println("Green     =", Green)
  fmt.Println("Blue      =", Blue)
  fmt.Println("ColorMask =", ColorMask)
  fmt.Println()
  fmt.Println("Low    =", Low)
  fmt.Println("Medium =", Medium)
  fmt.Println("High   =", High)
  fmt.Println()
  fmt.Println("i =", i)
  fmt.Println("j =", j)
  fmt.Println()
  fmt.Println("A =", A)
  fmt.Println("B =", B)
  fmt.Println("C =", C)
}
