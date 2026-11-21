package main

import "strawberryInterpreter/interpreter"

func runSource(source string) (string, error) {
	return interpreter.Run(source)
}
