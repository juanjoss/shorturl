package api

import (
	"bytes"
	"image"
	"image/png"
	"net/http"

	"github.com/gorilla/mux"
	grpc "github.com/juanjoss/shorturl/pkg/grpc/client"
)

func (s *server) getQR(w http.ResponseWriter, r *http.Request) {
	_, span := s.tracer.Start(r.Context(), "getQR")
	defer span.End()

	id := mux.Vars(r)["id"]

	qr, err := grpc.GenerateQR("http://localhost/" + id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	image, _, err := image.Decode(bytes.NewReader(qr))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	png.Encode(w, image)
}
