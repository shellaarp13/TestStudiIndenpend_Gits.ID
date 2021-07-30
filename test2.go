package main
import "fmt"
import "regexp"

func main ()  {
  var email string

  fmt.Println ("Masukkan email:")
  fmt.Scan (&email)
  if len(email) <= 20{
    checkEmail := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+[.id][.co.id]$`)
  	 if checkEmail.MatchString(email) {
  			fmt.Println("Email Valid")
  		} else {
  			fmt.Println("Email Tidak Valid")
  		}
  }
}
