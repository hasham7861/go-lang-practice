package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota // creates auto indexes from 0 to n for following values
  StateConnected
  StateError
  StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

//  you can overwrite the string of given enum with map
func (ss ServerState) String() string {
    return stateName[ss]
}

func main () {

	ns := transition(StateIdle)
	fmt.Println(ns)


	ns2 := transition(StateConnected)
	fmt.Println(ns2)
	
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
      return StateConnected
    case StateConnected, StateRetrying:
      return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}