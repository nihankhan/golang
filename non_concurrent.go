package main

import (
  "fmt"
  "time"
  )
  
  func main(){
    countFruit("Apole")
    countFruit("Orange")
  }
 
 func countFruit(name string){
   for i := 1; true; i++{
     fmt.Println(name)
     time.Sleep(time.Second * 1)
   }
 }