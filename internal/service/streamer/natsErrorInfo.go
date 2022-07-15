package streamer

import (
	"L0/internal/repository"
	"github.com/nats-io/stan.go"
)

const (
	errTypeInvalidMsg   = `invalid msg`
	errTypeInvalidOrder = `invalid order`
)

type chanErrInfo struct {
	repo repository.Repository
	natsChan
}

func newChanErrInfo(p natChanParam, repo repository.Repository) chanErrInfo {
	c := chanErrInfo{
		repo: repo,
		natsChan: natsChan{
			name:      p.name,
			natStream: p.n,
		},
	}
	c.natsChan.addChannel(p.name, c.handleErrInfo)

	return c
}

func (e chanErrInfo) handleErrInfo(msg *stan.Msg) { e.repo.ErrorCache.Put(string(msg.Data)) }
