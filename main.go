package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println("problem with remote ntp server, use server time")
	}
	fmt.Println(currentTime)
}
