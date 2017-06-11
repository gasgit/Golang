package main


import "fmt"

func main(){


  /*
    demo variable types, print functions, formating

  */

  /*
    var assingment, static, dynamic
  */
  var a float32  = 10.0
  b := 20

  var x,y,z = 22, 2.3, "two point four"

  

  /*
    Printf - formatting %f(decimal) %d(int) %T(type) %s(string)
    Println - print line

  */
  fmt.Printf("Value of a : %f\n",a)
  fmt.Printf("Value of b : %d\n", b)

  fmt.Printf("Type a: %T\n", a)
  fmt.Printf("Type b: %T\n", b)

  fmt.Printf("Value of x: %d\n", x)
  fmt.Printf("Value for y: %f\n", y)
  fmt.Printf("Value of z: %s\n", z)

  fmt.Printf("Type x: %T\n",x)
  fmt.Printf("Type y: %T\n", y)
  fmt.Printf("Type z: %T\n", z)
}
