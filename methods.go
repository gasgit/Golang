package main

import "fmt"

type actions struct{
  x, y float32
}

func (a *actions) add() float32{
  return a.x + a.y
}

func (a actions) sub() float32{
  return a.x - a.y
}

func main(){

  a := actions{x: 10.3, y: 30.99}
  fmt.Println(a.add())
  fmt.Println(a.sub())
}
