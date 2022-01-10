package main

import (
	_ "miniprogram-backend/repository"
	"miniprogram-backend/router"
)

func main() {
	r := router.NewRouter()
	// r.Run("0.0.0.0:80")
	r.RunTLS("0.0.0.0:8080", "/eventbackend/test.pem", "/eventbackend/test.key")
}

