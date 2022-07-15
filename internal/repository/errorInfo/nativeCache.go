package errorInfo

import (
	"L0/pkg/customLogger"
)

const (
	msgNoError = `none`
	empty      = ``
)

type NatCache struct {
	nat []string
	log customLogger.Logger
}

func NewErrCache(log customLogger.Logger, nat []string) NatCache {
	return NatCache{
		nat: nat,
		log: log,
	}
}

func (n NatCache) Put(error string) {
	if n.nat[len(n.nat)-1] != empty {
		n.clearAll()
		n.nat[0] = error
	} else {
		for i, element := range n.nat {
			ind := i
			if element == `` {
				n.nat[ind] = error
				return
			}
		}
	}
}

func (n NatCache) GetLast() string {
	result := msgNoError
	if n.nat[0] != empty {
		result = n.nat[0]
	}
	for i, element := range n.nat {
		ind := i - 1
		if element == empty {
			if i > 0 && n.nat[ind] != empty {
				result = n.nat[ind]
			} else {
				break
			}
		}
	}

	return result
}

func (n NatCache) clearAll() {
	for i, _ := range n.nat {
		n.nat[i] = empty
	}
}
