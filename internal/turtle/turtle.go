package turtle

import (
  "bufio"
  "os"
)

type Turtle struct {
  instructionSet []Instruction
}

type NewTurtle(set []Instruction) Turtle {
  return Turtle{
    instructionSet = set
  }
}

type Action func(modifier int)

type Instruction struct {
  Command string
  action Action 
}


// Receives a raw string and returns an Instruction
func (t Turtle) parse(query string) (Instruction, error) {

}

// Receives an instruction and calls its assigned action 
func (t Turtle) execute(instr Instruction) (error) {

}

// Reads from stdin and calls executions
func (t Turtle) read(){
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    line := input.Text() 
  }
}

// -- Instruction Set --

