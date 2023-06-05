package dsmnd

type DB_type string

const (
	DB_SQLite DB_type = "sqlite"
)

// DbDescr is generic.
// type DbDescr Datum

// DbTblSpec specifies a DB table (but not its columns).
// Field usage is as follows:
//   - Fundatype: TABL
//   - StorName: a short name for use in field names (primary & foreign keys)
//   - DispName: a long name, for the name of the table in the DB
//   - Description: long description
type TableSpec Datum // DbDescr

// ColumnSpec specifies a datum (i.e. a struct field and/or a
// DB column), including its generic/portable/DB-independent
// representation using the enumeration Datum.Fundatype).
// Some values for common DB columns are defined in file
// datadxnry.go . Field usage is as follows:
//   - Fundatype: TEXT or INTG // TODO Replace w semantic
//   - StorName: the field name in the DB
//   - DispName: short description
//   - Description: long description
type ColumnSpec Datum // DbDescr

// DbColInDb describes a column as-is in the DB (as obtained via
// reflection), and has a slot to include the value (as a string).
// Usd only in utils/repo/sqlite/meta_table.go
type DbColInDb Datum // DbDescr
