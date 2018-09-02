package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"

	"sample-test-one"

	builder "github.com/joaosoft/builder"
)

func main() {
	fmt.Println("test of changes")
	err := builder.NewBuilder().Start(nil)
	if err != nil {
		fmt.Printf("error, %s", err)
	}

	fmt.Println(sample.NewSampleTestOne())

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	<-termChan
}
