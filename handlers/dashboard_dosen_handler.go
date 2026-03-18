package handlers

import (
	"context"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type DashboardDosenHandler struct {
	DashboardDosenService *services.DashboardDosenService
}

func NewDashboardDosenHandler(service *services.DashboardDosenService) *DashboardDosenHandler {
	return &DashboardDosenHandler{
		DashboardDosenService: service,
	}
}

// GetDashboardDosenOverview mendapatkan  overview dashboard dosen
// @Summary      Get Dashboard Dosen Overview
// @Description  Mendapatkan overview dashboard dosen
// @Tags         Dashboard Dosen
// @Accept       json
// @Produce      json
// @Param        tahun        query     int  false  "tahun"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DashboardCardPegawai}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-dosen/overview [get]
func (h *DashboardDosenHandler) GetDashboardDosenOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun int

	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	result, err := h.DashboardDosenService.GetDashboardDosenOverview(ctx, tahun)
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

// GetDrilldownDosenFakultas mendapatkan drilldown fakultas dashboard Dosen
// @Summary      Get drilldown fakultas dosen Overview
// @Description  Mendapatkan drilldown fakultas dashboard Dosen
// @Tags         Dashboard Dosen
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        tahun        query     int  false  "tahun"
// @Param        statusPegawai        query     int  false  "statusPegawai"
// @Param        statusKeaktifan        query     int  false  "statusKeaktifan"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-dosen/fakultas [get]
func (h *DashboardDosenHandler) GetDrilldownDosenFakultas(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardDosenService.GetDrilldownDosenFakultas(ctx, tahun, statusPegawai, statusKeaktifan)
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

// GetDrilldownDosenJurusan mendapatkan drilldown jurusan dashboard Dosen
// @Summary      Get drilldown jurusan dosen Overview
// @Description  Mendapatkan drilldown jurusan dashboard Dosen
// @Tags         Dashboard Dosen
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        tahun        query     int  false  "tahun"
// @Param        statusPegawai        query     int  false  "statusPegawai"
// @Param        statusKeaktifan        query     int  false  "statusKeaktifan"
// @Param        kodeFakultas        query     string  false  "kodeFakultas"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-dosen/jurusan [get]
func (h *DashboardDosenHandler) GetDrilldownDosenJurusan(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")
	kodeFakultas := c.Query("kodeFakultas")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardDosenService.GetDrilldownDosenJurusan(ctx, tahun, statusPegawai, statusKeaktifan, kodeFakultas)
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

// GetDrilldownDosenProdi mendapatkan drilldown prodi dashboard Dosen
// @Summary      Get drilldown prodi dosen Overview
// @Description  Mendapatkan drilldown prodi dashboard Dosen
// @Tags         Dashboard Dosen
// @Accept       json
// @Produce      json
// @Param        status        query     string  false  "status"
// @Param        tahun        query     int  false  "tahun"
// @Param        statusPegawai        query     int  false  "statusPegawai"
// @Param        statusKeaktifan        query     int  false  "statusKeaktifan"
// @Param        kodeJurusan        query     string  false  "kodeJurusan"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DrilldownItem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-dosen/prodi [get]
func (h *DashboardDosenHandler) GetDrilldownDosenProdi(c *gin.Context) {
	tahunStr := c.Query("tahun")
	statusPegawaiStr := c.Query("statusPegawai")
	statusKeaktifanStr := c.Query("statusKeaktifan")
	kodeJurusan := c.Query("kodeJurusan")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun, statusPegawai, statusKeaktifan int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}

	tahun = utils.StringToInt(tahunStr, 0)
	statusPegawai = utils.StringToInt(statusPegawaiStr, 0)
	statusKeaktifan = utils.StringToInt(statusKeaktifanStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	items, total, err := h.DashboardDosenService.GetDrilldownDosenProdi(ctx, tahun, statusPegawai, statusKeaktifan, kodeJurusan)
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
