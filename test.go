package main

import (
  "fmt"
  "math/rand"
)

func main(){
  
  var dice1 = rand.Intn(6)+1
  var dice2 = rand.Intn(6)+1

  fmt.Println("First dice is " + dice1 + " and the second is " + dice2 + ".")

}
