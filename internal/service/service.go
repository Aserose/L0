package service

import (
	"L0/internal/datastruct"
	"L0/internal/repository"
	errorInfo "L0/internal/service/errorInfo"
	"L0/internal/service/purchaseInfo"
	"L0/internal/service/streamer"
	"L0/pkg/customLogger"
)

type Streamer interface {
	ChanPurchase() streamer.SendFunc
}

type PurchaseInfo interface {
	Put(data datastruct.Order) int
	Get(id int) []byte
	GetCount() int
	GetCached(id int) []byte
}

type ErrorInfo interface {
	GetLast() string
}

type Service struct {
	Streamer
	PurchaseInfo
	ErrorInfo
}

func NewService(log customLogger.Logger, repo repository.Repository) Service {
	return Service{
		Streamer:     streamer.NewStreamer(log, repo),
		PurchaseInfo: purchaseInfo.NewPurchaseInfo(repo),
		ErrorInfo:    errorInfo.NewErrInfo(repo),
	}
}
