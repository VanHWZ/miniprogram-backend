package main

import (
	_ "miniprogram-backend/repository"
	"miniprogram-backend/router"
)

func main() {
	r := router.NewRouter()
	r.Run("127.0.0.1:8888")
}

