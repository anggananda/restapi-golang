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

// GetDashboardMhsOverview mendapatkan  overview dashboard mahasiswa
// @Summary      Get Dashboard Mhs Overview
// @Description  Mendapatkan overview dashboard mahasiswa
// @Tags         Dashboard Mahasiswa
// @Accept       json
// @Produce      json
// @Param        tahun        query     int  false  "tahun"
// @Param        semester        query     int  false  "semester"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DashboardCard}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-mhs/overview [get]
func (h *DashboardMhsHandler) GetDashboardMhsOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")
	contentType := c.DefaultQuery("contentType", "json")

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

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{"message": "OK", "datas": result})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "datas": result})
	}
}

// GetDrilldownMhsFakultas mendapatkan drilldown fakultas dashboard mahasiswa
// @Summary      Get drilldown fakultas Mhs Overview
// @Description  Mendapatkan drilldown fakultas dashboard mahasiswa
// @Tags         Dashboard Mahasiswa
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        tahun        query     int  false  "tahun"
// @Param        semester        query     int  false  "semester"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-mhs/fakultas [get]
func (h *DashboardMhsHandler) GetDrilldownMhsFakultas(c *gin.Context) {
	status := c.Query("status")
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")
	contentType := c.DefaultQuery("contentType", "json")

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

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	}
}

// GetDrilldownMhsJurusan mendapatkan drilldown jurusan dashboard mahasiswa
// @Summary      Get drilldown jurusan Mhs Overview
// @Description  Mendapatkan drilldown jurusan dashboard mahasiswa
// @Tags         Dashboard Mahasiswa
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        kodeFakultas        query     string  false  "kodeFakultas"
// @Param        tahun        query     int  false  "tahun"
// @Param        semester        query     int  false  "semester"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-mhs/jurusan [get]
func (h *DashboardMhsHandler) GetDrilldownMhsJurusan(c *gin.Context) {
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")
	status := c.Query("status")
	kodeFakultas := c.Query("kodeFakultas")
	contentType := c.DefaultQuery("contentType", "json")

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

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	}
}

// GetDrilldownMhsProdi mendapatkan drilldown prodi dashboard mahasiswa
// @Summary      Get drilldown prodi Mhs Overview
// @Description  Mendapatkan drilldown prodi dashboard mahasiswa
// @Tags         Dashboard Mahasiswa
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        kodeJurusan        query     string  false  "kodeJurusan"
// @Param        tahun        query     int  false  "tahun"
// @Param        semester        query     int  false  "semester"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-mhs/prodi [get]
func (h *DashboardMhsHandler) GetDrilldownMhsProdi(c *gin.Context) {
	status := c.Query("status")
	kodeJurusan := c.Query("kodeJurusan")
	tahunStr := c.Query("tahun")
	semesterStr := c.Query("semester")
	contentType := c.DefaultQuery("contentType", "json")

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

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"datas":   items,
			"total":   total,
		})
	}
}
