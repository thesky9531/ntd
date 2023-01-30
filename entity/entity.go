package entity

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/thesky9531/network-traffic-detection/config"
	//"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"log"
	"time"
)

func ini() {

}

func FlowCollection(cf *config.Conf) {
	handle, err := pcap.OpenLive(
		cf.Ntd.Input.NicName,
		int32(65535),
		true,
		pcap.BlockForever,
	)
	if err != nil {
		log.Fatal("pcap open live is failed: ", err)
	}

	// 设置过滤器
	filter := GetFilter()
	filter = ""
	if err := handle.SetBPFFilter(filter); err != nil {
		log.Fatal("set bpf filter is failed: ", err)
	}
	defer handle.Close()

	//start get packet
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packageSource.NoCopy = true

	for packet := range packageSource.Packets() {
		//if len(packageSource.Packets())>100 {
		//	fmt.Println(len(packageSource.Packets()))
		//}
		/*
			正式代码
		*/
		//1.先将packet进行简单的信息过滤，根据不同的类型，传递到对应的协程里面分析。

		/* part 1 除TCP，UDP以外的 数据包如何捕获*/
		fmt.Print(".")
		aa := packet.Layers()
		if packet.TransportLayer() == nil && aa[1].LayerType().String() != "ARP" {
			//if packet.ApplicationLayer()==nil{
			//	continue
			//}else{
			//	if packet.ApplicationLayer().LayerType().String()=="Payload"{
			//		continue
			//	}
			//}
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("linkLayer", packet.LinkLayer().LinkFlow(), "\n", packet.LinkLayer().LayerType().String())
			fmt.Println("networkLayer", packet.NetworkLayer())
			fmt.Println("transportLayer", packet.TransportLayer())
			if packet.ApplicationLayer() != nil {
				fmt.Println("applicationLayer", packet.ApplicationLayer().LayerType().String())
			}
			fmt.Println(packet)
			//fmt.Println(packet.String())
			aa := packet.Layers()
			fmt.Println("aa[2]: ", string(aa[1].LayerType().String()))

		}

		/*
			part 2.0  基于协议种类划分

			if aa:=packet.Layers();aa[1].LayerType().String()=="ARP"{
				fmt.Println(packet)
				arp := packet.Layer(layers.LayerTypeARP)
				aaa := arp.(*layers.ARP)
				fmt.Println("protocol:",aaa.Protocol)
				fmt.Println("arp: ",arp)
				fmt.Println("LayerType:",arp.LayerType().String())
				fmt.Println("LaerContents:",string(arp.LayerContents()))
				fmt.Println("LaerPayLoad:",string(arp.LayerPayload()))
			}*/

		/*
			if packet.TransportLayer().LayerType() == layers.LayerTypeTCP{
				//fmt.Println("发现一个tcp数据包：")
				//fmt.Println(packet)
				//time.Sleep(time.Minute/10)
				//fmt.Print(".")
				//go handle11(packet)

				//time.Sleep(time.Microsecond)
			}else{
				fmt.Println(packet)
				time.Sleep(time.Second)
			}*/
	}

}

var goroutineNum int64

func handle11(packet gopacket.Packet) {
	goroutineNum++
	//fmt.Print(".")
	fmt.Println("goroutineNum: ", goroutineNum)
	time.Sleep(time.Minute / 60)
	goroutineNum--
}

func GetFilter() string {
	return ""
}
