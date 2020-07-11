package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"net/http"

	"github.com/larien/service/cmd/internal/platform/web"
)

func health(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	i := rand.Intn(100)
	if i%2 == 0 {
		return web.NewRequestError(errors.New("trusted error"), http.StatusBadRequest)
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return json.NewEncoder(w).Encode(status)
}
