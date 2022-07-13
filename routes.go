package main

import (
	"x/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/supertokens/supertokens-golang/recipe/session"
)

func routes(r *chi.Mux) {

	r.Get("/login", handlers.Login)
	r.Get("/auth/callback/google", handlers.Callback)
	r.Get("/printuser", session.VerifySession(nil, handlers.PrintUserDetails))
}
