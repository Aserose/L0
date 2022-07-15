package purchaseInfo

import (
	"L0/pkg/customLogger"
)

type natCache struct {
	nat map[int][]byte
	log customLogger.Logger
}

func newNatCache(log customLogger.Logger, data map[int][]byte) natCache {
	return natCache{
		nat: data,
		log: log,
	}
}

func (n natCache) put(id int, data []byte) {
	n.nat[id] = data
}

func (n natCache) get(id int) []byte {
	data, isExist := n.nat[id]
	if !isExist {
		return nil
	}
	return data
}

func (n natCache) delete(id int) {
	delete(n.nat, id)
}
