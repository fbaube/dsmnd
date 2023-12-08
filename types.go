package dsmnd

type DB_type string

const (
	DB_SQLite DB_type = "sqlite"
)

// DbColInDb describes a column as-is in the DB (as obtained via
// reflection), and has a slot to include the value (as a string).
// Used only in utils/repo/sqlite/meta_table.go
// . 
type DbColInDb Datum
