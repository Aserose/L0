package repository

import (
	"L0/internal/datastruct"
	"L0/internal/repository/db"
	er "L0/internal/repository/errorInfo"
	"L0/internal/repository/purchaseInfo"
	"L0/pkg/customLogger"
)

type ErrorCache interface {
	Put(error string)
	GetLast() string
}

type PurchaseInfo interface {
	Put(data datastruct.Order) (id int)
	Get(id int) (res []byte)
	GetCached(id int) (res []byte)
	GetCount() int
	Delete(id int)
}

type Repository struct {
	ErrorCache
	PurchaseInfo
}

func NewRepository(log customLogger.Logger) Repository {
	return Repository{
		ErrorCache:   er.NewErrCache(log, make([]string, 1)),
		PurchaseInfo: purchaseInfo.NewPsqlPurchase(db.NewPsql(log), log),
	}
}
