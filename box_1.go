package main

import "fmt"

func main()
{
  var (
        t int 
        n int      
    )
    
    fmt.Scanln(&t)
    
    for i := 1; i <= t; i++{
      fmt.Scanln(&n)
      for j := 1; j <= n; j++{
        for k := 1; k <= n; k++{
          fmt.Println("*")
        }
        
        fmt.Println("\n")
      }
      
      if i != t{
        fmt.Println("\n")
      }
    }
}