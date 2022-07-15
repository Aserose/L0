package streamer

import (
	"L0/internal/repository"
	"L0/pkg/customLogger"
	"github.com/nats-io/stan.go"
)

const (
	defaultClusterName = "test-cluster"
	defaultClientID    = "NAEZVQ2VNL6G635GTQUGEUXCUWZ7USZFJTTNLEHDLG75HNGU5RE3CQXF"
)

type natStream struct {
	conn stan.Conn
	repo repository.Repository
	log  customLogger.Logger
}

func newNatStream(log customLogger.Logger) natStream {
	c, err := stan.Connect(defaultClusterName, defaultClientID)
	if err != nil {
		log.Error(log.CallInfoStr(), err.Error())
	}

	return natStream{
		conn: c,
		log:  log,
	}
}

func (ns natStream) addChannels(chns map[string]stan.MsgHandler) {
	for k := range chns {
		ns.addChannel(k, chns[k])
	}
}

func (ns natStream) addChannel(subjectName string, handleFunc stan.MsgHandler) {
	_, err := ns.conn.Subscribe(subjectName, handleFunc)
	if err != nil {
		ns.log.Error(ns.log.CallInfoStr(), err.Error())
	}
}

func (ns natStream) close() {
	if err := ns.conn.Close(); err != nil {
		ns.log.Error(ns.log.CallInfoStr(), err.Error())
	}
}

func (ns natStream) sendMsg(subjectName string, msg []byte) error {
	return ns.conn.Publish(subjectName, msg)
}
