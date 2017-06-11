package main

import "fmt"
import "math/rand"


/*
  decisions, if, if-else
*/

func main(){
  var x int = 12
  // if
  if( x < 15 ){
    fmt.Printf("x is less than 15\n")
  }
  fmt.Printf("Value of x: %d\n",x)
  // if else
  if( x > 10){
    fmt.Printf("x is greater than 10\n")
  }else{
    fmt.Printf("x is not greater than 10\n")
  }

  // if else-if else
  if( x == 10){
    fmt.Println("x is 10\n")
  }else if(x == 20){
    fmt.Println("x is 200\n")
  }else{
    fmt.Println("x not found")
  }

  // switch


  choice := rand.Intn(3 - 1) + 1

  fmt.Printf("Rand int: %d\n", choice)


  switch choice {
  case 1:
    fmt.Println("1 selected")
  case 2:
    fmt.Println("2 selected")
  case 3:
    fmt.Println("3 selected")

  }


}
