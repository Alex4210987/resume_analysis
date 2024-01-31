package service

import (
	"github.com/gin-gonic/gin"
	"resume_analysis/model"
)

func CreatePDF(ctx *gin.Context, pdf *model.PDF) error {

}
func IsPDFExisted(ctx *gin.Context, pdfTested *model.PDF) (*model.PDF, bool, error) {

}

func UpdatePDF(ctx *gin.Context, pdf *model.PDF) error {

}

func GetMyPDF(ctx *gin.Context, limit int, offset int) ([]map[string]interface{}, error) {

}
