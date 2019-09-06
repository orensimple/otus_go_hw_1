package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Printf("Remote ntp server error: %v", err)
	}
	fmt.Println(currentTime)
}
