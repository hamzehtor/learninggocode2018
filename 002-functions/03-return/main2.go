// another progrem with the same result
package main

import "fmt"

func main()  {
  data := []float64{43,25,52,2,6}
  fmt.Printf("%T", data)
  fmt.Println("\n")
  n := average(data...)
  // "..." say that one by one gay of the list is the goal
  fmt.Println(n ,"\n")
  fmt.Printf("%T", n)
  fmt.Println("\n")
}

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
