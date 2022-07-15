package handler

import (
	"L0/internal/datastruct"
	"L0/internal/service"
	"L0/pkg/customLogger"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	formChannel = `ch`
	formCache   = `cache`
	formId      = `id`
	enable      = `1`
)

type purchaseInfo struct {
	service service.Service
	log     customLogger.Logger
}

func newPurchaseInfo(log customLogger.Logger, s service.Service) purchaseInfo {
	return purchaseInfo{
		service: s,
		log:     log,
	}
}

func (p purchaseInfo) setup(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getInfo(w, r)
	case http.MethodPut:
		p.putInfo(w, r)
	}
}

func (p purchaseInfo) putInfo(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue(formChannel) {
	case enable:
		p.putInfoCh(w, r)
	default:
		p.putInfoDB(w, r)
	}
}

func (p purchaseInfo) putInfoDB(w http.ResponseWriter, r *http.Request) {
	order := datastruct.Order{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.Unmarshal(b, &order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !order.IsValid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}
	_, err = w.Write([]byte(strconv.Itoa(p.service.Put(order))))
	if err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func (p purchaseInfo) getCountFormat(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(`count: ` + strconv.Itoa(p.service.GetCount())))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (p purchaseInfo) getCount(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(strconv.Itoa(p.service.GetCount())))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (p purchaseInfo) putInfoCh(w http.ResponseWriter, r *http.Request) {
	sendFunc := p.service.Streamer.ChanPurchase()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := sendFunc(b); err != nil {
		p.log.Error(p.log.CallInfoStr(), err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (p purchaseInfo) getInfo(w http.ResponseWriter, r *http.Request) {
	val, err := strconv.Atoi(r.FormValue(formId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if val <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch r.FormValue(formCache) {
	case enable:
		_, err = w.Write(p.service.GetCached(val))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		_, err = w.Write(p.service.Get(val))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
