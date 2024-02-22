package entity

type Monitorable struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	URI  string `db:"uri"`
	// Method string `db:"method"` // get, post, put, delete
}
