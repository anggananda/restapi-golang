package routes

import (
	"restapi-golang/handlers"
	"restapi-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, casHandler *handlers.CASHandler, userHandler *handlers.UserHandler, mhs *handlers.MhsHandler, dashboardMhsHandler *handlers.DashboardMhsHandler, perpemHandler *handlers.PerpemHandler, angketMhsHandler *handlers.AngketMhsHandler, kritikSaranHandler *handlers.KritikSaranHandler, agendaMengajarHandler *handlers.AgendaMengajarHandler, mhsWisudaHandler *handlers.MhsWisudaHandler, rekapPMBHandler *handlers.RekapPMBHandler, khsHandler *handlers.KHSHandler, penawaranHandler *handlers.PenawaranHandler, karyaAkhirHandler *handlers.KaryaAkhirHandler, kerjasamaHandler *handlers.KerjasamaHandler, realisasiUnitHandler *handlers.RealisasiUnitHandler, realisasiBulanHandler *handlers.RealisasiBulanHandler, penelitianHandler *handlers.PenelitianHandler, pengabdianHandler *handlers.PengabdianHandler, jurnalHandler *handlers.JurnalHandler, hkiHandler *handlers.HkiHandler, prosidingHandler *handlers.ProsidingHandler, bukuHandler *handlers.BukuHandler, beasiswaHandler *handlers.BeasiswaHandler, tracerHandler *handlers.TracerHandler, unitKerjaHandler *handlers.UnitKerjaHandler) {
	// implement swagger middleware
	r.Use(middlewares.SwaggerMockMiddleware())

	router := r.Group("/api/v1")
	{
		router.GET("/health-check", handlers.HealthCheckHandler)
	}

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/login", casHandler.LoginHandler)
		authGroup.GET("/callback", casHandler.CASMiddleware(), casHandler.CallbackHandler)
		authGroup.GET("/logout", casHandler.LogoutHandler)
		authGroup.GET("/health", casHandler.HealthCheckHandler)
	}

	private := router.Group("/", middlewares.AuthMiddleware())
	{
		setUserRoutes(private, userHandler)
		setUpMhsRoutes(private, mhs)
		setUpDashboardMhsRoutes(private, dashboardMhsHandler)
		setUpPerpemRoutes(private, perpemHandler)
		setUpAngketMhsRoutes(private, angketMhsHandler)
		setUpKritikSaranRoutes(private, kritikSaranHandler)
		setUpAgendaMengajarRoutes(private, agendaMengajarHandler)
		setUpMhsWisudaRoutes(private, mhsWisudaHandler)
		setUpRekapPMBRoutes(private, rekapPMBHandler)
		setUpKHSRoutes(private, khsHandler)
		setUpPenawaranRoutes(private, penawaranHandler)
		setUpKaryaAkhirRoutes(private, karyaAkhirHandler)
		setUpKerjasamaRoutes(private, kerjasamaHandler)
		setUpRealisasiUnitRoutes(private, realisasiUnitHandler)
		setUpRealisasiBulanRoutes(private, realisasiBulanHandler)
		setUpPenelitianRoutes(private, penelitianHandler)
		setUpPengabdianRoutes(private, pengabdianHandler)
		setUpJurnalRoutes(private, jurnalHandler)
		setUpHkiRoutes(private, hkiHandler)
		setUpProsidingRoutes(private, prosidingHandler)
		setUpBukuRoutes(private, bukuHandler)
		setUpBeasiswaRoutes(private, beasiswaHandler)
		setUpTracerRoutes(private, tracerHandler)
		setUpUnitKerjaRoutes(private, unitKerjaHandler)
	}
}

func setUserRoutes(rg *gin.RouterGroup, userHandler *handlers.UserHandler) {
	rg.GET("/user/details", userHandler.GetDataProfile)
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

func setUpPerpemRoutes(rg *gin.RouterGroup, perpemHandler *handlers.PerpemHandler) {
	rg.GET("/perpem", perpemHandler.GetPerpemFiltered)
}

func setUpAngketMhsRoutes(rg *gin.RouterGroup, angketMhsHandler *handlers.AngketMhsHandler) {
	rg.GET("/angket-mhs", angketMhsHandler.GetAngketMhsFiltered)
}

func setUpKritikSaranRoutes(rg *gin.RouterGroup, kritikSaranHandler *handlers.KritikSaranHandler) {
	rg.GET("/kritik-saran", kritikSaranHandler.GetKritikSaranFiltered)
}

func setUpAgendaMengajarRoutes(rg *gin.RouterGroup, agendaMengajarHandler *handlers.AgendaMengajarHandler) {
	rg.GET("/agenda-mengajar", agendaMengajarHandler.GetAgendaMengajarFiltered)
}

func setUpMhsWisudaRoutes(rg *gin.RouterGroup, mhsWisudaHandler *handlers.MhsWisudaHandler) {
	rg.GET("/mhs-wisuda", mhsWisudaHandler.GetMhsWisudaFiltered)
}

func setUpRekapPMBRoutes(rg *gin.RouterGroup, rekapPMBHandler *handlers.RekapPMBHandler) {
	rg.GET("/rekap-pmb", rekapPMBHandler.GetRekapPMBFiltered)
}

func setUpKHSRoutes(rg *gin.RouterGroup, khsHandler *handlers.KHSHandler) {
	rg.GET("/khs", khsHandler.GetKHSFiltered)
}

func setUpPenawaranRoutes(rg *gin.RouterGroup, penawaranHandler *handlers.PenawaranHandler) {
	rg.GET("/penawaran", penawaranHandler.GetPenawaranFiltered)
}

func setUpKaryaAkhirRoutes(rg *gin.RouterGroup, karyaAkhirHandler *handlers.KaryaAkhirHandler) {
	rg.GET("/karya-akhir", karyaAkhirHandler.GetKaryaAkhirFiltered)
}
func setUpKerjasamaRoutes(rg *gin.RouterGroup, kerjasamaHandler *handlers.KerjasamaHandler) {
	rg.GET("/kerjasama", kerjasamaHandler.GetKerjasamaFiltered)
}

func setUpRealisasiUnitRoutes(rg *gin.RouterGroup, realisasiUnitHandler *handlers.RealisasiUnitHandler) {
	rg.GET("/realisasi-unit", realisasiUnitHandler.GetRealisasiUnitFiltered)
}

func setUpRealisasiBulanRoutes(rg *gin.RouterGroup, realisasiBulanHandler *handlers.RealisasiBulanHandler) {
	rg.GET("/realisasi-bulan", realisasiBulanHandler.GetRealisasiBulanFiltered)
}

func setUpPenelitianRoutes(rg *gin.RouterGroup, penelitianHandler *handlers.PenelitianHandler) {
	rg.GET("/penelitian", penelitianHandler.GetPenelitianFiltered)
}

func setUpPengabdianRoutes(rg *gin.RouterGroup, pengabdianHandler *handlers.PengabdianHandler) {
	rg.GET("/pengabdian", pengabdianHandler.GetPengabdianFiltered)
}

func setUpJurnalRoutes(rg *gin.RouterGroup, jurnalHandler *handlers.JurnalHandler) {
	rg.GET("/jurnal", jurnalHandler.GetJurnalFiltered)
}

func setUpHkiRoutes(rg *gin.RouterGroup, hkiHandler *handlers.HkiHandler) {
	rg.GET("/hki", hkiHandler.GetHkiFiltered)
}

func setUpProsidingRoutes(rg *gin.RouterGroup, prosidingHandler *handlers.ProsidingHandler) {
	rg.GET("/prosiding", prosidingHandler.GetProsidingFiltered)
}

func setUpBukuRoutes(rg *gin.RouterGroup, bukuHandler *handlers.BukuHandler) {
	rg.GET("/buku", bukuHandler.GetBukuFiltered)
}

func setUpBeasiswaRoutes(rg *gin.RouterGroup, beasiswaHandler *handlers.BeasiswaHandler) {
	rg.GET("/beasiswa", beasiswaHandler.GetBeasiswaFiltered)
}

func setUpTracerRoutes(rg *gin.RouterGroup, tracerHandler *handlers.TracerHandler) {
	rg.GET("/tracer", tracerHandler.GetTracerFiltered)
}

func setUpUnitKerjaRoutes(rg *gin.RouterGroup, unitKerjaHandler *handlers.UnitKerjaHandler) {
	rg.GET("/unit-kerja", unitKerjaHandler.GetUnitKerja)
}
