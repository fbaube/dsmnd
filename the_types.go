package dsmnd

type DB_type string

const (
     // DB_SQLite is NOT "sqlite3"! 
	DB_SQLite DB_type = "sqlite"
)

// DbColInDb describes a column as-is in the DB (as obtained via
// reflection), and has a slot to include the value (as a string).
// Used only in utils/repo/sqlite/meta_table.go
// . 
type DbColInDb Datum

// TypeType enumerates the basic types defined in this package.
// [Datum] is not in this list.
// . 
type TypeType string 

const(
	TT_NIL = TypeType("nil")
	// TT_SQLITE is [SqliteDatatype] 
        TT_SQLITE TypeType = TypeType("sqlite")
	// TT_BASIC is [BasicDatatype] 
	TT_BASIC = TypeType("basic")
	// TT_SEMFIELD is [SemanticFieldType] 
	TT_SEMFIELD = TypeType("semfield")
	// TT_SEMLIST is [SemanticListType]
	TT_SEMLIST = TypeType("semlist")
	// TT_SEMCLXN is [SemanticClxnType]
	TT_SEMCLXN = TypeType("semclxn")
)

