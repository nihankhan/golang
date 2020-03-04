//go 1.10.4

package main
import "fmt"
func main() {
  var (
    n int
    i int
    j int
    
    )
  
  fmt.Scanln(&n)
  for i = 1; i <= n; i++{
    for j = 1; j <= i; j++{
      fmt.Printf("*")
    }
    
    fmt.Printf("\n")
  }
}