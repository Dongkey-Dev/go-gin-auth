package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  r := gin.Default()
  public := r.Group("/api")
  public.POST("/user", func(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{"data": "user post endpoint."})
  })

  public.GET("/ping", func(c *gin.Context) {
    c.String(200, "pong")
  })
  return r
}

func main() {
  r := setupRouter()
  r.Run(":8080")
}
