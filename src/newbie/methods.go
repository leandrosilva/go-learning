package main

import "fmt"

type integer int

func (i integer) logByValue() {
  i++
  fmt.Println("[logByValue] i =", i)
}

func (i *integer) logByPointer() {
  *i++
  fmt.Println("[logByPointer] i =", *i)
}

type Entry struct {
  Val string
  count int
}

func (e Entry) LogByValue() {
  e.count++
  fmt.Println("[LogByValue] e = ", e)
}

func (e *Entry) LogByPointer() {
  e.count++
  fmt.Println("[LogByPointer] e = ", e)
}

func main() {
  var i1 integer = 1000  
  i1.logByValue()
  fmt.Println("i1 =", i1, "\n")

  var i2 integer = 1000  
  i2.logByPointer()
  fmt.Println("i2 =", i2, "\n")
  
  fmt.Println("e1 := Entry{Val: 'x1', count: 10}")
  e1 := Entry{Val: "x1", count: 10}
  e1.LogByValue()
  fmt.Println("e1.count =", e1.count, "\n")
  
  fmt.Println("e2 := &Entry{Val: 'x2', count: 10}")
  e2 := &Entry{Val: "x2", count: 10}
  e2.LogByValue()
  fmt.Println("e2.count =", e2.count, "\n")
  
  fmt.Println("e3 := Entry{Val: 'x3', count: 10}")
  e3 := Entry{Val: "x3", count: 10}
  e3.LogByPointer()
  fmt.Println("e3.count =", e3.count, "\n")
  
  fmt.Println("e4 := &Entry{Val: 'x4', count: 10}")
  e4 := &Entry{Val: "x4", count: 10}
  e4.LogByPointer()
  fmt.Println("e4.count =", e4.count)
}
