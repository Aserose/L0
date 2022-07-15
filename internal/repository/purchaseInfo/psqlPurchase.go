package purchaseInfo

import (
	"L0/internal/datastruct"
	"L0/pkg/customLogger"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type cache interface {
	put(id int, data []byte)
	get(id int) []byte
	delete(id int)
}

type PsqlPurchase struct {
	db *sqlx.DB
	cache
	log customLogger.Logger
}

func NewPsqlPurchase(db *sqlx.DB, log customLogger.Logger) PsqlPurchase {
	return PsqlPurchase{
		db:    db,
		cache: newNatCache(log, getPurchaseData(db, log)),
		log:   log,
	}
}

func (p PsqlPurchase) Put(data datastruct.Order) (id int) {
	a, _ := json.Marshal(data)

	if err := p.db.QueryRow("INSERT INTO purchase (info) VALUES($1) RETURNING id", a).Scan(&id); err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
	}

	if id > 0 {
		p.cache.put(id, a)
	}

	return
}

func (p PsqlPurchase) GetCached(id int) (res []byte) { return p.cache.get(id) }

func (p PsqlPurchase) GetCount() (count int) {
	if err := p.db.QueryRow("SELECT count(*) FROM purchase").Scan(&count); err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
	}
	return
}

func (p PsqlPurchase) Get(id int) (res []byte) {
	if err := p.db.Get(&res, "SELECT info FROM purchase WHERE id=$1", id); err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
	}
	return
}

func (p PsqlPurchase) Delete(id int) {
	_, err := p.db.Exec("DELETE FROM purchase WHERE id=$1", strconv.Itoa(id))
	if err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
	}
	p.cache.delete(id)
}

func getPurchaseData(db *sqlx.DB, log customLogger.Logger) map[int][]byte {
	r := make(map[int][]byte)
	id := 0
	data := []byte{}

	rows, err := db.Query("SELECT * FROM purchase")
	if err == nil {
		for rows.Next() {
			if err := rows.Scan(&id, &data); err != nil {
				log.Error(log.CallInfoStr(), err.Error())
			}
			r[id] = data
		}
	}

	return r
}
