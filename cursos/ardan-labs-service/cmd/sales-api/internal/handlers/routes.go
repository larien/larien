package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/larien/service/cmd/internal/mid"
	"github.com/larien/service/cmd/internal/platform/web"
)

// API constructs an http.Handler with all application routes defined.
func API(build string, shutdown chan os.Signal, log *log.Logger) *web.App { // deixar o desacoplamento para o caller

	app := web.NewApp(shutdown, mid.Logger(log))

	app.Handle(http.MethodGet, "/test", health)

	return app
}
