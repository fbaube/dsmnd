package dsmnd

// === aside: Semantics (for uses, see other files s?t_*.go) ===

type SemanticType	string
type SemanticDescriptor Datum

// === this file: Basic types ===

// SqliteDatatype is one of the five basic ones defined
// by & for SQLite, plus the addition of a date-time.
// See https://sqlite.org/c3ref/c_blob.html
//
// There's also complicating factors like
// NUMERIC (and REAL).
// . 
type SqliteDatatype int

const(
	SQLITE_INTEGER = SqliteDatatype(1) // 64-bit signed integer
	SQLITE_FLOAT   = SqliteDatatype(2) // 64-bit IEEE FP number
	SQLITE_TEXT    = SqliteDatatype(3) // string; incl JSON 
	SQLITE_BLOB    = SqliteDatatype(4) // incl JSONB 
	SQLITE_NULL    = SqliteDatatype(5)
	SQLYT_DATETIME = SqliteDatatype(6) // dusnt exist, but should 
)

// BasicDatatype is a string that expands upon the SQLite 
// datatypes, mainly by adding "key" and "list" and "clxn"
// (collection).
//  - Symbol names: "BDT_" + FOUR UPPER CASE letters
//  - Symbol values: four lower case letters
//  - Lists and Collections are included
//
// These are appropriate for describing DB schemas.
// For more detail tho, use [SemFieldTypes]. 
//
// But in any case see https://sqlite.org/c3ref/c_blob.html
// .
type BasicDatatype string

// NOTE: These BasicDatatype's _ARE_ USED

const (
BDT_INTG = BasicDatatype("1234") // INT   // SQLITE_INTEGER 1
BDT_FLOT = BasicDatatype("1.0f") // FLOAT // SQLITE_FLOAT   2
BDT_TEXT = BasicDatatype("AaZz") // TEXT  // SQLITE_TEXT    3
BDT_BLOB = BasicDatatype("blOb") // BLOB  // SQLITE_BLOB    4
BDT_NULL = BasicDatatype("null") // NULL  // SQLITE_NULL    5
BDT_DYTM = BasicDatatype("dytm") // DATETIME SQLYT_DATETIME 6
BDT_KEYY = BasicDatatype("keyy") // PRIMARY/FOREIGN/OTHER KEY (SQLite "INTEGER")
BDT_LIST = BasicDatatype("list") // List (simple one-dimensional lists) 
BDT_CLXN = BasicDatatype("clxn") // Collection (more-complicated data strux)
BDT_OTHR = BasicDatatype("othr") // reserved: expansion 
BDT_NONE = BasicDatatype("none") // reserved 
)

func (bdt BasicDatatype) S() string {
     return string(bdt)
}

