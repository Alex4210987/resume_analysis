package api

import (
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func useRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		pdf := api.Group("/pdf")
		{
			pdf.POST("/", CreatePDF)
			pdf.PUT("/", UpdatePDF)
			pdf.GET("/", GetMyPDF)
		}
	}
}

func init() {
	Engine = gin.Default()
	useRouter(Engine)
}
