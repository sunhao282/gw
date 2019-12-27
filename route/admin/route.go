package admin

import (
	"gw/pkg/admin"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.POST("/req/add/api", admin.Add)
	r.POST("/req/add/group", admin.AddGroup)
}
