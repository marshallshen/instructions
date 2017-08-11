package main

import (
  "github.com/gin-gonic/gin"
  "./app"
)

func SetupRouter() *gin.Engine {
  router := gin.Default()

  v1 := router.Group("api/v1") 
  {
    v1.GET("/instructions", app.GetInstructions)
    v1.GET("/instructions/:id", app.GetInstruction)
    v1.POST("/instructions", app.PostInstruction)
    v1.PUT("/instructions/:id", app.UpdateInstruction)
    v1.DELETE("/instructions/:id", app.DeleteInstruction)
  }

  return router
}

func main() {
  router := SetupRouter()
  router.Run(":8080")
}
