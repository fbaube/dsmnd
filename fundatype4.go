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

// These [Fundatype]s expand on the SQLite fundamental types,
// mainly by adding "pkey" and "fkey". Each symbol is four
// upper case characters; each value is four lower-case.
// But in any case see https://sqlite.org/c3ref/c_blob.html
const (
	TABL Fundatype = "tabl" // This datum describes a table !
	//
	INTG = "1234" // "INT"   // SQLITE_INTEGER 1
	FLOT = "1.0f" // "FLOAT" // SQLITE_FLOAT   2
	TEXT = "AaZz" // "TEXT"  // SQLITE_TEXT    3
	BLOB = "blob" // "BLOB"  // SQLITE_BLOB    4
	NULL = "null" // "NULL"  // SQLITE_NULL    5
	PKEY = "pkey" // PRIMARY KEY (SQLite "INTEGER")
	FKEY = "fkey" // FOREIGN KEY (SQLite "INTEGER")
	LIST = "list" // table type, one per enumeration ??
	OTHR = "othr"
)
