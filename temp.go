package main
  
import (
  "fmt"
  "time"
)
  
func Print() {
  fmt.Println("printing")
}

func main() {
  go Print()
  time.Sleep(time.Second)
}
