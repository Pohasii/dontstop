package main

import (
	"dontstop/cycle"
	"dontstop/world"
	"dontstop/wsServer"
	"fmt"
)

// "dontstop/world"

func main() {

	go cycle.Frame(world.World, 6)
	wsServer.Start()

	defer fmt.Println("Program has been finished")
}
