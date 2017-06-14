package main

import "fmt"


// defer: statment is defered until surrounding function returns

func main(){
  begin()
  defer end()
  loop()
  defer_demo_loop()

}

func begin(){
  fmt.Println("Begin")

}

func end(){
  fmt.Println("End")

}

func loop(){
  for i := 0; i<5; i++{
     fmt.Println(i)
    }
}

// defer calls are stored in a LIFO order

func defer_demo_loop(){
  for j := 0; j<5; j++{
    defer fmt.Println("Defered!:", j)
  }
}
