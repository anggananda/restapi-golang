package routes

import (
	"restapi-golang/handlers"
	"restapi-golang/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, casHandler *handlers.CASHandler, userHandler *handlers.UserHandler, mhs *handlers.MhsHandler, dosenHandler *handlers.DosenHandler, pegawaiHandler *handlers.PegawaiHandler, dashboardMhsHandler *handlers.DashboardMhsHandler, dashboardDosenHandler *handlers.DashboardDosenHandler, dashboardPegawaiHandler *handlers.DashboardPegawaiHandler, perpemHandler *handlers.PerpemHandler, evaluasiDosenHandler *handlers.EvaluasiDosenHandler, angketMhsHandler *handlers.AngketMhsHandler, kritikSaranHandler *handlers.KritikSaranHandler, agendaMengajarHandler *handlers.AgendaMengajarHandler, mhsWisudaHandler *handlers.MhsWisudaHandler, rekapPMBHandler *handlers.RekapPMBHandler, khsHandler *handlers.KHSHandler, penawaranHandler *handlers.PenawaranHandler, karyaAkhirHandler *handlers.KaryaAkhirHandler, kerjasamaHandler *handlers.KerjasamaHandler, realisasiUnitHandler *handlers.RealisasiUnitHandler, realisasiBulanHandler *handlers.RealisasiBulanHandler, penelitianHandler *handlers.PenelitianHandler, pengabdianHandler *handlers.PengabdianHandler, jurnalHandler *handlers.JurnalHandler, hkiHandler *handlers.HkiHandler, prosidingHandler *handlers.ProsidingHandler, bukuHandler *handlers.BukuHandler, beasiswaHandler *handlers.BeasiswaHandler, tracerHandler *handlers.TracerHandler, unitKerjaHandler *handlers.UnitKerjaHandler, statusHandler *handlers.StatusHandler) {
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
		setUpDosenRoutes(private, dosenHandler)
		setUpPegawaiRoutes(private, pegawaiHandler)
		setUpDashboardMhsRoutes(private, dashboardMhsHandler)
		setUpDashboardDosenRoutes(private, dashboardDosenHandler)
		setUpDashboardPegawaiRoutes(private, dashboardPegawaiHandler)
		setUpPerpemRoutes(private, perpemHandler)
		setUpEvaluasiDosenRoutes(private, evaluasiDosenHandler)
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
		setUpStatusRoutes(private, statusHandler)
	}
}

func setUserRoutes(rg *gin.RouterGroup, userHandler *handlers.UserHandler) {
	rg.GET("/user/details", userHandler.GetDataProfile)
}

func setUpMhsRoutes(rg *gin.RouterGroup, mhs *handlers.MhsHandler) {
	rg.GET("/mhs/:nim", mhs.GetDetailMhs)
	rg.GET("/mhs/history", mhs.GetMahasiswaHistoryFiltered)
	rg.GET("/mhs/history/export-csv", mhs.ExportMhsCSV)
}

func setUpDosenRoutes(rg *gin.RouterGroup, dosenHandler *handlers.DosenHandler) {
	rg.GET("/dosen/:niu", dosenHandler.GetDetailDosen)
	rg.GET("/dosen/history", dosenHandler.GetDosenHistoryFiltered)
	rg.GET("/dosen/history/export-csv", dosenHandler.ExportDosenCSV)
}

func setUpPegawaiRoutes(rg *gin.RouterGroup, pegawaiHandler *handlers.PegawaiHandler) {
	rg.GET("/pegawai/:niu", pegawaiHandler.GetDetailPegawai)
	rg.GET("/pegawai/history", pegawaiHandler.GetPegawaiHistoryFiltered)
	rg.GET("/pegawai/history/export-csv", pegawaiHandler.ExportPegawaiCSV)
}

func setUpDashboardMhsRoutes(rg *gin.RouterGroup, dashboardMhsHandler *handlers.DashboardMhsHandler) {
	rg.GET("/dashboard-mhs/overview", dashboardMhsHandler.GetDashboardMhsOverview)
	rg.GET("/dashboard-mhs/fakultas", dashboardMhsHandler.GetDrilldownMhsFakultas)
	rg.GET("/dashboard-mhs/jurusan", dashboardMhsHandler.GetDrilldownMhsJurusan)
	rg.GET("/dashboard-mhs/prodi", dashboardMhsHandler.GetDrilldownMhsProdi)
}

func setUpDashboardDosenRoutes(rg *gin.RouterGroup, dashboardDosenHandler *handlers.DashboardDosenHandler) {
	rg.GET("/dashboard-dosen/overview", dashboardDosenHandler.GetDashboardDosenOverview)
	rg.GET("/dashboard-dosen/fakultas", dashboardDosenHandler.GetDrilldownDosenFakultas)
	rg.GET("/dashboard-dosen/jurusan", dashboardDosenHandler.GetDrilldownDosenJurusan)
	rg.GET("/dashboard-dosen/prodi", dashboardDosenHandler.GetDrilldownDosenProdi)
}

func setUpDashboardPegawaiRoutes(rg *gin.RouterGroup, dashboardPegawaiHandler *handlers.DashboardPegawaiHandler) {
	rg.GET("/dashboard-pegawai/overview", dashboardPegawaiHandler.GetDashboardPegawaiOverview)
	rg.GET("/dashboard-pegawai/fakultas", dashboardPegawaiHandler.GetDrilldownPegawaiFakultas)
	rg.GET("/dashboard-pegawai/jurusan", dashboardPegawaiHandler.GetDrilldownPegawaiJurusan)
	rg.GET("/dashboard-pegawai/prodi", dashboardPegawaiHandler.GetDrilldownPegawaiProdi)
}

func setUpPerpemRoutes(rg *gin.RouterGroup, perpemHandler *handlers.PerpemHandler) {
	rg.GET("/perpem", perpemHandler.GetPerpemFiltered)
	rg.GET("/perpem/export-csv", perpemHandler.ExportPerpemCSV)
}
func setUpEvaluasiDosenRoutes(rg *gin.RouterGroup, evaluasiDosenHandler *handlers.EvaluasiDosenHandler) {
	rg.GET("/evaluasi-dosen", evaluasiDosenHandler.GetEvaluasiDosenFiltered)
	rg.GET("/evaluasi-dosen/export-csv", evaluasiDosenHandler.ExportEvaluasiDosenCSV)
}

func setUpAngketMhsRoutes(rg *gin.RouterGroup, angketMhsHandler *handlers.AngketMhsHandler) {
	rg.GET("/angket-mhs", angketMhsHandler.GetAngketMhsFiltered)
	rg.GET("/angket-mhs/export-csv", angketMhsHandler.ExportAngketMhsCSV)
}

func setUpKritikSaranRoutes(rg *gin.RouterGroup, kritikSaranHandler *handlers.KritikSaranHandler) {
	rg.GET("/kritik-saran", kritikSaranHandler.GetKritikSaranFiltered)
	rg.GET("/kritik-saran/export-csv", kritikSaranHandler.ExportKritikSaranCSV)
}

func setUpAgendaMengajarRoutes(rg *gin.RouterGroup, agendaMengajarHandler *handlers.AgendaMengajarHandler) {
	rg.GET("/agenda-mengajar", agendaMengajarHandler.GetAgendaMengajarFiltered)
	rg.GET("/agenda-mengajar/export-csv", agendaMengajarHandler.ExportAgendaMengajarCSV)
}

func setUpMhsWisudaRoutes(rg *gin.RouterGroup, mhsWisudaHandler *handlers.MhsWisudaHandler) {
	rg.GET("/mhs-wisuda", mhsWisudaHandler.GetMhsWisudaFiltered)
	rg.GET("/mhs-wisuda/export-csv", mhsWisudaHandler.ExportMhsWisudaCSV)
}

func setUpRekapPMBRoutes(rg *gin.RouterGroup, rekapPMBHandler *handlers.RekapPMBHandler) {
	rg.GET("/rekap-pmb", rekapPMBHandler.GetRekapPMBFiltered)
	rg.GET("/rekap-pmb/export-csv", rekapPMBHandler.ExportRekapPMBCSV)
}

func setUpKHSRoutes(rg *gin.RouterGroup, khsHandler *handlers.KHSHandler) {
	rg.GET("/khs", khsHandler.GetKHSFiltered)
	rg.GET("/khs/export-csv", khsHandler.ExportKhsCSV)
}

func setUpPenawaranRoutes(rg *gin.RouterGroup, penawaranHandler *handlers.PenawaranHandler) {
	rg.GET("/penawaran", penawaranHandler.GetPenawaranFiltered)
	rg.GET("/penawaran/export-csv", penawaranHandler.ExportPenawaranCSV)
}

func setUpKaryaAkhirRoutes(rg *gin.RouterGroup, karyaAkhirHandler *handlers.KaryaAkhirHandler) {
	rg.GET("/karya-akhir", karyaAkhirHandler.GetKaryaAkhirFiltered)
	rg.GET("/karya-akhir/export-csv", karyaAkhirHandler.ExportKaryaAkhirCSV)
}
func setUpKerjasamaRoutes(rg *gin.RouterGroup, kerjasamaHandler *handlers.KerjasamaHandler) {
	rg.GET("/kerjasama", kerjasamaHandler.GetKerjasamaFiltered)
	rg.GET("/kerjasama/export-csv", kerjasamaHandler.ExportKerjasamaCSV)
}

func setUpRealisasiUnitRoutes(rg *gin.RouterGroup, realisasiUnitHandler *handlers.RealisasiUnitHandler) {
	rg.GET("/realisasi-unit", realisasiUnitHandler.GetRealisasiUnitFiltered)
	rg.GET("/realisasi-unit/export-csv", realisasiUnitHandler.ExportRealisasiUnitCSV)
}

func setUpRealisasiBulanRoutes(rg *gin.RouterGroup, realisasiBulanHandler *handlers.RealisasiBulanHandler) {
	rg.GET("/realisasi-bulan", realisasiBulanHandler.GetRealisasiBulanFiltered)
	rg.GET("/realisasi-bulan/export-csv", realisasiBulanHandler.ExportRealisasiBulanCSV)
}

func setUpPenelitianRoutes(rg *gin.RouterGroup, penelitianHandler *handlers.PenelitianHandler) {
	rg.GET("/penelitian", penelitianHandler.GetPenelitianFiltered)
	rg.GET("/penelitian/export-csv", penelitianHandler.ExportPenelitianCSV)
}

func setUpPengabdianRoutes(rg *gin.RouterGroup, pengabdianHandler *handlers.PengabdianHandler) {
	rg.GET("/pengabdian", pengabdianHandler.GetPengabdianFiltered)
	rg.GET("/pengabdian/export-csv", pengabdianHandler.ExportPengabdianCSV)
}

func setUpJurnalRoutes(rg *gin.RouterGroup, jurnalHandler *handlers.JurnalHandler) {
	rg.GET("/jurnal", jurnalHandler.GetJurnalFiltered)
	rg.GET("/jurnal/export-csv", jurnalHandler.ExportJurnalCSV)
}

func setUpHkiRoutes(rg *gin.RouterGroup, hkiHandler *handlers.HkiHandler) {
	rg.GET("/hki", hkiHandler.GetHkiFiltered)
	rg.GET("/hki/export-csv", hkiHandler.ExportHkiCSV)
}

func setUpProsidingRoutes(rg *gin.RouterGroup, prosidingHandler *handlers.ProsidingHandler) {
	rg.GET("/prosiding", prosidingHandler.GetProsidingFiltered)
	rg.GET("/prosiding/export-csv", prosidingHandler.ExportProsidingCSV)
}

func setUpBukuRoutes(rg *gin.RouterGroup, bukuHandler *handlers.BukuHandler) {
	rg.GET("/buku", bukuHandler.GetBukuFiltered)
	rg.GET("/buku/export-csv", bukuHandler.ExportBukuCSV)
}

func setUpBeasiswaRoutes(rg *gin.RouterGroup, beasiswaHandler *handlers.BeasiswaHandler) {
	rg.GET("/beasiswa", beasiswaHandler.GetBeasiswaFiltered)
	rg.GET("/beasiswa/export-csv", beasiswaHandler.ExportBeasiswaCSV)
}

func setUpTracerRoutes(rg *gin.RouterGroup, tracerHandler *handlers.TracerHandler) {
	rg.GET("/tracer", tracerHandler.GetTracerFiltered)
	rg.GET("/tracer/export-csv", tracerHandler.ExportTracerCSV)
}

func setUpUnitKerjaRoutes(rg *gin.RouterGroup, unitKerjaHandler *handlers.UnitKerjaHandler) {
	rg.GET("/unit-kerja", unitKerjaHandler.GetUnitKerja)
}

func setUpStatusRoutes(rg *gin.RouterGroup, statusHandler *handlers.StatusHandler) {
	rg.GET("/status-mhs", statusHandler.GetStatusMahasiswa)
	rg.GET("/status-pegawai", statusHandler.GetStatusPegawai)
	rg.GET("/status-keaktifan-pegawai", statusHandler.GetStatusKeaktifanPegawai)
}
