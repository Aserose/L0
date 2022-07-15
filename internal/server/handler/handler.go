package handler

import (
	"L0/internal/server/handler/htmlPage"
	"L0/internal/service"
	"L0/pkg/customLogger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	purInfo purchaseInfo
	errInfo errorInfo
	html    htmlPage.Html
}

func NewHandler(log customLogger.Logger, s service.Service) Handler {
	return Handler{
		purInfo: newPurchaseInfo(log, s),
		errInfo: newErrInfo(log, s),
		html:    htmlPage.NewHtml(),
	}
}

func (h Handler) SetupRoutes() *http.ServeMux {
	router := http.DefaultServeMux

	router.HandleFunc("/", h.main)
	router.HandleFunc("/purchase", h.purInfo.setup)
	router.HandleFunc("/error", h.errInfo.setup)

	return router
}

func (h Handler) main(w http.ResponseWriter, r *http.Request) {
	h.html.TemplateExecute(w, "main.html")

	if val := r.FormValue("channel"); val != `` {
		h.outputLastError(w, r)
	} else if r.FormValue("id") != `` {
		h.outputPurchaseInfo(w, r)
	}
}

func (h Handler) outputPurchaseInfo(w http.ResponseWriter, r *http.Request) {
	h.purInfo.getCountFormat(w, r)
	h.purInfo.getInfo(w, r)
}

func (h Handler) outputLastError(w http.ResponseWriter, r *http.Request) {
	h.updateRequestBody(r, r.FormValue("channel"))
	h.purInfo.putInfoCh(w, r)
	time.Sleep(250 * time.Millisecond)
	h.errInfo.getLastFormat(w, r)
}

func (h Handler) updateRequestBody(r *http.Request, value string) {
	newBody := value
	r.Body = ioutil.NopCloser(strings.NewReader(newBody))
	r.ContentLength = int64(len(newBody))
}
