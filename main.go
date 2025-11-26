package main

import (
	"fmt"
	"os"
	"restapi-golang/config"
	"restapi-golang/docs"
	"restapi-golang/handlers"
	"restapi-golang/middlewares"
	"restapi-golang/repositories"
	"restapi-golang/routes"
	"restapi-golang/services"

	_ "restapi-golang/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Executive Information System RESTful API
// @version         1.0
// @description  This repository contains the **RESTful API** for the **Executive Information System (EIS) Undiksha**, built with **Golang** using the **Gin framework**.
// @description The API consumes data from **MongoDB** (populated via ETL from SQL in a separate repository) and uses **Redis** for caching.
// @description The system is integrated with **Undiksha Single Sign-On (SSO)** and implements **JWT Authentication** for secure access.

// @termsOfService  https://github.com/anggananda
// @contact.name    Dwi Angga
// @contact.url     https://github.com/anggananda
// @contact.email   anggadek857@gmail.com

// @license.name    MIT License
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /api/v1
// @schemes         http https

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	router := gin.Default()
	config.ConnectDB()
	// Redis Hold Dulu
	config.ConnectRedis()

	router.Use(middlewares.CorsMiddleware())
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "localhost:8080"
	}

	scheme := os.Getenv("APP_SCHEME")
	if scheme == "" {
		scheme = "http"
	}

	docs.SwaggerInfo.Host = host
	docs.SwaggerInfo.Schemes = []string{scheme}
	docs.SwaggerInfo.BasePath = "/api/v1"

	userRepo := repositories.NewUserMongoRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	casHandler, err := handlers.NewCASHandler(config.CASUrl, config.FrontendUrl, config.HostUrl, userService)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize CAS: %v", err))
	}

	mhsRepo := repositories.NewMhsMongoRepository(config.DB)
	mhsService := services.NewMhsService(mhsRepo)
	mhsHandler := handlers.NewMhsHandler(mhsService)

	dosenRepo := repositories.NewDosenMongoRepository(config.DB)
	dosenService := services.NewDosenService(dosenRepo)
	dosenHandler := handlers.NewDosenHandler(dosenService)

	pegawaiRepo := repositories.NewPegawaiMongoRepository(config.DB)
	pegawaiService := services.NewPegawaiService(pegawaiRepo)
	pegawaiHandler := handlers.NewPegawaiHandler(pegawaiService)

	dashboardMhsRepo := repositories.NewDashboardMhsRepository(config.DB)
	dashboardMhsservice := services.NewDashboardMhsService(dashboardMhsRepo)
	dashboardMhsHandler := handlers.NewDashboardMhsHandler(dashboardMhsservice)

	dashboardDosenRepo := repositories.NewDashboardDosenRepository(config.DB)
	dashboardDosenservice := services.NewDashboardDosenService(dashboardDosenRepo)
	dashboardDosenHandler := handlers.NewDashboardDosenHandler(dashboardDosenservice)

	dashboardPegawaiRepo := repositories.NewDashboardPegawaiRepository(config.DB)
	dashboardPegawaiservice := services.NewDashboardPegawaiService(dashboardPegawaiRepo)
	dashboardPegawaiHandler := handlers.NewDashboardPegawaiHandler(dashboardPegawaiservice)

	perpemRepo := repositories.NewPerpemMongoRepository(config.DB)
	perpemService := services.NewPerpemService(perpemRepo)
	perpemHandler := handlers.NewPerpemHandler(perpemService)

	evaluasiDosenRepo := repositories.NewEvaluasiDosenMongoRepository(config.DB)
	evaluasiDosenService := services.NewEvaluasiDosenService(evaluasiDosenRepo)
	evaluasiDosenHandler := handlers.NewEvaluasiDosenHandler(evaluasiDosenService)

	angketMhsRepo := repositories.NewAngketMhsMongoRepository(config.DB)
	angketMhsService := services.NewAngketMhsService(angketMhsRepo)
	angketMhsHandler := handlers.NewAngketMhsHandler(angketMhsService)

	kritikSaranRepo := repositories.NewKritikSaranMongoRepository(config.DB)
	kritikSaranService := services.NewKritikSaranService(kritikSaranRepo)
	kritikSaranHandler := handlers.NewKritikSaranHandler(kritikSaranService)

	agendaMengajarRepo := repositories.NewAgendaMengajarMongoRepository(config.DB)
	agendaMengajarService := services.NewAgendaMengajarService(agendaMengajarRepo)
	agendaMengajarHandler := handlers.NewAgendaMengajarHandler(agendaMengajarService)

	mhsWisudaRepo := repositories.NewMhsWisudaMongoRepository(config.DB)
	mhsWisudaService := services.NewMhsWisudaService(mhsWisudaRepo)
	mhsWisudaHandler := handlers.NewMhsWisudaHandler(mhsWisudaService)

	rekapPMBRepo := repositories.NewRekapPMBMongoRepository(config.DB)
	rekapPMBService := services.NewRekapPMBService(rekapPMBRepo)
	rekapPMBHandler := handlers.NewRekapPMBHandler(rekapPMBService)

	khsRepo := repositories.NewKHSMongoRepository(config.DB)
	khsService := services.NewKHSService(khsRepo)
	khsHandler := handlers.NewKHSHandler(khsService)

	penawaranRepo := repositories.NewPenawaranMongoRepository(config.DB)
	penawaranService := services.NewPenawaranService(penawaranRepo)
	penawaranHandler := handlers.NewPenawaranHandler(penawaranService)

	karyaAkhirRepo := repositories.NewKaryaAkhirMongoRepository(config.DB)
	karyaAkhirService := services.NewKaryaAkhirService(karyaAkhirRepo)
	karyaAkhirHandler := handlers.NewKaryaAkhirHandler(karyaAkhirService)

	kerjasamaRepo := repositories.NewKerjasamaMongoRepository(config.DB)
	kerjasamaService := services.NewKerjasamaService(kerjasamaRepo)
	kerjasamaHandler := handlers.NewKerjasamaHandler(kerjasamaService)

	realisasiUnitRepo := repositories.NewRealisasiUnitMongoRepository(config.DB)
	realisasiUnitService := services.NewRealisasiUnitService(realisasiUnitRepo)
	realisasiUnitHandler := handlers.NewRealisasiUnitHandler(realisasiUnitService)

	realisasiBulanRepo := repositories.NewRealisasiBulanMongoRepository(config.DB)
	realisasiBulanService := services.NewRealisasiBulanService(realisasiBulanRepo)
	realisasiBulanHandler := handlers.NewRealisasiBulanHandler(realisasiBulanService)

	//Module Kinerja
	penelitianRepo := repositories.NewPenelitianMongoRepository(config.DB)
	penelitianService := services.NewPenelitianService(penelitianRepo)
	penelitianHandler := handlers.NewPenelitianHandler(penelitianService)

	pengabdianRepo := repositories.NewPengabdianMongoRepository(config.DB)
	pengabdianService := services.NewPengabdianService(pengabdianRepo)
	pengabdianHandler := handlers.NewPengabdianHandler(pengabdianService)

	jurnalRepo := repositories.NewJurnalMongoRepository(config.DB)
	jurnalService := services.NewJurnalService(jurnalRepo)
	jurnalHandler := handlers.NewJurnalHandler(jurnalService)

	hkiRepo := repositories.NewHkiMongoRepository(config.DB)
	hkiService := services.NewHkiService(hkiRepo)
	hkiHandler := handlers.NewHkiHandler(hkiService)

	prosidingRepo := repositories.NewProsidingMongoRepository(config.DB)
	prosidingService := services.NewProsidingService(prosidingRepo)
	prosidingHandler := handlers.NewProsidingHandler(prosidingService)

	bukuRepo := repositories.NewBukuMongoRepository(config.DB)
	bukuService := services.NewBukuService(bukuRepo)
	bukuHandler := handlers.NewBukuHandler(bukuService)

	beasiswaRepo := repositories.NewBeasiswaMongoRepository(config.DB)
	beasiswaService := services.NewBeasiswaService(beasiswaRepo)
	beasiswaHandler := handlers.NewBeasiswaHandler(beasiswaService)

	tracerRepo := repositories.NewTracerMongoRepository(config.DB)
	tracerService := services.NewTracerService(tracerRepo)
	tracerHandler := handlers.NewTracerHandler(tracerService)

	unitKerjaRepo := repositories.NewUnitKerjaMongoRepository(config.DB)
	unitKerjaService := services.NewUnitKerjaService(unitKerjaRepo)
	unitKerjaHandler := handlers.NewUnitKerjaHandler(unitKerjaService, config.RDB)

	statusMahasiswaRepo := repositories.NewStatusMahasiswaMongoRepository(config.DB)
	statusPegawaiRepo := repositories.NewStatusPegawaiMongoRepository(config.DB)
	statusKeaktifanPegawaiRepo := repositories.NewStatusKeaktifanPegawaiMongoRepository(config.DB)

	statusService := services.NewStatusService(statusMahasiswaRepo, statusPegawaiRepo, statusKeaktifanPegawaiRepo)
	statusHandler := handlers.NewStatusHandler(statusService, config.RDB)

	routes.SetUpRoutes(router, casHandler, userHandler, mhsHandler, dosenHandler, pegawaiHandler, dashboardMhsHandler, dashboardDosenHandler, dashboardPegawaiHandler, perpemHandler, evaluasiDosenHandler, angketMhsHandler, kritikSaranHandler, agendaMengajarHandler, mhsWisudaHandler, rekapPMBHandler, khsHandler, penawaranHandler, karyaAkhirHandler, kerjasamaHandler, realisasiUnitHandler, realisasiBulanHandler, penelitianHandler, pengabdianHandler, jurnalHandler, hkiHandler, prosidingHandler, bukuHandler, beasiswaHandler, tracerHandler, unitKerjaHandler, statusHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
