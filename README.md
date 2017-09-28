Nugget
===================

Nugget is a domain specific language ([DSL](https://en.wikipedia.org/wiki/Domain-specific_language)) for Digital Forensics. 

Building
----------
Nugget does not (currently) conform to the Go standard of buildling (i.e., keeping all code under ./src/github.com/cdstelly/code). 

To build from source: 
```
git clone https://github.com/cdstelly/nugget.git
cd nugget
export GOPATH=`pwd`
go get "github.com/antlr/antlr4/runtime/Go/antlr"
go get "github.com/google/gopacket"
go get "github.com/google/gopacket/pcap"
go build nugtest.go
```

Using
------------
After either building from source or downloading binaries:

```
[arch@arch nugget]$ ./nugtest -h
Usage of ./nugtest:
  -input string
    	Path to input (default "input.nug")
```

Examples
--------------

```
nugget> myhashes = "file.dd" | extract  as ntfs | filter filename == "*.pdf",ctime>"01/01/01" | md5
nugget> print(myhashes)
    [{c10c4d40735cc699bd16d4d18c2c6b09} {cc285f386f167c2206dd9ff6546dcd0a} ... }]

nugget> mypcap = "G:\school\sample.pcap" | extract as pcap
nugget> myhttp = mypcap | filter packetfilter == "tcp and port 80 and http"
nugget> print(myhttp)
    [ GET /site=0000127709/mnum=0000162763/genr=1/logs=0/mdtm=1077726643/bins=1 HTTP/1.1
    User-Agent: Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0) Opera 7.11  [en]
    Host: opera2-servedby.advertising.com
    Accept: application/x-shockwave-flash,text/xml,application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,video/x-mng,image/png,image/jpeg,image/gif;q=0.2,text/css,*/*;q=0.1
    Accept-Language: en
    ...

```


Expanding
--------------

Bugs
---------------
Please use the github issue tracker. 
