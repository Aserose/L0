package purchaseInfo

import (
	"L0/internal/datastruct"
	"L0/internal/repository"
)

type PurchaseInfo struct {
	repo repository.Repository
}

func NewPurchaseInfo(repo repository.Repository) PurchaseInfo {
	return PurchaseInfo{
		repo: repo,
	}
}

func (p PurchaseInfo) Put(order datastruct.Order) int {
	return p.repo.PurchaseInfo.Put(order)
}

func (p PurchaseInfo) GetCached(id int) []byte {
	return p.repo.GetCached(id)
}

func (p PurchaseInfo) Get(id int) []byte {
	return p.repo.Get(id)
}

func (p PurchaseInfo) GetCount() int {
	return p.repo.GetCount()
}
