package main

import "github.com/mikelinesta/turtle/internal/turtle"

// Main function in charge of reading input from stdin and calling
// the Turtle package
func main(){

  instructions := make(map[string]turtle.Action)
  instructions["p"] = func (penNumber int){println("select pen ",penNumber)}
  instructions["d"] = func (_ int){println("pen down")}
 
  t := turtle.NewTurtle(
    instructions,
  )
  // Start reading input
  t.Read()

}

