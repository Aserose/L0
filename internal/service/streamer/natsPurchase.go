package streamer

import (
	"L0/internal/datastruct"
	"L0/internal/repository"
	"encoding/json"
	"github.com/nats-io/stan.go"
)

type chanPurchase struct {
	maxMsgSize int
	repo       repository.Repository
	natsChan
	chanErrInfo
}

func newChanPurchase(p natChanParam, repo repository.Repository, errChan chanErrInfo) chanPurchase {
	c := chanPurchase{
		maxMsgSize:  15,
		repo:        repo,
		chanErrInfo: errChan,
		natsChan: natsChan{
			name:      p.name,
			natStream: p.n,
		},
	}
	c.natsChan.addChannel(p.name, c.handlePurchase)

	return c
}

func (o chanPurchase) handlePurchase(msg *stan.Msg) {
	order := datastruct.Order{}
	if err := json.Unmarshal(msg.Data, &order); err != nil {
		o.chanErrInfo.natsChan.Send(o.newErrorMsg(errTypeInvalidMsg, string(msg.Data)))
		return
	}

	switch order.IsValid() {
	case true:
		o.repo.PurchaseInfo.Put(order)
	case false:
		o.chanErrInfo.natsChan.Send(o.newErrorMsg(errTypeInvalidOrder, order.Entry))
	}
}

func (o chanPurchase) newErrorMsg(errorType string, data string) []byte {
	byteValue, _ := json.Marshal(datastruct.ErrorInfo{
		Type: errorType,
		Data: o.cutStr(data),
	})
	return byteValue
}

func (o chanPurchase) cutStr(str string) string {
	if len(str) > o.maxMsgSize {
		return str[:o.maxMsgSize]
	}
	return str
}
