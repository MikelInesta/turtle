package turtle

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Turtle struct {
  instructionSet map[string]Action
}

func NewTurtle(instructions map[string]Action) Turtle {
  return Turtle{
    instructionSet: instructions,
  }
}

type Instruction struct {
  command string 
  modifier int
}

// The actual function to be called to run an instruction
type Action func(modifier int)


// Receives a raw string and returns an Instruction 
func (t Turtle) parse(query string) (*Instruction, error) {
  trimmed := strings.TrimSpace(query)
  split := strings.Split(trimmed, " ")
  if (len(split) < 1){
    return nil, errors.New("query must have at least one argument")
  }
  command := split[0]
  lowerCaseCommand := strings.ToLower(command)
  if (len(split) > 1){
    modifier, err := strconv.Atoi(split[1])
    if err != nil {
      return nil, err
    }

    return &Instruction{
      command: lowerCaseCommand,
      modifier: modifier,
    }, nil
  }

  return &Instruction{
    command: lowerCaseCommand,
  }, nil
}


// Reads from stdin and executes actions based on commands 
func (t Turtle) Read() error {
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {

    line := input.Text() 

    instruction, err := t.parse(line)
    if (err != nil){
      return err
    }

    action, ok := t.instructionSet[instruction.command]
    if(!ok){
      return errors.New("unsupported command")
    }

    action(instruction.modifier)
  }
  return nil
}