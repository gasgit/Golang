package main

import "fmt"

func main(){
  var a [4] int
  fmt.Println("Empty array:", a)
  fmt.Println("Length a:", len(a))

  b := [4]int{0,1,2,3}
  fmt.Println("b:", b)
  fmt.Println("Value in pos 2:", b[2])



}
