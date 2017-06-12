package main

import "fmt"

type vechicle struct{
  make string
  model string
  engine_size float32

}

func main(){
  
  p1 := vechicle{make: "toyota", model: "avensis", engine_size: 1.6}

  fmt.Println(p1)
  fmt.Println(p1.make)
}
