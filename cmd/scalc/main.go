package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fschnko/scalc"
)

func main() {

	if len(os.Args) != 2 {
		log.Println("Wrong arguments count")
		usage()
		os.Exit(1)
	}

	calc, err := scalc.New(os.Args[1])
	if err != nil {
		log.Fatalf("Create a new instance of calc: %s", err)
	}

	result, err := calc.Calculate()
	if err != nil {
		log.Fatalf("Calculate expression: %s", err)
	}

	for _, r := range result {
		fmt.Println(r)
	}
}

func usage() {
	fmt.Println(`
	scalc a basic sets calculator.

	Grammar of calculator is given:
	expression := “[“ operator sets “]”
	sets := set | set sets
	set := file | expression
	operator := “SUM” | “INT” | “DIF”

	Each file should contain sorted integers, one integer in a line.
	Meaning of operators:
	SUM - returns union of all sets
	INT - returns intersection of all sets
	DIF - returns difference of first set and the rest ones

	Usage: ./scalc "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"
	`)
}
