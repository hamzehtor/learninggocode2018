package main
import "fmt"

func waapper() func() int {
  x := 0
  return func() int {
    x++
    return x
  }
}
func main()  {
  increment := waapper()
  fmt.Println(increment())
  fmt.Println(increment())
}
