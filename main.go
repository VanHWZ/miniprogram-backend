package main

import (
	"miniprogram-backend/repository"
	"miniprogram-backend/router"
)

func main() {
	repository.DatabaseInit()
	r := router.NewRouter()
	r.Run("127.0.0.1:8888")
}

