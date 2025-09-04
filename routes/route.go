package routes

import (
	"net/http"
	"restapi-golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, mhs *handlers.MhsHandler, dashboardMhsHandler *handlers.DashboardMhsHandler) {
	router := r.Group("/api/v1")
	{
		router.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}
	setUpMhsRoutes(router, mhs)
	setUpDashboardMhsRoutes(router, dashboardMhsHandler)
}

func setUpMhsRoutes(rg *gin.RouterGroup, mhs *handlers.MhsHandler) {
	rg.GET("/mhs/:nim", mhs.GetDetailMhs)
	rg.GET("/mhs/histories/:status", mhs.GetMahasiswaHistoryByStatus)
	rg.GET("/mhs/history", mhs.GetMahasiswaHistoryFiltered)
}
func setUpDashboardMhsRoutes(rg *gin.RouterGroup, dashboardMhsHandler *handlers.DashboardMhsHandler) {
	rg.GET("/dashboard/overview", dashboardMhsHandler.GetDashboardOverview)
	rg.GET("/dashboard/fakultas", dashboardMhsHandler.GetDrilldownFakultas)
	rg.GET("/dashboard/jurusan", dashboardMhsHandler.GetDrilldownJurusan)
	rg.GET("/dashboard/prodi", dashboardMhsHandler.GetDrilldownProdi)
}
