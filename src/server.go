package src

import (
	"fmt"
	"net/http"

	"github.com/cjmarkham/GoBB/config"
)

func ProvideServer(router http.Handler, cfg *config.ServerConfig) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
}
