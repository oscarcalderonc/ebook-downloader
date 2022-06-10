package config

import "github.com/gin-gonic/gin"

func Start() {
	r := gin.Default()

	r.Run()
}
