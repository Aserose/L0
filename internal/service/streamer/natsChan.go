package streamer

type natsChan struct {
	name string
	natStream
}

func (s natsChan) getName() string {
	return s.name
}

type natChanParam struct {
	name string
	n    natStream
}

func newNatChanParam(name string, n natStream) natChanParam {
	return natChanParam{
		name: name,
		n:    n,
	}
}

func (s natsChan) Send(msg []byte) error { return s.natStream.sendMsg(s.name, msg) }
