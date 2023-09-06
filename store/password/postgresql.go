package password

// Postgres implements the Repository interface for PostgreSQL.
type Postgres struct {
}

// NewPostgres returns a new instance of Postgres.
func NewPostgres() *Postgres {
	return &Postgres{}
}
