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

// GetDashboardPegawaiOverview mendapatkan  overview dashboard pegawai
// @Summary      Get Dashboard pegawai Overview
// @Description  Mendapatkan overview dashboard pegawai
// @Tags         Dashboard Pegawai
// @Accept       json
// @Produce      json
// @Param        tahun        query     int  false  "tahun"
// @Success      200           {object}  models.ListDetailResponse{datas=models.DashboardCardPegawai}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dashboard-pegawai/overview [get]
func (h *DashboardPegawaiHandler) GetDashboardPegawaiOverview(c *gin.Context) {
	tahunStr := c.Query("tahun")
	contentType := c.DefaultQuery("contentType", "json")

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

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{"message": "OK", "datas": result})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "datas": result})
	}
}

// GetDrilldownPegawaiFakultas mendapatkan drilldown fakultas dashboard pegawai
// @Summary      Get drilldown fakultas pegawai Overview
// @Description  Mendapatkan drilldown fakultas dashboard pegawai
// @Tags         Dashboard Pegawai
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
// @Router       /dashboard-pegawai/fakultas [get]
func (h *DashboardPegawaiHandler) GetDrilldownPegawaiFakultas(c *gin.Context) {
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

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiFakultas(ctx, tahun, statusPegawai, statusKeaktifan)
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

// GetDrilldownPegawaiJurusan mendapatkan drilldown jurusan dashboard pegawai
// @Summary      Get drilldown jurusan pegawai Overview
// @Description  Mendapatkan drilldown jurusan dashboard pegawai
// @Tags         Dashboard Pegawai
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
// @Router       /dashboard-pegawai/jurusan [get]
func (h *DashboardPegawaiHandler) GetDrilldownPegawaiJurusan(c *gin.Context) {
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

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiJurusan(ctx, tahun, statusPegawai, statusKeaktifan, kodeFakultas)
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

// GetDrilldownPegawaiProdi mendapatkan drilldown prodi dashboard pegawai
// @Summary      Get drilldown prodi pegawai Overview
// @Description  Mendapatkan drilldown prodi dashboard pegawai
// @Tags         Dashboard Pegawai
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
// @Router       /dashboard-pegawai/prodi [get]
func (h *DashboardPegawaiHandler) GetDrilldownPegawaiProdi(c *gin.Context) {
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

	items, total, err := h.DashboardPegawaiService.GetDrilldownPegawaiProdi(ctx, tahun, statusPegawai, statusKeaktifan, kodeJurusan)
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
