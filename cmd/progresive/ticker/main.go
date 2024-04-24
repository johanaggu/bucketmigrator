package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	perSecond := os.Getenv("PER_SECOND")

	fmt.Println("-----111", perSecond, "11------")
	div, _ := strconv.Atoi(perSecond)
	ticker := time.NewTicker(time.Minute / time.Duration(div))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			hello()
			fmt.Println("-------2222", "222------")
			perSecond = os.Getenv("PER_SECOND")

		}
	}

}

func hello() {
	//time.Sleep(time.Second * 3)
	fmt.Println("ok")

}
