package main

import (
	"fmt"
	"time"
)

func main() {
	throttle := time.Tick(time.Minute / 60)

	for next := range throttle {
		<-throttle
		fmt.Printf("Hello %v :: %v \n", next.Minute(), next.Second())
	}

}
