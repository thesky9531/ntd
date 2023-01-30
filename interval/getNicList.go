package interval

import (
	"fmt"
	"github.com/google/gopacket/pcap"
)

func GetNicList() {
	fmt.Println("test")
	devices, err := pcap.FindAllDevs()
	if err != nil {
		fmt.Println(err)
	}
	for _, device := range devices {
		for _, address := range device.Addresses {
			fmt.Println(device.Name, " - ", address.IP.String())
		}
	}
}
