package handlers

import (
	"fmt"
	"net/http"
	"restapi-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DashboardMhsHandler struct {
	DashboardMhsService *services.DashboardMhsService
}

func NewDashboardMhsHandler(service *services.DashboardMhsService) *DashboardMhsHandler {
	return &DashboardMhsHandler{
		DashboardMhsService: service,
	}
}

// ==========================
// GET /dashboard/overview?tahun=2023&semester=1
// ==========================
func (h *DashboardMhsHandler) GetDashboardOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")
	semester := c.Query("semester")

	tahun, err := strconv.Atoi(tahunStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "tahun harus berupa angka"})
		return
	}

	result, err := h.DashboardMhsService.GetDashboardOverview(c.Request.Context(), tahun, semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// ==========================
// GET /dashboard/fakultas?tahun=2023&semester=1&status=A
// ==========================
func (h *DashboardMhsHandler) GetDrilldownFakultas(c *gin.Context) {
	tahun, _ := strconv.Atoi(c.Query("tahun"))
  semester := c.Query("semester")
	status := c.Query("status")
	fmt.Println("Statusnya: " + status)

	items, total, err := h.DashboardMhsService.GetDrilldownFakultas(c.Request.Context(), tahun, semester, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
	})
}

// ==========================
// GET /dashboard/jurusan?tahun=2023&semester=1&status=A&fakultasKode=FT01
// ==========================
func (h *DashboardMhsHandler) GetDrilldownJurusan(c *gin.Context) {
	tahun, _ := strconv.Atoi(c.Query("tahun"))
  semester := c.Query("semester")
	status := c.Query("status")
	fakultasKode := c.Query("fakultasKode")

	items, total, err := h.DashboardMhsService.GetDrilldownJurusan(c.Request.Context(), tahun, semester, status, fakultasKode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
	})
}

// ==========================
// GET /dashboard/prodi?tahun=2023&semester=1&status=A&jurusanKode=FT01_TI
// ==========================
func (h *DashboardMhsHandler) GetDrilldownProdi(c *gin.Context) {
	tahun, _ := strconv.Atoi(c.Query("tahun"))
  semester := c.Query("semester")
	status := c.Query("status")
	jurusanKode := c.Query("jurusanKode")

	items, total, err := h.DashboardMhsService.GetDrilldownProdi(c.Request.Context(), tahun, semester, status, jurusanKode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
		"total": total,
	})
}
