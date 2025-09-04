package handlers

import (
	"context"
	"math"
	"net/http"
	"restapi-golang/models"
	"restapi-golang/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MhsHandler struct {
	MhsService *services.MhsService
}

func NewMhsHandler(service *services.MhsService) *MhsHandler {
	return &MhsHandler{
		MhsService: service,
	}
}

func (h *MhsHandler) GetDetailMhs(c *gin.Context) {
	nim := c.Param("nim")
	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	mh, err := h.MhsService.GetDetailMhs(ctx, nim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": mh, "message": "OK"})
}

func (h *MhsHandler) GetMahasiswaHistoryByStatus(c *gin.Context) {
	status := c.Param("status")
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	tahunStr := c.Query("tahun")
	semester := c.Query("semester")

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun harus berupa angka"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}
	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	results, total, err := h.MhsService.GetMahasiswaHistoryByStatus(ctx, status, page, limit, tahun, semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   results,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *MhsHandler) GetMahasiswaHistoryFiltered(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")
	tahunStr := c.Query("tahun")
	semester := c.Query("semester")

	// Convert to integers dengan error handling
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tahun must be a number"})
		return
	}

	filter := models.MahasiswaHistoryRequest{
		Nama:            c.Query("Nama"),
		Tahun:           tahun,
		Semester:        semester,
		Page:            page,
		Limit:           limit,
		Fakultas:        c.Query("fakultas"),
		Jurusan:         c.Query("jurusan"),
		Prodi:           c.Query("prodi"),
		Status:          c.Query("status"),
		Kewarganegaraan: c.Query("kewarganegaraan"),
		NIM:             c.Query("nim"),
	}

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	results, total, err := h.MhsService.GetMahasiswaHistoryFiltered(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   results,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
