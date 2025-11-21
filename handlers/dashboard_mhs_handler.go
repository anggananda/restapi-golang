package handlers

import (
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

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

func (h *DashboardMhsHandler) GetDashboardMhsOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")

	var tahun, semester int

	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)
	semester = utils.StringToInt(semesterStr, 0)

	result, err := h.DashboardMhsService.GetDashboardMhsOverview(c.Request.Context(), tahun, semester)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OK", "datas": result})
}

func (h *DashboardMhsHandler) GetDrilldownMhsFakultas(c *gin.Context) {
	status := c.Query("status")
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")

	var tahun, semester int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)
	semester = utils.StringToInt(semesterStr, 0)

	items, total, err := h.DashboardMhsService.GetDrilldownMhsFakultas(c.Request.Context(), tahun, semester, status)
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

func (h *DashboardMhsHandler) GetDrilldownMhsJurusan(c *gin.Context) {
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")
	status := c.Query("status")
	kodeFakultas := c.Query("kodeFakultas")

	var tahun, semester int

	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)
	semester = utils.StringToInt(semesterStr, 0)

	items, total, err := h.DashboardMhsService.GetDrilldownMhsJurusan(c.Request.Context(), tahun, semester, status, kodeFakultas)
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

func (h *DashboardMhsHandler) GetDrilldownMhsProdi(c *gin.Context) {
	status := c.Query("status")
	kodeJurusan := c.Query("kodeJurusan")
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")

	var tahun, semester int

	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)
	semester = utils.StringToInt(semesterStr, 0)

	items, total, err := h.DashboardMhsService.GetDrilldownMhsProdi(c.Request.Context(), tahun, semester, status, kodeJurusan)
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
