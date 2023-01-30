package main

import (
	"github.com/thesky9531/network-traffic-detection/cmd"
	"os"
)

func main() {
	if err := cmd.NetworkTrafficDetection("123"); err != nil {
		os.Exit(1)
	}
}
