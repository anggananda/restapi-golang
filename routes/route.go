package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"restapi-golang/handlers"
	"restapi-golang/middlewares"
	"time"

	"github.com/gin-gonic/gin"
)

// Middleware untuk prevent caching pada auth endpoints
func NoCacheMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")
		c.Writer.Header().Set("Vary", "*")
		c.Next()
	}
}

func SetUpRoutes(r *gin.Engine, casHandler *handlers.CASHandler, userHandler *handlers.UserHandler, mhs *handlers.MhsHandler, dashboardMhsHandler *handlers.DashboardMhsHandler, perpemHandler *handlers.PerpemHandler, angketMhsHandler *handlers.AngketMhsHandler, kritikSaranHandler *handlers.KritikSaranHandler, agendaMengajarHandler *handlers.AgendaMengajarHandler, mhsWisudaHandler *handlers.MhsWisudaHandler, rekapPMBHandler *handlers.RekapPMBHandler, khsHandler *handlers.KHSHandler, penawaranHandler *handlers.PenawaranHandler, karyaAkhirHandler *handlers.KaryaAkhirHandler, realisasiUnitHandler *handlers.RealisasiUnitHandler, realisasiBulanHandler *handlers.RealisasiBulanHandler, penelitianHandler *handlers.PenelitianHandler, pengabdianHandler *handlers.PengabdianHandler, jurnalHandler *handlers.JurnalHandler, hkiHandler *handlers.HkiHandler, prosidingHandler *handlers.ProsidingHandler, bukuHandler *handlers.BukuHandler, beasiswaHandler *handlers.BeasiswaHandler, tracerHandler *handlers.TracerHandler) {

	router := r.Group("/api/v1")
	{
		router.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
		})
	}

	// Auth Group dengan middleware untuk prevent caching
	authGroup := router.Group("/auth")
	authGroup.Use(NoCacheMiddleware()) // Apply no-cache middleware untuk semua auth routes
	{
		// Login endpoint - tidak perlu CAS middleware di sini
		authGroup.GET("/login", casHandler.LoginHandler)

		// Callback endpoint - perlu CAS middleware untuk handle authentication
		authGroup.GET("/callback", casHandler.CASMiddleware(), casHandler.CallbackHandler)

		// Regular logout
		authGroup.GET("/logout", casHandler.LogoutHandler)

		// Force logout - untuk kasus yang lebih agresif
		authGroup.GET("/force-logout", casHandler.ForceLogoutHandler)

		// Destroy session - menghancurkan JSESSIONID dan semua CAS session
		authGroup.GET("/destroy-session", casHandler.DestroySessionHandler)

		// Pre-logout - clear session sebelum login baru
		authGroup.POST("/pre-logout", casHandler.PreLogoutHandler)

		// Super aggressive logout - untuk testing
		authGroup.GET("/nuclear-logout", func(c *gin.Context) {
			service := c.Query("service")
			log.Printf("Nuclear logout called with service: %s", service)

			// Clear semua cookies dengan cara paling agresif
			allCookieNames := []string{
				"CASTGC", "TGC", "JSESSIONID", "SESSION",
				"cas-session", "TICKET_GRANTING_COOKIE",
				"token", "session", "user", "auth",
			}

			// Clear dengan semua variasi path dan domain
			for _, cookieName := range allCookieNames {
				// Clear dengan empty domain
				c.SetCookie(cookieName, "", -1, "/", "", false, true)
				c.SetCookie(cookieName, "", -1, "/cas", "", false, true)
				c.SetCookie(cookieName, "", -1, "/api", "", false, true)
			}

			// Clear dari existing cookies
			for _, cookie := range c.Request.Cookies() {
				c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
			}

			// Set aggressive headers
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Writer.Header().Set("Pragma", "no-cache")
			c.Writer.Header().Set("Expires", "0")
			c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\", \"executionContexts\"")

			// Build logout URL dengan parameter paling agresif
			timestamp := time.Now().UnixNano()
			var logoutURL string

			if service != "" {
				// Logout dengan service URL dan semua parameter forcing
				logoutURL = fmt.Sprintf("%s/logout?service=%s&renew=true&gateway=true&_=%d",
					casHandler.GetCASURL().String(),
					url.QueryEscape(service),
					timestamp)
			} else {
				// Logout tanpa service tapi dengan forcing parameters
				logoutURL = fmt.Sprintf("%s/logout?renew=true&gateway=true&_=%d",
					casHandler.GetCASURL().String(),
					timestamp)
			}

			log.Printf("Nuclear logout redirect to: %s", logoutURL)
			c.Redirect(http.StatusFound, logoutURL)
		})

		// Clear session endpoint untuk debugging/manual clearing
		authGroup.POST("/clear-session", func(c *gin.Context) {
			// Set headers untuk prevent caching dan clear site data
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Writer.Header().Set("Pragma", "no-cache")
			c.Writer.Header().Set("Expires", "0")
			c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\"")

			// Clear semua cookies yang ada
			cookies := c.Request.Cookies()
			for _, cookie := range cookies {
				// Clear dengan path dan domain asli
				c.SetCookie(cookie.Name, "", -1, cookie.Path, cookie.Domain, false, true)

				// Clear dengan variasi path dan domain
				c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
				c.SetCookie(cookie.Name, "", -1, "/", c.Request.Host, false, true)
			}

			// Clear known CAS cookies secara eksplisit
			casCookies := []string{"CASTGC", "TGC", "JSESSIONID", "SESSION"}
			for _, cookieName := range casCookies {
				c.SetCookie(cookieName, "", -1, "/", "", false, true)
				c.SetCookie(cookieName, "", -1, "/", c.Request.Host, false, true)
			}

			c.JSON(http.StatusOK, gin.H{
				"message":   "Session cleared successfully",
				"timestamp": time.Now().Unix(),
				"cleared_cookies": len(cookies),
			})
		})

		// Health check untuk auth system
		authGroup.GET("/health", casHandler.HealthCheckHandler)

		// Status endpoint untuk check authentication state
		authGroup.GET("/status", func(c *gin.Context) {
			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			c.Writer.Header().Set("Pragma", "no-cache")
			c.Writer.Header().Set("Expires", "0")

			// Simple status check tanpa CAS middleware
			c.JSON(http.StatusOK, gin.H{
				"authenticated": false,
				"timestamp":     time.Now().Unix(),
				"server":        "ready",
			})
		})
	}

	// Protected routes group (existing routes yang memerlukan authentication)
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
}
}

// package routes

// import (
// 	"net/http"
// 	"restapi-golang/handlers"
// 	"restapi-golang/middlewares"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// // Middleware untuk prevent caching pada auth endpoints
// func NoCacheMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 		c.Writer.Header().Set("Pragma", "no-cache")
// 		c.Writer.Header().Set("Expires", "0")
// 		c.Writer.Header().Set("Vary", "*")
// 		c.Next()
// 	}
// }

// func SetUpRoutes(r *gin.Engine, casHandler *handlers.CASHandler, userHandler *handlers.UserHandler, mhs *handlers.MhsHandler, dashboardMhsHandler *handlers.DashboardMhsHandler, perpemHandler *handlers.PerpemHandler, angketMhsHandler *handlers.AngketMhsHandler, kritikSaranHandler *handlers.KritikSaranHandler, agendaMengajarHandler *handlers.AgendaMengajarHandler, mhsWisudaHandler *handlers.MhsWisudaHandler, rekapPMBHandler *handlers.RekapPMBHandler, khsHandler *handlers.KHSHandler, penawaranHandler *handlers.PenawaranHandler, karyaAkhirHandler *handlers.KaryaAkhirHandler, realisasiUnitHandler *handlers.RealisasiUnitHandler, realisasiBulanHandler *handlers.RealisasiBulanHandler, penelitianHandler *handlers.PenelitianHandler, pengabdianHandler *handlers.PengabdianHandler, jurnalHandler *handlers.JurnalHandler, hkiHandler *handlers.HkiHandler, prosidingHandler *handlers.ProsidingHandler, bukuHandler *handlers.BukuHandler, beasiswaHandler *handlers.BeasiswaHandler, tracerHandler *handlers.TracerHandler) {

// 	router := r.Group("/api/v1")
// 	{
// 		router.GET("/health", func(c *gin.Context) {
// 			c.JSON(http.StatusOK, gin.H{"message": "OK"})
// 		})
// 	}

// 	// Auth Group dengan middleware untuk prevent caching
// 	authGroup := router.Group("/auth")
// 	authGroup.Use(NoCacheMiddleware()) // Apply no-cache middleware untuk semua auth routes
// 	{
// 		// Login endpoint - tidak perlu CAS middleware di sini
// 		authGroup.GET("/login", casHandler.LoginHandler)

// 		// Callback endpoint - perlu CAS middleware untuk handle authentication
// 		authGroup.GET("/callback", casHandler.CASMiddleware(), casHandler.CallbackHandler)

// 		// Regular logout
// 		authGroup.GET("/logout", casHandler.LogoutHandler)

// 		// Force logout - untuk kasus yang lebih agresif
// 		authGroup.GET("/force-logout", casHandler.ForceLogoutHandler)

// 		// Clear session endpoint untuk debugging/manual clearing
// 		authGroup.POST("/clear-session", func(c *gin.Context) {
// 			// Set headers untuk prevent caching dan clear site data
// 			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 			c.Writer.Header().Set("Pragma", "no-cache")
// 			c.Writer.Header().Set("Expires", "0")
// 			c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\"")

// 			// Clear semua cookies yang ada
// 			cookies := c.Request.Cookies()
// 			for _, cookie := range cookies {
// 				// Clear dengan path dan domain asli
// 				c.SetCookie(cookie.Name, "", -1, cookie.Path, cookie.Domain, false, true)

// 				// Clear dengan variasi path dan domain
// 				c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
// 				c.SetCookie(cookie.Name, "", -1, "/", c.Request.Host, false, true)
// 			}

// 			// Clear known CAS cookies secara eksplisit
// 			casCookies := []string{"CASTGC", "TGC", "JSESSIONID", "SESSION"}
// 			for _, cookieName := range casCookies {
// 				c.SetCookie(cookieName, "", -1, "/", "", false, true)
// 				c.SetCookie(cookieName, "", -1, "/", c.Request.Host, false, true)
// 			}

// 			c.JSON(http.StatusOK, gin.H{
// 				"message":         "Session cleared successfully",
// 				"timestamp":       time.Now().Unix(),
// 				"cleared_cookies": len(cookies),
// 			})
// 		})

// 		// Health check untuk auth system
// 		authGroup.GET("/health", casHandler.HealthCheckHandler)

// 		// Status endpoint untuk check authentication state
// 		authGroup.GET("/status", func(c *gin.Context) {
// 			c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 			c.Writer.Header().Set("Pragma", "no-cache")
// 			c.Writer.Header().Set("Expires", "0")

// 			// Simple status check tanpa CAS middleware
// 			c.JSON(http.StatusOK, gin.H{
// 				"authenticated": false,
// 				"timestamp":     time.Now().Unix(),
// 				"server":        "ready",
// 			})
// 		})
// 	}

// private := router.Group("/", middlewares.AuthMiddleware())
// {
// 	setUserRoutes(private, userHandler)
// 	setUpMhsRoutes(private, mhs)
// 	setUpDashboardMhsRoutes(private, dashboardMhsHandler)
// 	setUpPerpemRoutes(private, perpemHandler)
// 	setUpAngketMhsRoutes(private, angketMhsHandler)
// 	setUpKritikSaranRoutes(private, kritikSaranHandler)
// 	setUpAgendaMengajarRoutes(private, agendaMengajarHandler)
// 	setUpMhsWisudaRoutes(private, mhsWisudaHandler)
// 	setUpRekapPMBRoutes(private, rekapPMBHandler)
// 	setUpKHSRoutes(private, khsHandler)
// 	setUpPenawaranRoutes(private, penawaranHandler)
// 	setUpKaryaAkhirRoutes(private, karyaAkhirHandler)
// 	setUpRealisasiUnitRoutes(private, realisasiUnitHandler)
// 	setUpRealisasiBulanRoutes(private, realisasiBulanHandler)
// 	setUpPenelitianRoutes(private, penelitianHandler)
// 	setUpPengabdianRoutes(private, pengabdianHandler)
// 	setUpJurnalRoutes(private, jurnalHandler)
// 	setUpHkiRoutes(private, hkiHandler)
// 	setUpProsidingRoutes(private, prosidingHandler)
// 	setUpBukuRoutes(private, bukuHandler)
// 	setUpBeasiswaRoutes(private, beasiswaHandler)
// 	setUpTracerRoutes(private, tracerHandler)
// }
// }

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
