package main

import "fmt"

type Node int

const (
	Node0 Node = iota
	Node1
	Node2
	Node3
)

type Automaton struct {
	Current Node
	IsValid bool
}

func (a *Automaton) Transition(input *string, output *string) (bool, string) {
	if len(*input) == 0 {
		return a.IsValid, *output
	}

	next := (*input)[0]
	*input = (*input)[1:]
	success, successor := a.Current.Follow(next)

	if !success {
		return false, "invalid transition"
	}

	a.Current = successor
	a.IsValid = a.Current == Node2
	*output = *output + string(next)

	return a.Transition(input, output)
}

func (n *Node) Follow(next byte) (bool, Node) {
	switch *n {
	case Node0:
		switch next {
		case 'a':
			return true, Node1
		case 'b':
			return true, Node2
		default:
			return false, *n
		}

	case Node1:
		switch next {
		case 'd':
			return true, Node2
		case 'c':
			return true, Node3
		default:
			return false, *n
		}

	case Node2:
		switch next {
		case 'f':
			return true, Node3
		default:
			return false, *n
		}

	case Node3:
		switch next {
		case 'e':
			return true, Node1
		default:
			return false, *n
		}
	default:
		return false, *n
	}
}

func e1_3() {
	a := Automaton{Current: Node0, IsValid: true}
	input := "bfecedfecedfed"
	output := ""

	valid, output := a.Transition(&input, &output)
	if valid {
		fmt.Printf("Valid: %v\n", output)
	} else {
		fmt.Printf("Invalid: %v\n", output)
	}
}
