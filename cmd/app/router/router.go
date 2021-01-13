package router

import (
	"github.com/ahmedashrafdev/nozom_ecommerce/app/app"

	"github.com/go-chi/chi"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("GET", "/", app.HandleIndex)
	return r
}
