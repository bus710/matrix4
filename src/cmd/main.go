package main

import (
	"log"
	"sync"

	matrix "github.com/bus710/matrixd/src/matrixd/app/matrix"
	signal "github.com/bus710/matrixd/src/matrixd/app/signal"
	ui "github.com/bus710/matrixd/src/matrixd/app/ui"
)

func main() {
	log.Println("Hello!")

	windowCloseIndicator := make(chan bool, 1)
	waitInstance := sync.WaitGroup{}

	matrix.Matrix.Init(&waitInstance)
	ui.Window.Init(&waitInstance, windowCloseIndicator)
	signal.Signal.Init(&waitInstance, windowCloseIndicator)

	waitInstance.Add(1)
	go signal.Signal.Run()
	waitInstance.Add(1)
	go matrix.Matrix.Run()
	waitInstance.Add(1)
	go ui.Window.Run()

	waitInstance.Wait()

	log.Println("Bye!")
}
