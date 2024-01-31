package api

import (
	"net/http"
	"resume_analysis/model"
	"resume_analysis/service"
	"time"

	"github.com/gin-gonic/gin"
)

type createPDFParam struct {
	PDFUrl    string `form:"pdf_url" binding:"required"`
	PDFName   string `form:"pdf_name" binding:"required"`
	PDFSize   int64  `form:"pdf_size" binding:"required"`
	Content   string `form:"content" binding:"required"`
	ValidDays int    `form:"valid_days" binding:"required"`
}

func CreatePDF(ctx *gin.Context) {
	p := &createPDFParam{}
	if err := ctx.ShouldBind(p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	pdfCreated := &model.PDF{
		PDFUrl:   p.PDFUrl,
		PDFName:  p.PDFName,
		PDFSize:  p.PDFSize,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		ExpireAt: time.Now().AddDate(0, 0, p.ValidDays),
	}
	pdfGet, existed, err := service.IsPDFExisted(ctx, pdfCreated)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	} else if existed {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg":    "PDF已存在",
			"pdf_id": pdfGet.PDFId,
		})
		return
	}
	if err := service.CreatePDF(ctx, pdfCreated); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":    "发布成功",
		"pdf_id": pdfCreated.PDFId,
	})
}

type updatePDFParam struct {
	PDFId     int64  `form:"pdf_id" binding:"required"`
	PDFUrl    string `form:"pdf_url" binding:"required"`
	PDFName   string `form:"pdf_name" binding:"required"`
	PDFSize   int64  `form:"pdf_size" binding:"required"`
	Content   string `form:"content" binding:"required"`
	ValidDays int    `form:"valid_days" binding:"required"`
}

func UpdatePDF(ctx *gin.Context) {
	p := &updatePDFParam{}
	if err := ctx.ShouldBind(p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	pdfUpdated := &model.PDF{
		PDFId:    p.PDFId,
		PDFUrl:   p.PDFUrl,
		PDFName:  p.PDFName,
		PDFSize:  p.PDFSize,
		Content:  p.Content,
		UpdateAt: time.Now(),
	}
	if err := service.UpdatePDF(ctx, pdfUpdated); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}

type getMyPDFParams struct {
	Limit  int `form:"limit" binding:"required"`
	Offset int `form:"offset" binding:"required"`
}

func GetMyPDF(ctx *gin.Context) {
	p := &getMyPDFParams{}
	if err := ctx.ShouldBind(p); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	pdfs, err := service.GetMyPDF(ctx, p.Limit, p.Offset)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg":  "获取成功",
		"pdfs": pdfs,
	})
}
