package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	routerFns = append(routerFns, func(group *gin.RouterGroup) {
		useSysRouter(group)
	})
}

func useSysRouter(group *gin.RouterGroup) {

	group.GET("/list", List)

}
func List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"2": 1})

}
