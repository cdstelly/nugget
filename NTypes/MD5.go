package NTypes

type MD5 struct {
	Digest string
	HashOf interface{}
}

func (m MD5) String() string {
	return string(m.Digest)
}
