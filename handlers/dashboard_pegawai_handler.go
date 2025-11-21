package handlers

import (
	"context"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardPegawaiHandler struct {
	DashboardPegawaiService *services.DashboardPegawaiService
}

func NewDashboardPegawaiHandler(service *services.DashboardPegawaiService) *DashboardPegawaiHandler {
	return &DashboardPegawaiHandler{
		DashboardPegawaiService: service,
	}
}

func (h *DashboardPegawaiHandler) GetDashboardPegawaiOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")

	var tahun int

	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	result, err := h.DashboardPegawaiService.GetDashboardPegawaiOverview(ctx, tahun)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "datas": result})
}

func (h *DashboardPegawaiHandler) GetDrilldownPegawaiFakultas(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiFakultas(ctx, tahun, statusPegawai, statusKeaktifan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"datas":   items,
		"total":   total,
	})
}

func (h *DashboardPegawaiHandler) GetDrilldownPegawaiJurusan(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")
	kodeFakultas := c.Query("kodeFakultas")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiJurusan(ctx, tahun, statusPegawai, statusKeaktifan, kodeFakultas)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"datas":   items,
		"total":   total,
	})
}

func (h *DashboardPegawaiHandler) GetDrilldownPegawaiProdi(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")
	kodeJurusan := c.Query("kodeJurusan")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiProdi(ctx, tahun, statusPegawai, statusKeaktifan, kodeJurusan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"datas":   items,
		"total":   total,
	})
}
