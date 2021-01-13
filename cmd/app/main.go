package main

import (
	"fmt"
	"net/http"

	"github.com/ahmedashrafdev/nozom_ecommerce/app/router"
	"github.com/ahmedashrafdev/nozom_ecommerce/config"
	lr "github.com/ahmedashrafdev/nozom_ecommerce/utl/config"
)

func main() {
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
	appRouter := router.New()
	address := fmt.Sprintf(":%d", appConf.Server.Port)
	logger.Info().Msgf("Starting server %v", address)
	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal().Err(err).Msg("Server startup failed")
	}
}
func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
