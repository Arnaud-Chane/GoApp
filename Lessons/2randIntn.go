package main

import (
  "fmt"
  "math/rand"
)

func main (){

  var dice1 = rand.Intn(6)+1
  var dice2 = rand.Intn(6)+1
  fmt.Printf("First dice is %v ", dice1)
  fmt.Printf("and the second one is %v.", dice2)

}
