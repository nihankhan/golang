package main

import "fmt"

func main(){
  var(
       p int
       q int
       i int
       n int
       t int
       sum int
    )
    
    fmt.Scanln(&t)
    for i = 1; i <= t; i++{
      fmt.Scanln(&n)
      p = n % 10
      q = n / 10000
      sum = p + q
      
      fmt.Println("Sum : ", sum)
    }
}