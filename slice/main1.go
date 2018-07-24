package main
import "fmt"
func main()  {
  myslice := make([]int, 3, 5)

  fmt.Println("----------------------")
  fmt.Println(myslice)
  fmt.Println(len(myslice))
  fmt.Println(cap(myslice))
  fmt.Println("----------------------")

  for i := 1; i < 10; i++ {
    if panic {
      myslice[i-1] := append(i)
    }
    // I have to woke on this case.
    myslice[i-1] = i
    fmt.Println("Len", len(myslice), "capacity", cap(myslice), "value" ,myslice[i])
    fmt.Println(myslice)
  }
}
