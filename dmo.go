package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(time.Second * 5)
	for i := 0; i < 3; i++ {
		fmt.Println(<-t.C)
	}

}
