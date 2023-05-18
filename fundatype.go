package dsmnd

// SQLite fundamental datatypes:
// see https://sqlite.org/c3ref/c_blob.html
const (
	SQLITE_INTEGER = 1 // 64-bit signed integer
	SQLITE_FLOAT   = 2 // 64-bit IEEE floating point number
	SQLITE_TEXT    = 3 // string
	SQLITE_BLOB    = 4
	SQLITE_NULL    = 5
)

// Fundatype is a string that (in a different way)
// specifies the (SQLite) type of a field.
type Fundatype string

// These [Fundatype]s expand on the SQLite fundamental types.
// But in any case see https://sqlite.org/c3ref/c_blob.html
const (
	D_TABL Fundatype = "Tbl" // This is describing a table !
	//
	D_INTG = "123"  // "INT"   // SQLITE_INTEGER 1
	D_FLOT = "1.0"  // "FLOAT" // SQLITE_FLOAT   2
	D_TEXT = "A-z"  // "TEXT"  // SQLITE_TEXT    3
	D_BLOB = "bin"  // "BLOB"  // SQLITE_BLOB    4
	D_NULL = "null" // "NULL"  // SQLITE_NULL    5
	D_PKEY = "pkey" // PRIMARY KEY (SQLite "INTEGER")
	D_FKEY = "fkey" // FOREIGN KEY (SQLite "INTEGER")
	D_LIST = "list" // table type, one per enumeration ??
	D_OTHR = "other"
)
