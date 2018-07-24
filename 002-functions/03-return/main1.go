package main

import "fmt"

func main()  {
  n:=average(43,25,52,2,6)
  fmt.Println(n ,"\n")
  fmt.Printf("%T", n)
//
  fmt.Println("\n")
}
// float64 is a type like int
// for undrastanding this type run the code
// variadic "..." len of inputs is not known
func average(sf ...float64) float64 {
  fmt.Println(sf, "\n")
  fmt.Printf("%T", sf)
  fmt.Println("\n")
  total := 0.0
  for  _ , v := range sf {
    total += v
  }
  return total / float64(len(sf))
}
