package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	defaultSpan = 5
	minSpan     = 3
	maxSpan     = 20

	defaultStart = 0
	defaultCount = 10

	separator = "\n"
)

func main() {
	var (
		start, count, span int
	)

	flag.IntVar(&start, "start", defaultStart, "starting point")
	flag.IntVar(&count, "count", defaultCount, "max count of random integers")
	flag.IntVar(&span, "span", defaultSpan, "average step between values")
	flag.Parse()

	if span < minSpan || span > maxSpan {
		span = defaultSpan
	}

	for i, v := 0, start; i < count; i++ {
		v += jitter(span)
		fmt.Print(v, separator)
	}
}

var r = rand.NewSource(time.Now().Unix())

func jitter(span int) int {
	return int((r.Int63() % int64(span))) + span
}
