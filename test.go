<<<<<<< HEAD
=======
package main

import (
  "fmt"
  "math/rand"
)

func main(){
  
  var dice1 = rand.Intn(6)+1
  var dice2 = rand.Intn(6)+1

  fmt.Println("First dice is %v and the second is %v." , dice1, dice2)

  var name string
  fmt.Print("You are ")
  fmt.Scanln(&name)

}
>>>>>>> 5029223 (changing lessons by name more explicit)
