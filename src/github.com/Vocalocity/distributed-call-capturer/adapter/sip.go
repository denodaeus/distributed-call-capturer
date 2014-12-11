package adapter

import (
	"code.google.com/p/gopacket"
	"code.google.com/p/gopacket/dumpcommand"
	"code.google.com/p/gopacket/layers"
	"code.google.com/p/gopacket/pcap"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

// var iface = "eth0"
var promisc = "true"

type Sip struct {
	CallId string
	Port   int
}

func (s Sip) listIfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	for _, iface := range ifaces {
		wg.Add(1)

		go func(iface net.Interface) {
			defer wg.Done()
			if err := scan(&iface); err != nil {
				log.Printf("interface %v: %v", iface.Name, err)
			}

		}(iface)
	}
	wg.Wait()
}

func (s Sip) Trace(filter string) {

	log.Printf("Trace :: starting trace for filter=" + filter)
	var addr *net.IPNet
	if addrs, err := iface.Addrs(); err != nil {
		return err
	} else {
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					addr = &net.IPNet{
						IP:   ip4,
						Mask: ipnet.Mask[len(ipnet.Mask)-4:],
					}
					break
				}
			}
		}
	}
	if addr == nil {
		return fmt.Errorf("no IP network found")
	} else if addr.IP[0] == 127 {
		return fmt.Errorf("skipping localhost")
	} else if addr.Mask[0] != 0xff || addr.Mask[1] != 0xff {
		return fmt.Errorf("network mask is too large")
	}
	log.Printf("Using network range %v for interface %v", addr, iface.Name)

	handle, err := pcap.OpenLive(iface.Name, 65536, true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	stop := make(chan struct{})
	go readSIP(handle, iface, stop)
}

func (s Sip) constructDataSource(iface string) {
	log.Println("constructDataSource :: ")
	// packetSource := // construct packet source
	// for packet := range packetSource.Packets() {
	// 	handlePacket(packet)
	// }
}

func (s Sip) readSIP(handle *pcap.Handle, iface *net.Interface, stop chan struct{}) {
	src := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	in := src.Packets()
	for {
		var packet gopacket.Packet
		select {
		case <-stop:
			return
		case packet = <-in:
			udpLayer := packet.Layer(layers.LayerTypeUDP)
			if udpLayer == nil {
				continue
			}
			udp := udpLayer.(*layers.UDP)
			log.Printf("IP %v is at %v", net.IP(udp.SourceProtAddress), net.HardwareAddr(udp.SourceHwAddress))
		}
	}
}

func (s Sip) handlePacket(p Packet) {
	log.Println("handlePacket :: ")
}
