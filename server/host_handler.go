package server

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/M4xwell/network-service/host"
	"gitlab.com/M4xwell/network-service/pkg/logger"
	"html/template"
	"net/http"
)

type HostHandler struct {
	Service host.Service

	Logger logger.Logger
}

func (h *HostHandler) ListHosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hosts, err := h.Service.ListHosts()
	if err != nil {
		h.Logger.Errorf("could not get hosts: %v", err)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		h.Logger.Errorf("could not parse html: %v", err)
		return
	}

	h.Logger.Infof("found: %v hosts", len(hosts))

	err = tmpl.Execute(w, hosts)
	if err != nil {
		h.Logger.Errorf("could not execute template: %v", err)
		return
	}
}

func (h *HostHandler) DetailHost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")

	hst, err := h.Service.FindHostByName(name)
	if err != nil {
		h.Logger.Errorf("could not find host: %v", err)
		return
	}

	tmpl, err := template.ParseFiles("static/detail.html")
	if err != nil {
		h.Logger.Errorf("could not parse html: %v", err)
		return
	}

	err = tmpl.Execute(w, hst)
	if err != nil {
		h.Logger.Errorf("could execute template: %v", err)
		return
	}
}
