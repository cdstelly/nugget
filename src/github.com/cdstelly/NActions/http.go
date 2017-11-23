package NActions

import (
	"../NTypes"
	"bufio"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"io"
	"net/http"
	"reflect"

)

type HTTPAction struct {
	executed  bool
	dependsOn BaseAction
	filters   []NTypes.Filter

	Results []NTypes.HTTP
}

func (na *HTTPAction) BeenExecuted() bool {
	return na.executed
}

func (na *HTTPAction) DependencySatisfied() bool {
	if na.dependsOn == nil {
		return true
	}
	return na.dependsOn.BeenExecuted()
}

func (na *HTTPAction) SetDependency(action BaseAction) {
	na.dependsOn = action
}

func (thisAction *HTTPAction) Execute() {
	// Set up assembly streams, reference: https://godoc.org/github.com/google/gopacket/tcpassembly
	streamFactory := &httpStreamFactory{}
	streamPool := tcpassembly.NewStreamPool(streamFactory)
	assembler := tcpassembly.NewAssembler(streamPool)

	operateOn := thisAction.dependsOn.GetResults()
	if _, ok := operateOn.([]NTypes.NPacket); ok {
		var packets []NTypes.NPacket
		packets = operateOn.([]NTypes.NPacket)

		unusablePrinted := false
		for _, packet := range packets {
			pkt := packet.Pkt
			if pkt.NetworkLayer() == nil || pkt.TransportLayer() == nil || pkt.TransportLayer().LayerType() != layers.LayerTypeTCP {
				//log.Println("Unusable packet")
				if unusablePrinted == false {
					fmt.Println("Warning: unusable packets given to HTTP extractor")
					unusablePrinted = true
				}
				continue
			}
			tcp := pkt.TransportLayer().(*layers.TCP)
			assembler.AssembleWithTimestamp(pkt.NetworkLayer().NetworkFlow(), tcp, pkt.Metadata().Timestamp)
			//fmt.Println(tcp.Payload)
		}
		for _,r := range myrequests {
			thisAction.Results = append(thisAction.Results, nuggetHTTPFromRequest(r))
		}
//		fmt.Println("done len:" , len(myrequests))
	} else {
		fmt.Println("incorect type, found ", reflect.TypeOf(operateOn))
	}
	thisAction.executed = true
}

func (na *HTTPAction) GetResults() interface{} {
	if !na.BeenExecuted() {
		na.Execute()
	}
	return na.Results
}

func (na *HTTPAction) SetFilters(filters []NTypes.Filter) {
	//TODO: investigate if resetting executed status will be a problem:
	na.executed = false
	na.filters = filters
}

// Packet -> HTTP helpers, taken from https://github.com/google/gopacket/blob/master/examples/httpassembly/main.go
// httpStreamFactory implements tcpassembly.StreamFactory
type httpStreamFactory struct{}

// httpStream will handle the actual decoding of http requests.
type httpStream struct {
	net, transport gopacket.Flow
	r              tcpreader.ReaderStream
}
var myrequests []http.Request  //todo find best way to pass in reference to instance's results variable
func (h *httpStreamFactory) New(net, transport gopacket.Flow) tcpassembly.Stream {
	hstream := &httpStream{
		net:       net,
		transport: transport,
		r:         tcpreader.NewReaderStream(),
	}
	go hstream.run() // Important... we must guarantee that data from the reader stream is read.

	// ReaderStream implements tcpassembly.Stream, so we can return a pointer to it.
	return &hstream.r
}

func nuggetHTTPFromRequest(r http.Request) NTypes.HTTP {
	var httpHolder NTypes.HTTP
	httpHolder.Host = r.Host
	r.Body.Read(httpHolder.Data)
	httpHolder.Method = r.Method
	httpHolder.RawRequest = r
	return httpHolder
}

func (h *httpStream) run() {
	buf := bufio.NewReader(&h.r)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			// We must read until we see an EOF... very important!
			return
		} else if err != nil {
			//todo investigate how errors can occur, invalid data for example
			//log.Println("Error reading stream", h.net, h.transport, ":", err)
			continue
		} else {
			req.Body.Close()
			//fmt.Println("request header: ", req.Host)
			myrequests = append(myrequests,*req)
		}
		//fmt.Println("HTTP Request host: " + req.Host)
	}
}
