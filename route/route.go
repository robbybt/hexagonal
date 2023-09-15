package route

import (
	"hexagonal/handler"
	"net/http"
)

func InitializeRoutes(h *handler.Handlers) {
	http.HandleFunc("/fetchData", h.ATCHandler)
}
