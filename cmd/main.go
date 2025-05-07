package main

import (
   "turtle"
)

// Main function in charge of reading input from stdin and calling
// the Turtle package
func main(){

  t := NewTurtle()
  // P 2
  i1 := Instruction{
    Command: p,
    Action: print("select pen " + Modifier) 
  }
  i2 := Instruction{
    Command: d,
    Action: print("pen down") 
  }
  i3 := Instruction{
    Command: w,
    Action: print("draw west 2cm") 
  }
  i4 := Instruction{
    Command: n,
    Action: print("then north 1") 
  }
  i5 := Instruction{
    Command: e,
    Action: print("then east 2") 
  }
  i6 := Instruction{
    Command: s,
    Action: print("then back south") 
  }
  i7 := Instruction{
    Command: p,
    Action: print("pen up") 
  }
  t.read()

}

