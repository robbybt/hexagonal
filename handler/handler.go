package handler

import (
	"context"
	"encoding/json"
	"hexagonal/usecase"
	"net/http"
)

type Handlers struct {
	UC usecase.UsecaseInterface
}

func NewHandlers(uc *usecase.UseCases) *Handlers {
	return &Handlers{UC: uc}
}

func (h *Handlers) ATCHandler(w http.ResponseWriter, r *http.Request) {
	resp, _ := h.UC.DoAtc(context.Background(), usecase.RequestFrontEndATC{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
