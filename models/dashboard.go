package models

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
	Level string `json:"level"`
}
