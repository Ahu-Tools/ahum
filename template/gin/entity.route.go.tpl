package {{.EntityName}}

import (
	"github.com/gin-gonic/gin"
	// @ahum: imports
)

func RegisterRoutes(r *gin.RouterGroup) {
	h := NewHandler()
	r.GET("/health", h.Check)
	// @ahum: routes
}
