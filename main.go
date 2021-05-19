package main

import (
	"gitee.com/zeehug/network-traffic-detection/cmd"
	"os"
)

func main() {
	if err := cmd.NetworkTrafficDetection("123"); err != nil {
		os.Exit(1)
	}
}
