package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fschnko/scalc"
)

var (
	help   = flag.Bool("help", false, "display usage information and exit")
	verify = flag.Bool("verify", false, "verify an expression syntax and exit")
)

func main() {

	flag.Parse()

	if *help {
		usage()
		return
	}

	if len(os.Args) < 2 {
		log.Println("Wrong arguments count")
		usage()
		os.Exit(1)
	}

	expr, err := scalc.New(os.Args[len(os.Args)-1]).Process()
	if err != nil {
		log.Fatalf("Process: %s", err)
	}

	if *verify {
		fmt.Println("The expression has the correct syntax.")
		return
	}

	result, err := expr.Value()
	if err != nil {
		log.Fatalf("Calculate: %s", err)
	}

	for _, r := range result {
		fmt.Println(r)
	}
}

func usage() {
	fmt.Println(`scalc a basic sets calculator.

Usage:
	scalc [OPTIONS...] "<EXPRESSION>"

Options:`)
	flag.PrintDefaults()
	fmt.Println(`
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

Example: 
	scalc "[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"
	`)
}
