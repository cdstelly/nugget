package NTypes

import (
	"github.com/google/gopacket"
	"encoding/json"
	"fmt"
)

type NPacket struct {
	Pkt gopacket.Packet
}

func (p NPacket) String() string {
	applicationLayer := p.Pkt.ApplicationLayer()
	if applicationLayer != nil {
		return string(applicationLayer.Payload()) + "\n"
	}
	return "" //todo: perhaps return a 'packet unable to be decoded' string
}

func (p NPacket) JSON() string {
	jsonEncoding, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return string(jsonEncoding)
}
//todo function to interpret as http and print the results of specific fields, such as host and cookie
