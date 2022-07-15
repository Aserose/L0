package streamer

import (
	"L0/internal/repository"
	"L0/pkg/customLogger"
)

const (
	chNamePurchase = `purchaseInfo`
	chNameErr      = `errorInfo`
)

type SendFunc func(msg []byte) error

type Streamer struct {
	chanPurchase
}

func NewStreamer(log customLogger.Logger, repo repository.Repository) Streamer {
	ns := newNatStream(log)
	return Streamer{
		chanPurchase: newChanPurchase(newNatChanParam(chNamePurchase, ns), repo, newChanErrInfo(newNatChanParam(chNameErr, ns), repo)),
	}
}

func (s Streamer) ChanPurchase() SendFunc {
	return s.chanPurchase.natsChan.Send
}
