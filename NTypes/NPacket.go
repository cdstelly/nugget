package NTypes

import (
	"github.com/google/gopacket"
)

type NPacket struct {
	Pkt gopacket.Packet
}

func (p NPacket) String() string {
	applicationLayer := p.Pkt.ApplicationLayer()
	if applicationLayer != nil {
		return string(applicationLayer.Payload()) + "\n"
	}
	return ""  //todo: perhaps return a 'packet unable to be decoded' string
}