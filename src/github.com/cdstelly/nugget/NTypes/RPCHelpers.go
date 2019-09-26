package NTypes

type NugArg struct {
	TheData []byte
	Inode   string
}

type NuggetPCAPArgs struct {
	TheData      []byte
	PcapLocation string
	Filters      []Filter
}

type NugData int
