package main

import (
	"dontstop/cicle"
	"dontstop/world"
	"dontstop/wsServer"
	"fmt"
)

// "dontstop/world"

func main() {

	go cicle.Frame(world.World, 6)
	wsServer.Start()

	defer fmt.Println("Program has been finished")
}
