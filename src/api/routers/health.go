package routers

import (
	"github.com/amirhasanpour/bama-clone/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
