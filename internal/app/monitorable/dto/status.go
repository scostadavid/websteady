package dto

type MonitorableStatus struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URL  string `db:"url"`
	Up   bool   `db:"up"`
}
