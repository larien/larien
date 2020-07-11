package mid

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/larien/service/cmd/internal/platform/web"
)

// Logger ...
func Logger(log *log.Logger) web.Middleware {
	m := func(before web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
			if !ok {
				return nil //return web.NewShutdownError("web value missing from context")
			}

			before(ctx, w, r, params)

			log.Printf("%s : (%d) : %s %s -> %s (%s)",
				v.TraceID, v.StatusCode,
				r.Method, r.URL.Path,
				r.RemoteAddr, time.Since(v.Now),
			)

			return nil
		}
		return h
	}
	return m
}
