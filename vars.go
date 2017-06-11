package main
import "fmt"

func main(){


  var a float32  = 10.0
  b := 20

  var x,y,z = 22, 2.3, "two point four"


  fmt.Println(a)
  fmt.Println(b)

  fmt.Printf("A = %T\n", a)
  fmt.Printf("B = %T\n", b)

  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)

  fmt.Printf("Type x: %T\n",x)
  fmt.Printf("Type y: %T\n", y)
  fmt.Printf("Type z: %T\n", z)
}
