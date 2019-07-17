package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(host string) {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("web/app/public/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H {
			"title": "dup - A duplication checker tool on GitHub",
		})
	})

	r.Run(host)
}
