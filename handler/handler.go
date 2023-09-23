package handler

import (
	"context"
	"encoding/json"
	"hexagonal/usecase"
	"net/http"
)

type Handlers struct {
	uc usecase.UsecaseInterface
}

func NewHandlers(uc *usecase.UseCases) *Handlers {
	return &Handlers{
		uc: uc,
	}
}

func (h *Handlers) ATCHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := h.uc.DoAtc(context.Background(), usecase.RequestFrontEndATC{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) CartlistHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := h.uc.DoCartList(context.Background(), usecase.RequestFrontEndCartlist{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
