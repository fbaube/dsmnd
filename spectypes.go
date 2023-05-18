package dsmnd

type DB_type string

const (
	DB_SQLite DB_type = "sqlite"
)

// DbDescr is generic.
type DbDescr Datum

// DbTblSpec specifies a DB table (but not its columns).
// Field usage is as follows:
//   - TxtIntKeyEtc: D_TBL
//   - Code: a short name for use in field names (primary & foreign keys)
//   - Name: a long name, for the name of the table in the DB
//   - Descr: long description
type TableSpec DbDescr

// ColumnSpec specifies a datum (i.e. a struct field and/or a
// DB column), including its generic/portable/DB-independent
// representation using the enumeration Datum.TxtIntKeyEtc).
// Some values for common DB columns are defined in the D_*
// series below. Field usage is as follows:
//   - TxtIntKeyEtc: D_TXT or D_INT
//   - Code: the field name in the DB
//   - Name: short description
//   - Descr: long description
type ColumnSpec DbDescr

// DbColInDb describes a column as-is in the DB (as obtained via
// reflection), and has a slot to include the value (as a string).
// Not used so far.
type DbColInDb DbDescr
