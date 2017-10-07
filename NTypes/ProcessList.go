package NTypes

type ProcessList struct {
	Processes string
}

func (pl ProcessList) String() string {
	return string(pl.Processes)
}
