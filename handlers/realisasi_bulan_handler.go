package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type RealisasiBulanHandler struct {
	RealisasiBulanService *services.RealisasiBulanService
}

func NewRealisasiBulanHandler(service *services.RealisasiBulanService) *RealisasiBulanHandler {
	return &RealisasiBulanHandler{
		RealisasiBulanService: service,
	}
}

// GetRealisasiBulanFiltered mendapatkan data realisasi bulan dengan filter dan pagination
// @Summary      Get filtered realisasi bulan
// @Description  Mendapatkan data realisasi bulan berdasarkan filter dengan pagination
// @Tags         Realisasi Bulan
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.RealisasiBulan}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /realisasi-bulan [get]
func (h *RealisasiBulanHandler) GetRealisasiBulanFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahun := c.Query("tahun")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	realisasiBulan, total, err := h.RealisasiBulanService.GetRealisasiBulanFiltered(ctx, tahun, search, page, limit)
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
		"datas":  realisasiBulan,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

// ExportRealisasiBulanCSV mengekspor data realisasi bulan ke format CSV
// @Summary      Export Realisasi Bulan ke CSV
// @Description  Mengekspor daftar realisasi bulan yang telah difilter ke dalam file CSV.
// @Tags         Realisasi Bulan
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /realisasi-bulan/export-csv [get]
func (h *RealisasiBulanHandler) ExportRealisasiBulanCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	tahun := c.Query("tahun")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	realisasiBulan, _, err := h.RealisasiBulanService.GetRealisasiBulanFiltered(ctx, tahun, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Kode ID", "Tahun Anggaran", "Bulan",
		"Realisasi Total PNBP", "Realisasi Total RM", "Realisasi Total RM BOPTN",
	}

	var csvData [][]string

	for _, item := range realisasiBulan {

		idStr := strconv.Itoa(item.ID)

		row := []string{

			idStr,
			item.Kode,
			item.TahunAnggaran,
			item.Bulan,

			item.RealisasiTotalPNBP,
			item.RealisasiTotalRM,
			item.RealisasiTotalRMBOPTN,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("realisasi_bulan_%s_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
