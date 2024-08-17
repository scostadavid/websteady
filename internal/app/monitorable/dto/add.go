package dto

type AddMonitorable struct {
	Name string `db:"name"`
	URL  string `db:"url"`
}
