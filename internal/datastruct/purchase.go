package datastruct

import (
	"sync"
	"time"
)

type Order struct {
	CustomerID        string    `json:"customer_id"`
	DateCreated       time.Time `json:"date_created"`
	DeliveryService   string    `json:"delivery_service"`
	Entry             string    `json:"entry"`
	InternalSignature string    `json:"internal_signature"`
	Locale            string    `json:"locale"`
	OofShard          string    `json:"oof_shard"`
	OrderUid          string    `json:"order_uid"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	TrackNumber       string    `json:"track_number"`
	Delivery          `json:"delivery"`
	Items             []Item `json:"items"`
	Payment           `json:"payment"`
}

func (o Order) IsValid() bool {
	reqFields := []string{
		o.CustomerID,
		o.OrderUid,
		o.TrackNumber,
	}

	for _, req := range reqFields {
		if req == `` {
			return false
		}
	}

	if !o.Payment.IsValid() {
		return false
	}

	for _, item := range o.Items {
		if !item.isValid() {
			return false
		}
	}

	return true
}

type Payment struct {
	Amount       int    `json:"amount"`
	Bank         string `json:"bank"`
	Currency     string `json:"currency"`
	CustomFee    int    `json:"custom_fee"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	PaymentDt    int    `json:"payment_dt"`
	Provider     string `json:"provider"`
	RequestID    string `json:"request_id"`
	Transaction  string `json:"transaction"`
}

func (p Payment) IsValid() bool {
	reqFields := []string{
		p.Transaction,
		p.Currency,
		p.Provider,
		p.Bank,
	}

	for _, req := range reqFields {
		if req == `` {
			return false
		}
	}

	return true
}

type Delivery struct {
	Address string `json:"address"`
	City    string `json:"city"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Region  string `json:"region"`
	Zip     string `json:"zip"`
}

type Item struct {
	Brand       string `json:"brand"`
	ChrtID      int    `json:"chrt_id"`
	Name        string `json:"name"`
	NmID        int    `json:"nm_id"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	Status      int    `json:"status"`
	TotalPrice  int    `json:"total_price"`
	TrackNumber string `json:"track_number"`
}

func (i Item) isValid() bool {
	result := true
	wg := sync.WaitGroup{}

	reqFieldsStr := []string{
		i.Rid,
	}
	reqFieldsInt := []int{
		i.ChrtID,
	}

	wg.Add(1)
	go func() {
		wg.Done()
		for _, req := range reqFieldsStr {
			if req == `` && result == true {
				result = false
			} else {
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		for _, req := range reqFieldsInt {
			if req == 0 && result == true {
				result = false
			} else {
				return
			}
		}
	}()
	wg.Wait()

	return result
}
