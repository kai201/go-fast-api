package routers

import (
	"github.com/gin-gonic/gin"
)

// NewRouter create a new router
func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	return r
}
