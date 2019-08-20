package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	exitChannel := make(chan int)

	go webServer(exitChannel)

	for {
		select {
		case msg := <-exitChannel:
			{
				switch msg {
				case 0:
					fmt.Println("All tasks finished")
					os.Exit(0)
				default:
					fmt.Println("Message" + strconv.Itoa(msg))
				}
			}
		}
	}
}
