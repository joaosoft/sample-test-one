package main

import (
	"fmt"

	sampletwo "github.com/joaosoft/sample-test-two"
)

func main() {
	fmt.Println(NewSampleTestOne())
	fmt.Println(sampletwo.NewSampleTestTwo())
}

func NewSampleTestOne() string {
	return "hello, i'm the sample-test-onw"
}
