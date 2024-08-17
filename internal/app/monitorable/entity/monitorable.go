package entity

type Monitorable struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URL  string `db:"url"`
	// Method string `db:"method"` // get, post, put, delete
	// payload
}
