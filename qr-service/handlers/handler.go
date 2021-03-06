package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fuato1/shorturl/qr-service/services"
	"github.com/fuato1/shorturl/qr-service/services/commands"
)

type Handler struct {
	qrServices services.QrServices
}

func NewHandler(qrServices services.QrServices) *Handler {
	return &Handler{qrServices: qrServices}
}

func (h *Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	qrs, err := h.qrServices.Queries.GetAllQrsHandler.Handle()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	err = json.NewEncoder(w).Encode(qrs)
	if err != nil {
		return
	}
}

type AddQrRequestModel struct {
	Source string `json:"source"`
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	var qr AddQrRequestModel
	err := json.NewDecoder(r.Body).Decode(&qr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}

	err = h.qrServices.Commands.AddQrHandler.Handle(commands.AddQrRequest{
		Source: qr.Source,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}
