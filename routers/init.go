package routers

import (
	_Controllers "go-back/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/go-back")
	{
		grp1.POST("callupload", _Controllers.CallAPIUploadFile)
	}

	return r
}
