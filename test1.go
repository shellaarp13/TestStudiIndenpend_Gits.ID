package main
import "fmt"

func main(){
  var input int

  fmt.Println ("Masukkan input:")
  fmt.Scanln (&input)
  if (input %3 == 0) && (input %5 == 0) {
    fmt.Println ("Hello World")
  } else if input %5 == 0 {
    fmt.Println ("World")
  } else if input %3 == 0 {
    fmt.Println ("Hello ")
  } else{
    fmt.Println ("Error")
  }
 }
