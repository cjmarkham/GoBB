package GoBB

import (
	"fmt"
	"net/http"

	config2 "github.com/cjmarkham/GoBB/config"
)

func ProvideServer(router http.Handler, config *config2.ServerConfig) *http.Server {
	return &http.Server{
		Addr: fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
}
