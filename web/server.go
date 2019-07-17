package web

import "github.com/gin-gonic/gin"

func Run(host string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())



	r.Run(host)
}
