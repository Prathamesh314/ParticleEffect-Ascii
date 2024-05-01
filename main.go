
package main

import (
	"ascii/particles"
	"fmt"
	"time"

)

func main() {
	coffee := particles.NewCoffee(5, 3)
	timer := time.NewTicker(100 * time.Millisecond)
	coffee.Start()

	for {
		<-timer.C
    fmt.Print("\033[H\033[2J")
		coffee.Update()
		fmt.Println(coffee.Display())
	}
}

