// models/dashboard.go
package models

type DashboardRequest struct {
	Tahun    int    `form:"tahun" binding:"required"`
	Semester int    `form:"semester" binding:"required"`
	Status   string `form:"status"` // A, C, D
	Fakultas string `form:"fakultas"`
	Jurusan  string `form:"jurusan"`
	Prodi    string `form:"prodi"`
	Level    string `form:"level"` // overview, fakultas, jurusan, prodi
}

type DashboardCard struct {
	Title     string `json:"title"`
	Value     int64  `json:"value"`
	Status    string `json:"status"`
	Drilldown bool   `json:"drilldown"`
}

type DashboardCardPegawai struct {
	Title             string `json:"title"`
	Value             int64  `json:"value"`
	IDStatusPegawai   int64  `json:"id_status_pegawai"`
	IDStatusKeaktifan int64  `json:"id_status_keaktifan"`
	Drilldown         bool   `json:"drilldown"`
}

type DrilldownItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value int64  `json:"value"`
	Level string `json:"level"` // fakultas, jurusan, prodi
}

type DashboardResponse struct {
	Status    string          `json:"status"`
	Level     string          `json:"level"`     // overview, fakultas, jurusan, prodi
	Cards     []DashboardCard `json:"cards"`     // untuk level overview
	Drilldown []DrilldownItem `json:"drilldown"` // untuk level drill-down
	Total     int64           `json:"total"`     // total untuk level tersebut
}
