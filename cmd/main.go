package main

import (
	"crypto"
	"fmt"
	jwtMiddleware "github.com/codeby-student/go-service/pkg/middleware/jwt"
	"github.com/codeby-student/go-service/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	publicKey, err := os.ReadFile("production/public.key")
	if err != nil {
		log.Fatal(fmt.Errorf("error read public key: %w", err))
	}
	//privateKey, err := os.ReadFile("private.key")
	//if err != nil {
	//	log.Fatal(fmt.Errorf("error read private key: %w", err))
	//}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	jwt := jwtMiddleware.Jwt(crypto.SHA512, publicKey)
	r.Use(jwt)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, ok := jwtMiddleware.FromContext(r.Context())
		if !ok {
			utils.WriteResponse(w, http.StatusUnauthorized, "err.auth_required")
			return
		}

		utils.WriteResponse(w, http.StatusOK, fmt.Sprintf("Welcome, %s!", payload.Subject))
	})
	err = http.ListenAndServe(":8888", r)
	if err != nil {
		log.Fatal(err)
	}
}
