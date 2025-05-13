package main

import (
	"github.com/mikelinesta/turtle/internal/turtle"
)

// Main function in charge of reading input from stdin and calling
// the Turtle package
func main(){

  instructions := make(map[string]turtle.Action)
  instructions["p"] = func (penNumber int){println("select pen ",penNumber)}
  instructions["d"] = func (_ int){println("pen down")}
  instructions["n"] = func (distance int){println("draw north ",distance, "cm")}
  instructions["s"] = func (distance int){println("draw south",distance, "cm")}
  instructions["e"] = func (distance int){println("draw east",distance, "cm")}
  instructions["w"] = func (distance int){println("draw west",distance, "cm")}
  instructions["u"] = func (_ int){println("pen up")}
 
  t := turtle.NewTurtle(
    instructions,
  )
  // Start reading input

  for{
    err := t.Read()
    if err != nil{
      println(err.Error())
    }
  }

}

