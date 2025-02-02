// "Error values in Go aren't special. They're just values like
// any other, and so you have the entire language at your disposal" - Rob Pike
package main

import (
	"log"
	"errors"
	)


func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	// implementation
	return 42, nil
}