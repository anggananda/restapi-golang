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

type RealisasiUnitHandler struct {
	RealisasiUnitService *services.RealisasiUnitService
}

func NewRealisasiUnitHandler(service *services.RealisasiUnitService) *RealisasiUnitHandler {
	return &RealisasiUnitHandler{
		RealisasiUnitService: service,
	}
}

// GetRealisasiUnitFiltered mendapatkan data ralisasi unit dengan filter dan pagination
// @Summary      Get filtered ralisasi unit
// @Description  Mendapatkan data ralisasi unit berdasarkan filter dengan pagination
// @Tags         Realisasi Unit
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.RealisasiUnit}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /realisasi-unit [get]
func (h *RealisasiUnitHandler) GetRealisasiUnitFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	search := c.Query("search")
	tahun := c.Query("tahun")
	contentType := c.DefaultQuery("contentType", "json")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	realisasiUnit, total, err := h.RealisasiUnitService.GetRealisasiUnitFiltered(ctx, search, tahun, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{
			"status": "success",
			"datas":  realisasiUnit,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
			"datas":  realisasiUnit,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	}
}

// ExportRealisasiUnitCSV mengekspor data realisasi unit ke format CSV
// @Summary      Export Realisasi Unit ke CSV
// @Description  Mengekspor daftar realisasi unit yang telah difilter ke dalam file CSV.
// @Tags         Realisasi Unit
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /realisasi-unit/export-csv [get]
func (h *RealisasiUnitHandler) ExportRealisasiUnitCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	search := c.Query("search")
	tahun := c.Query("tahun")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	realisasiUnit, _, err := h.RealisasiUnitService.GetRealisasiUnitFiltered(ctx, search, tahun, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Kode ID", "Tahun Anggaran", "Kode Unit", "Nama Unit",
		"Pagu PNBP", "Pagu RM", "Pagu RM BOPTN",
		"Realisasi PNBP", "Realisasi RM", "Realisasi RM BOPTN",
	}

	var csvData [][]string

	for _, item := range realisasiUnit {

		idStr := strconv.Itoa(item.ID)

		row := []string{

			idStr,
			item.Kode,
			item.TahunAnggaran,
			item.KodeUnit,
			item.NamaUnit,

			item.PaguPNBP,
			item.PaguRM,
			item.PaguRMBOPTN,

			item.RealisasiPNBP,
			item.RealisasiRM,
			item.RealisasiRMBOPTN,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("realisasi_unit_%s_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
