package main

import (
  "fmt"
  "time"
  )
  
  func main() {
    go countFruit("Apple")
    go countFruit("Orange")
    time.Sleep(time.Second * 10)
  }
  
  func countFruit(name string){
    for i := 1; true; i++ {
      fmt.Println(name)
      time.Sleep(time.Second * 1)
    }
  }